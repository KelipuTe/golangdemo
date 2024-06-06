package gateway

import (
	"demo-golang/websocket"
	"log"
)

type InnerHandler struct {
	gateway *Gateway
}

func NewInnerHandler() *InnerHandler {
	return &InnerHandler{}
}

func (t *InnerHandler) HandleMsg(req *websocket.Msg, conn *websocket.AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	pkg := &Package{}
	_ = req.ParseJson(pkg)
	if pkg.Type == PackageTypeReq {
		//内部服务发来的请求
		switch pkg.Uri {
		case "/api/register":
			//注册
			t.gateway.innerPool[pkg.Data] = conn
		default:

		}
	} else {
		//内部服务发来的响应，回传给对应外部连接
		to, ok := t.gateway.openPool[pkg.From]
		if ok {
			toReq := websocket.NewUnmaskTextMsg()
			toReq.Payload = req.Payload
			_ = to.SendMsg(toReq)
		}
	}
}
