package client

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/protocol/http"
	"demo-golang/tcp-service/protocol/stream"
	"demo-golang/tcp-service/protocol/websocket"
	"fmt"
	"testing"
)

func Test_Client_HTTP(t *testing.T) {
	client := NewTCPClient("127.0.0.1", 9501, config.HTTPStr)
	client.SetName(fmt.Sprintf("%s-client", config.HTTPStr))
	client.OpenDebug()
	client.AfterConnConnect = func(conn *TCPConnection) {
		if conn.IsDebug() {
			req := http.NewRequest()
			req.SetMethod(http.MethodGet)
			msg := req.MakeMsg(fmt.Sprintf("this is %s.", conn.belongToClient.name))
			conn.SendMsg([]byte(msg))
		}
	}
	client.Start()
}

func Test_Client_Stream(t *testing.T) {
	client := NewTCPClient("127.0.0.1", 9502, config.StreamStr)
	client.SetName(fmt.Sprintf("%s-client", config.StreamStr))
	client.OpenDebug()
	client.AfterConnConnect = func(conn *TCPConnection) {
		if conn.IsDebug() {
			handler := conn.protocolHandler.(*stream.Stream)
			handler.SetDecodeMsg(fmt.Sprintf("this is %s.", conn.belongToClient.name))
			msg, _ := conn.protocolHandler.Encode()
			conn.SendMsg(msg)
		}
	}
	client.Start()
}

func Test_Client_WebSocket(t *testing.T) {
	client := NewTCPClient("127.0.0.1", 9503, config.WebSocketStr)
	client.SetName(fmt.Sprintf("%s-client", config.WebSocketStr))
	client.OpenDebug()
	client.AfterConnConnect = func(conn *TCPConnection) {
		//发送握手消息
		handler := conn.protocolHandler.(*websocket.WebSocket)
		respMsg, _ := handler.MakeHandShakeReq()
		conn.SendMsg(respMsg)
	}
	client.Start()
}
