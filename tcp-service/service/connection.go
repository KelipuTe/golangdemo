package service

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/protocol"
	"demo-golang/tcp-service/protocol/abs"
	"demo-golang/tcp-service/protocol/http"
	"demo-golang/tcp-service/protocol/stream"
	"demo-golang/tcp-service/protocol/websocket"
	"errors"
	"fmt"
	"io"
	"net"
)

// TCPConnection tcp连接，封装 net.Conn
type TCPConnection struct {
	runStatus       config.RunStatus //连接状态
	belongToService *TCPService      //连接所属的服务端

	protocolName    string        //协议名称
	protocolHandler abs.HandlerI9 //协议处理器

	netConn          net.Conn //tcp连接本体
	recvBuffer       []byte   //接收缓冲区
	recvBufferMaxLen uint64   //接收缓冲区最大长度
	recvBufferNowLen uint64   //接收缓冲区当前长度
}

func NewTCPConnection(service *TCPService, conn net.Conn) *TCPConnection {
	return &TCPConnection{
		runStatus:       config.RunStatusOff,
		belongToService: service,

		protocolName:    service.protocolName,
		protocolHandler: protocol.NewHandler(service.protocolName),

		netConn:          conn,
		recvBuffer:       make([]byte, config.RecvBufferMaxLen),
		recvBufferMaxLen: config.RecvBufferMaxLen,
		recvBufferNowLen: 0,
	}
}

// IsRun 是不是运行中
func (c *TCPConnection) IsRun() bool {
	return c.runStatus == config.RunStatusOn
}

// IsDebug 是不是debug模式
func (c *TCPConnection) IsDebug() bool {
	return c.belongToService.IsDebug()
}

// GetProtocolHandler 获取协议处理器
func (c *TCPConnection) GetProtocolHandler() abs.HandlerI9 {
	return c.protocolHandler
}

// GetNetConnRemoteAddr 获取连接的ip和端口号
func (c *TCPConnection) GetNetConnRemoteAddr() string {
	return c.netConn.RemoteAddr().String()
}

// HandleConnection 处理连接
func (c *TCPConnection) HandleConnection() {
	for c.IsRun() {
		//net.Conn.Read，系统调用，从socket读取数据
		byteNum, err := c.netConn.Read(c.recvBuffer[c.recvBufferNowLen:])

		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("[%s] tcpconn read [%d] bytes", c.belongToService.name, byteNum))
		}

		if err != nil {
			if err == io.EOF {
				//对端关闭了连接
				c.CloseConnection()
				return
			}
			c.belongToService.OnServiceError(c.belongToService, err)
			return
		}

		c.recvBufferNowLen += uint64(byteNum)

		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("[%s] tcpconn recvBufferNowLen=[%d]", c.belongToService.name, c.recvBufferNowLen))
			fmt.Println(fmt.Sprintf("[%s] tcpconn recvBuffer:", c.belongToService.name))
			fmt.Println(string(c.recvBuffer[0:c.recvBufferNowLen]))
		}

		c.HandleBuffer()
	}
}

// HandleBuffer 处理缓冲区
func (c *TCPConnection) HandleBuffer() {
	copyBuffer := c.recvBuffer[0:c.recvBufferNowLen]
	for c.recvBufferNowLen > 0 {
		firstMsgLen, err := c.protocolHandler.FirstMsgLen(copyBuffer)
		if err != nil {
			//处理解析异常
			if c.protocolName == config.HTTPStr {
				handler := c.protocolHandler.(*http.Handler)
				switch handler.ParseStatus {
				case http.ParseStatusRecvBufferEmpty,
					http.ParseStatusNotHTTP,
					http.ParseStatusIncomplete:
					//继续接收
				case http.ParseStatusParseErr:
					//明显出错
					c.CloseConnection()
				}
			}
			break
		}
		//取出第 1 条完整的消息
		firstMsg := c.recvBuffer[0:firstMsgLen]

		switch c.protocolName {
		case config.HTTPStr:
			//HTTP 1.1 协议，短连接。
			//处理完一条消息后，理论上客户端会关闭tcp连接
			c.HandleHTTPMsg(firstMsg)
			c.belongToService.OnConnRequest(c)
		case config.StreamStr:
			//自定义 MsgBody 协议，长链接
			//处理完一条消息后，理论上客户端不会关闭tcp连接
			c.HandleStreamMsg(firstMsg)
			c.belongToService.OnConnRequest(c)
		case config.WebSocketStr:
			//WebSocket 协议，长链接
			//处理完一条消息后，理论上客户端不会关闭tcp连接
			err := c.HandleWebSocketMsg(firstMsg)
			if err != nil {
				c.CloseConnection()
				return
			}
			c.belongToService.OnConnRequest(c)
			//如果握手成功，就直接响应一个固定的测试消息
			t1p1protocol := c.protocolHandler.(*websocket.WebSocket)
			if t1p1protocol.IsHandshakeStatusYes() {
				t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", c.belongToService.name))
				c.SendMsg([]byte{})
			}
		}

		//处理接收缓冲区中剩余的数据
		c.recvBuffer = c.recvBuffer[firstMsgLen:]
		//recvBufferNowLen 是 uint64 类型的，做减法的时候小心溢出
		if c.recvBufferNowLen <= firstMsgLen {
			c.recvBufferNowLen = 0
			break
		} else {
			c.recvBufferNowLen -= firstMsgLen
		}
	}
}

