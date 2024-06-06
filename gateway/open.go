package gateway

import (
	"demo-golang/websocket"
	"encoding/json"
	"log"
)

type OpenHandler struct {
	gateway *Gateway
}

func NewOpenHandler() *OpenHandler {
	return &OpenHandler{}
}

func (t *OpenHandler) HandleMsg(req *websocket.Msg, conn *websocket.AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	pkg := &Package{}
	_ = req.ParseJson(pkg)
	pkg.From = conn.GetRemoteAddr()
	if pkg.Type == PackageTypeReq {
		//外部服务发来的请求，转发给指定内部服务
		to, ok := t.gateway.innerPool[pkg.Service]
		if ok {
			t.gateway.openPool[pkg.From] = conn
			toReq := websocket.NewUnmaskTextMsg()
			pkgJson, _ := json.Marshal(pkg)
			toReq.Payload = string(pkgJson)
			_ = to.SendMsg(toReq)
		}
	} else {
		//外部服务发来的响应
	}
}
