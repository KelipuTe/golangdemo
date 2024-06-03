package websocket

import (
	"demo-golang/http"
	"log"
	"testing"
)

type TestServerHandler struct {
}

func NewTestServerHandler() *TestServerHandler {
	return &TestServerHandler{}
}

func (t *TestServerHandler) HandleMsg(req *Msg, conn *AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	data, _ := req.parseJson()
	if data["method"] == "/api/msg_only" {
		log.Println(data)
	} else if data["method"] == "/api/need_resp" {
		resp := NewUnmaskTextMsg()
		resp.Payload = "{\"method\":\"/api/msg_only\",\"msg\":\"server\"}"
		_ = conn.sendMsg(resp)
	}
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) HandleMsg(req *http.Request, resp *http.Response) {
	log.Println(req.Method)
	log.Println(req.Uri)
	log.Println(req.Version)
	if req.Method == http.MethodGet {
		log.Println(req.Query)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"name\":\"tom\"}"
	} else if req.Method == http.MethodPost {
		log.Println(req.Body)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"price\":100}"
	}
}

func Test_Server(t *testing.T) {
	h := NewTestServerHandler()
	s := NewServer("localhost", 9603, h)
	s.SupportHTTP(NewTestHandler())
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()
}
