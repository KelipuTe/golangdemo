package gateway

import (
	"demo-golang/official/signal"
	"demo-golang/websocket"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type TestOrderHandler struct {
}

func NewTestOrderHandler() *TestOrderHandler {
	return &TestOrderHandler{}
}

func (t *TestOrderHandler) HandleMsg(req *websocket.Msg, conn *websocket.DialConn) {
	pkg := &Package{}
	_ = req.ParseJson(pkg)

	if pkg.Type == PackageTypeReq {
		switch pkg.Uri {
		case "/api/get_order":
			//{"type":1,"service":"order","uri":"/api/get_order","data":"1"}
			respPkg := &Package{
				To:   pkg.From,
				Type: PackageTypeResp,
				Data: fmt.Sprintf("{\"id\":%s,\"price\":100}", pkg.Data),
			}
			resp := websocket.NewUnmaskTextMsg()
			respPkgJson, _ := json.Marshal(respPkg)
			resp.Payload = string(respPkgJson)
			_ = conn.SendMsg(resp)
		case "/api/send_push":
			//{"type":1,"service":"order","uri":"/api/send_push","data":"push data"}
			log.Println(pkg.Data)
		}
	}
}

func TestInnerOrder(t *testing.T) {
	h := NewTestOrderHandler()
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
		Data:    "order",
	}
	pkgJson, _ := json.Marshal(pkg)
	req.Payload = string(pkgJson)
	_ = c.Send(req)

	signal.WaitForSIGINT()
}
