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
	runStatus config.RunStatus //连接状态

	belongToClient  *TCPClient    //连接所属的客户端
	protocolName    string        //协议名称
	protocolHandler abs.HandlerI9 //协议处理器

	netConn net.Conn //tcp连接本体

	recvBuffer       []byte //接收缓冲区
	recvBufferMaxLen uint64 //接收缓冲区最大大小
	recvBufferNowLen uint64 //接收缓冲区当前大小
}

func NewTCPConnection(client *TCPClient, conn net.Conn) *TCPConnection {
	return &TCPConnection{
		runStatus: config.RunStatusOn,

		belongToClient:  client,
		protocolName:    client.protocolName,
		protocolHandler: protocol.NewHandler(client.protocolName),

		netConn: conn,

		recvBuffer:       make([]byte, config.RecvBufferMaxLen),
		recvBufferMaxLen: config.RecvBufferMaxLen,
		recvBufferNowLen: 0,
	}
}

func (t *TCPConnection) IsRun() bool {
	return t.runStatus == config.RunStatusOn
}

func (t *TCPConnection) IsDebug() bool {
	return t.belongToClient.IsDebug()
}

func (t *TCPConnection) GetClient() *TCPClient {
	return t.belongToClient
}

func (t *TCPConnection) GetProtocolName() string {
	return t.protocolName
}

func (t *TCPConnection) GetProtocolHandler() abs.HandlerI9 {
	return t.protocolHandler
}

func (t *TCPConnection) GetNetConnRemoteAddr() string {
	return t.netConn.RemoteAddr().String()
}

func (t *TCPConnection) HandleConnection(deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for t.IsRun() {
		byteNum, err := t.netConn.Read(t.recvBuffer[t.recvBufferNowLen:])

		if t.IsDebug() {
			fmt.Println(fmt.Sprintf("client [%s] read [%d] bytes", t.belongToClient.name, byteNum))
		}

		if err != nil {
			if err == io.EOF {
				t.CloseConnection()
				return
			}
			t.belongToClient.OnClientError(t.belongToClient, err)
			return
		}

		t.recvBufferNowLen += uint64(byteNum)

		if t.IsDebug() {
			fmt.Println(fmt.Sprintf("client [%s] recvBuffer:", t.belongToClient.name))
			fmt.Println(string(t.recvBuffer[0:t.recvBufferNowLen]))
		}

		t.HandleBuffer()
	}
}

func (t *TCPConnection) HandleBuffer() {
	copyBuffer := t.recvBuffer[0:t.recvBufferNowLen]
	for t.recvBufferNowLen > 0 {
		firstMsgLen, err := t.protocolHandler.FirstMsgLen(copyBuffer)
		if err != nil {
			//处理解析异常
			if t.protocolName == config.HTTPStr {
				handler := t.protocolHandler.(*http.Handler)
				switch handler.ParseStatus {
				case http.ParseStatusRecvBufferEmpty,
					http.ParseStatusNotHTTP,
					http.ParseStatusIncomplete:
					//继续接收
				case http.ParseStatusParseErr:
					//明显出错
					t.CloseConnection()
				}
			}
			break
		}
		firstMsg := t.recvBuffer[0:firstMsgLen]

		switch t.protocolName {
		case config.HTTPStr:
			t.HandleHTTPMsg(firstMsg)
			t.belongToClient.OnConnGetRequest(t)
		case config.StreamStr:
			// 自定义 Data 协议的消息，解析之后由外部实现的 OnConnGetRequest 继续处理
			t1p1protocol := t.protocolHandler.(*stream.Stream)
			t1p1protocol.Decode(firstMsg)

			if t.IsDebug() {
				fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.StreamStr.Decode: ", t.belongToClient.name))
				fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
			}
			t.belongToClient.OnConnGetRequest(t)
		case config.WebSocketStr:
			// WebSocket 协议的消息，需要判断是握手消息还是测试消息
			t1p1protocol := t.protocolHandler.(*websocket.WebSocket)
			t1p1protocol.Decode(firstMsg)

			if t1p1protocol.IsHandshakeStatusNo() {
				// 握手消息，校验一下服务端响应的握手消息
				err = t1p1protocol.CheckHandShakeResp()
				if err == nil {
					t1p1protocol.SetHandshakeStatusYes()
					t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", t.belongToClient.name))
					t.SendMsg([]byte{})
				} else {
					t.CloseConnection()
				}
			} else {
				// 测试消息，解析之后直接输出
				if t.IsDebug() {
					fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.WebSocketStr.Decode: ", t.belongToClient.name))
					fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
					t.belongToClient.OnConnGetRequest(t)
				}
			}
		}

		t.recvBuffer = t.recvBuffer[firstMsgLen:]
		// recvBufferNowLen 是 uint64 类型的，做减法的时候小心溢出
		if t.recvBufferNowLen <= firstMsgLen {
			t.recvBufferNowLen = 0
			break
		} else {
			t.recvBufferNowLen -= firstMsgLen
		}
	}
}

func (t *TCPConnection) HandleHTTPMsg(sli1firstMsg []byte) {
	t1p1protocol := t.protocolHandler.(*http.Handler)
	t1p1protocol.Decode(sli1firstMsg)

	if t.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.TCPConnection.HandelHTTPMsg.Decode: ", t.belongToClient.name))
		fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
	}
}

func (t *TCPConnection) SendMsg(msg []byte) {
	if t.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] send:", t.belongToClient.name))
		fmt.Println(string(msg))
	}
	_ = t.WriteData(msg)
}

func (t *TCPConnection) WriteData(data []byte) error {
	byteNum, err := t.netConn.Write(data)

	if t.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] write [%d] bytes data.", t.belongToClient.name, byteNum))
	}

	if err != nil {
		t.belongToClient.OnClientError(t.belongToClient, err)
		t.CloseConnection()
		return err
	}

	if byteNum != len(data) {
		return errors.New("write bytes != data len")
	}
	return nil
}

func (t *TCPConnection) CloseConnection() {
	t.runStatus = config.RunStatusOff
	_ = t.netConn.Close()
	t.recvBufferNowLen = 0
	t.belongToClient.AfterConnClose(t)
}
