package gateway

import (
	"demo-golang/signal"
	"demo-golang/websocket"
	"encoding/json"
	"log"
	"testing"
)

type TestUserHandler struct {
}

func NewTestUserHandler() *TestUserHandler {
	return &TestUserHandler{}
}

func (t *TestUserHandler) HandleMsg(req *websocket.Msg, conn *websocket.DialConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	pkg := &Package{}
	_ = req.ParseJson(pkg)

	if pkg.Type == PackageTypeReq {
		switch pkg.Uri {
		case "/api/msg_only":
			log.Println(pkg.Data)
		case "/api/need_resp":
			respPkg := &Package{
				From: pkg.From,
				Type: PackageTypeResp,
				Data: "{\"method\":\"/api/msg_only\",\"msg\":\"user\"}",
			}
			resp := websocket.NewUnmaskTextMsg()
			respPkgJson, _ := json.Marshal(respPkg)
			resp.Payload = string(respPkgJson)
			_ = conn.SendMsg(resp)
		}
	}
}

func TestInnerUser(t *testing.T) {
	h := NewTestUserHandler()
	c := websocket.NewClient("localhost", 9601, h)
	err := c.Start()
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	req := websocket.NewMaskTextMsg()
	pkg := &Package{
		Type:    PackageTypeReq,
		Service: "gateway",
		Uri:     "/api/register",
		Data:    "user",
	}
	pkgJson, _ := json.Marshal(pkg)
	req.Payload = string(pkgJson)
	_ = c.Send(req)

	signal.WaitForSIGINT()
}
