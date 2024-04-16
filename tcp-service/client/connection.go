package client

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
// 这个和 service 那里的那个 TCPConnection 很像
type TCPConnection struct {
	runStatus      config.RunStatus //连接状态
	belongToClient *TCPClient       //连接所属的客户端

	protocolName    string        //协议名称
	protocolHandler abs.HandlerI9 //协议处理器

	netConn          net.Conn //tcp连接本体
	recvBuffer       []byte   //接收缓冲区
	recvBufferMaxLen uint64   //接收缓冲区最大大小
	recvBufferNowLen uint64   //接收缓冲区当前大小
}

func NewTCPConnection(client *TCPClient, conn net.Conn) *TCPConnection {
	return &TCPConnection{
		runStatus:      config.RunStatusOn,
		belongToClient: client,

		protocolName:    client.protocolName,
		protocolHandler: protocol.NewHandler(client.protocolName),

		netConn:          conn,
		recvBuffer:       make([]byte, config.RecvBufferMaxLen),
		recvBufferMaxLen: config.RecvBufferMaxLen,
		recvBufferNowLen: 0,
	}
}

func (c *TCPConnection) IsRun() bool {
	return config.RunStatusOn == c.runStatus
}

func (c *TCPConnection) IsDebug() bool {
	return c.belongToClient.IsDebug()
}

func (c *TCPConnection) GetClient() *TCPClient {
	return c.belongToClient
}

func (c *TCPConnection) GetProtocolName() string {
	return c.protocolName
}

func (c *TCPConnection) GetProtocolHandler() abs.HandlerI9 {
	return c.protocolHandler
}

func (c *TCPConnection) GetNetConnRemoteAddr() string {
	return c.netConn.RemoteAddr().String()
}

func (c *TCPConnection) HandleConnection(deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for c.IsRun() {
		byteNum, err := c.netConn.Read(c.recvBuffer[c.recvBufferNowLen:])

		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.byteNum: %d", c.belongToClient.name, byteNum))
		}

		if err != nil {
			if err == io.EOF {
				c.CloseConnection()
				return
			}
			c.belongToClient.OnClientError(c.belongToClient, err)
			return
		}

		c.recvBufferNowLen += uint64(byteNum)

		if c.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.recvBufferNowLen: %d", c.belongToClient.name, c.recvBufferNowLen))
			fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.recvBuffer:", c.belongToClient.name))
			fmt.Println(string(c.recvBuffer[0:c.recvBufferNowLen]))
		}

		c.HandleBuffer()
	}
}

func (c *TCPConnection) HandleBuffer() {
	sli1Copy := c.recvBuffer[0:c.recvBufferNowLen]
	for c.recvBufferNowLen > 0 {
		firstMsgLength, err := c.protocolHandler.FirstMsgLen(sli1Copy)
		sli1firstMsg := c.recvBuffer[0:firstMsgLength]

		switch c.protocolName {
		case config.HTTPStr:
			// HandlerI9 1.1 协议的消息，解析之后由外部实现的 OnConnRequest 继续处理
			t1p1protocol := c.protocolHandler.(*http.Handler)
			t1p1protocol.Decode(sli1firstMsg)

			if c.IsDebug() {
				fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.StreamStr.Decode: ", c.belongToClient.name))
				fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
			}
			c.belongToClient.OnConnRequest(c)
		case config.StreamStr:
			// 自定义 Stream 协议的消息，解析之后由外部实现的 OnConnRequest 继续处理
			t1p1protocol := c.protocolHandler.(*stream.Stream)
			t1p1protocol.Decode(sli1firstMsg)

			if c.IsDebug() {
				fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.StreamStr.Decode: ", c.belongToClient.name))
				fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
			}
			c.belongToClient.OnConnRequest(c)
		case config.WebSocketStr:
			// WebSocket 协议的消息，需要判断是握手消息还是测试消息
			t1p1protocol := c.protocolHandler.(*websocket.WebSocket)
			t1p1protocol.Decode(sli1firstMsg)

			if t1p1protocol.IsHandshakeStatusNo() {
				// 握手消息，校验一下服务端响应的握手消息
				err = t1p1protocol.CheckHandShakeResp()
				if err == nil {
					t1p1protocol.SetHandshakeStatusYes()
					t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", c.belongToClient.name))
					c.SendMsg([]byte{})
				} else {
					c.CloseConnection()
				}
			} else {
				// 测试消息，解析之后直接输出
				if c.IsDebug() {
					fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.WebSocketStr.Decode: ", c.belongToClient.name))
					fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
					c.belongToClient.OnConnRequest(c)
				}
			}
		}

		c.recvBuffer = c.recvBuffer[firstMsgLength:]
		// recvBufferNowLen 是 uint64 类型的，做减法的时候小心溢出
		if c.recvBufferNowLen <= firstMsgLength {
			c.recvBufferNowLen = 0
			break
		} else {
			c.recvBufferNowLen -= firstMsgLength
		}
	}
}

func (c *TCPConnection) SendMsg(msg []byte) {
	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] send:", c.belongToClient.name))
		fmt.Println(string(msg))
	}
	_ = c.WriteData(msg)
}

func (c *TCPConnection) WriteData(data []byte) error {
	byteNum, err := c.netConn.Write(data)

	if c.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] write [%d] bytes data.", c.belongToClient.name, byteNum))
	}

	if err != nil {
		c.belongToClient.OnClientError(c.belongToClient, err)
		c.CloseConnection()
		return err
	}

	if byteNum != len(data) {
		return errors.New("write bytes != data len")
	}
	return nil
}

func (c *TCPConnection) CloseConnection() {
	c.runStatus = config.RunStatusOff
	_ = c.netConn.Close()
	c.recvBufferNowLen = 0
	c.belongToClient.OnConnClose(c)
}
