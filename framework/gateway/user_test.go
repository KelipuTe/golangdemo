package gateway

import (
	"demo-golang/official/signal"
	"demo-golang/websocket"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type CaseUserHandler struct {
}

func NewCaseUserHandler() *CaseUserHandler {
	return &CaseUserHandler{}
}

func (t *CaseUserHandler) HandleMsg(req *websocket.Msg, conn *websocket.DialConn) {
	pkg := &Package{}
	_ = req.ParseJson(pkg)

	if pkg.Type == PackageTypeReq {
		switch pkg.Uri {
		case "/api/get_user":
			//{"type":1,"service":"user","uri":"/api/get_user","data":"1"}
			respPkg := &Package{
				To:   pkg.From,
				Type: PackageTypeResp,
				Data: fmt.Sprintf("{\"id\":%s,\"name\":\"name\"}", pkg.Data),
			}
			resp := websocket.NewUnmaskTextMsg()
			respPkgJson, _ := json.Marshal(respPkg)
			resp.Payload = string(respPkgJson)
			_ = conn.SendMsg(resp)
		case "/api/send_chat":
			//{"type":1,"service":"user","uri":"/api/send_chat","data":"chat data"}
			log.Println(pkg.Data)
		}
	}
}

func TestInnerUser(t *testing.T) {
	h := NewCaseUserHandler()
	c := websocket.NewClient("localhost", 9601, h)
	err := c.Start()
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	//注册
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
