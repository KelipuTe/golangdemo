package websocket

import (
	http2 "demo-golang/protocol/http"
	"log"
	"testing"
)

type TestWSHandler struct {
}

func NewTestWSHandler() *TestWSHandler {
	return &TestWSHandler{}
}

func (t *TestWSHandler) HandleMsg(req *Msg, conn *AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	data := make(map[string]any)
	_ = req.ParseJson(&data)
	if data["method"] == "/api/msg_only" {
		log.Println(data)
	} else if data["method"] == "/api/need_resp" {
		resp := NewUnmaskTextMsg()
		resp.Payload = "{\"method\":\"/api/msg_only\",\"msg\":\"server\"}"
		_ = conn.SendMsg(resp)
	}
}

type TestHTTPHandler struct {
}

func NewTestHTTPHandler() *TestHTTPHandler {
	return &TestHTTPHandler{}
}

func (t *TestHTTPHandler) HandleMsg(req *http2.Request, resp *http2.Response) {
	log.Println(req.Method)
	log.Println(req.Uri)
	log.Println(req.Version)
	if req.Method == http2.MethodGet {
		log.Println(req.Query)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"name\":\"tom\"}"
	} else if req.Method == http2.MethodPost {
		log.Println(req.Body)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"price\":100}"
	}
}

func TestServer(t *testing.T) {
	h := NewTestWSHandler()
	s := NewServer(9601, h)
	s.SetHTTPHandler(NewTestHTTPHandler())
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()
}