func (c *TCPConnection) HandleHTTPMsg(sli1firstMsg []byte) {
	t1p1protocol := c.protocolHandler.(*http.Handler)
	t1p1protocol.Decode(sli1firstMsg)

	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.TCPConnection.HandelHTTPMsg.Decode: ", c.belongToService.name))
		fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
	}
}

func (c *TCPConnection) HandleStreamMsg(sli1firstMsg []byte) {
	t1p1protocol := c.protocolHandler.(*stream.Stream)
	t1p1protocol.Decode(sli1firstMsg)

	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.TCPConnection.HandelStreamMsg.Decode: ", c.belongToService.name))
		fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
	}
}

func (c *TCPConnection) HandleWebSocketMsg(sli1firstMsg []byte) error {
	t1p1protocol := c.protocolHandler.(*websocket.WebSocket)
	t1p1protocol.Decode(sli1firstMsg)

	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleWebSocketMsg.Decode: ", c.belongToService.name))
		fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
	}

	// 如果还没有握手成功，就走握手流程
	if t1p1protocol.IsHandshakeStatusNo() {
		sli1respMsg, err := t1p1protocol.CheckHandshakeReq()

		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleWebSocketMsg.CheckHandshakeReq: ", c.belongToService.name))
			fmt.Println(fmt.Sprintf("%+v", string(sli1respMsg)))
		}

		if err != nil {
			// 发送 400 给客户端，并且关闭连接
			resp := http.NewResponse()
			resp.SetStatusCode(http.StatusBadRequest)
			respStr := resp.MakeMsg(fmt.Sprintf("this is %s. handshake err: %s", c.belongToService.name, err))
			c.WriteData([]byte(respStr))

			return err
		} else {
			// 握手消息是通过 websocket.WebSocket 内部的 http.HandlerI9 处理的
			// 走 SendMsg 方法会判断成 WebSocket，走编码逻辑，所以这里通过 WriteData 方法直接发送
			err = c.WriteData(sli1respMsg)
			if nil == err {
				t1p1protocol.SetHandshakeStatusYes()
			}
		}
	}

	return nil
}

// SendMsg 发送数据
func (c *TCPConnection) SendMsg(sli1msg []byte) {
	switch c.protocolName {
	case config.HTTPStr:
		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", c.belongToService.name))
			fmt.Println(string(sli1msg))
		}
		c.WriteData(sli1msg)
	case config.StreamStr:
		t1sli1msg, _ := c.protocolHandler.Encode()
		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", c.belongToService.name))
			fmt.Println(string(t1sli1msg))
		}
		c.WriteData(t1sli1msg)
	}
}

// WriteData 发送数据
func (c *TCPConnection) WriteData(sli1data []byte) error {
	// net.Conn.Write，系统调用，用 socket 发送数据
	byteNum, err := c.netConn.Write(sli1data)

	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.TCPConnection.WriteData.byteNum: %d", c.belongToService.name, byteNum))
	}

	if err != nil {
		c.belongToService.OnServiceError(c.belongToService, err)
		c.CloseConnection()
	}

	if byteNum != len(sli1data) {
		return errors.New("write byte != data length")
	}
	return nil
}

// CloseConnection 关闭连接
func (c *TCPConnection) CloseConnection() {
	c.runStatus = config.RunStatusOff
	_ = c.netConn.Close()
	c.recvBufferNowLen = 0
	c.belongToService.DeleteConnection(c)
	c.belongToService.OnConnClose(c)
}
