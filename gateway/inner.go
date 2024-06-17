package gateway

import (
	"demo-golang/websocket"
)

type InnerHandler struct {
	gateway *Gateway
}

func NewInnerHandler() *InnerHandler {
	return &InnerHandler{}
}

func (t *InnerHandler) HandleMsg(req *websocket.Msg, conn *websocket.AcceptConn) {
	pkg := &Package{}
	_ = req.ParseJson(pkg)

	if pkg.Type == PackageTypeReq {
		//内部服务发来的请求
		switch pkg.Uri {
		case "/api/register": //注册
			t.gateway.innerPool[pkg.Data] = conn
		default:

		}
	} else {
		//内部服务发来的响应，转发给对应外部连接
		if to, ok := t.gateway.openPool[pkg.To]; ok {
			toReq := websocket.NewUnmaskTextMsg()
			toReq.Payload = pkg.Data
			_ = to.SendMsg(toReq)
		}
	}
}
