package websocket

import (
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

func Test_Server(t *testing.T) {
	h := NewTestServerHandler()
	s := NewServer("localhost", 9603, h)
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()
}
