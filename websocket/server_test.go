package websocket

import (
	"log"
	"testing"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) HandleMsg(req *Request, conn *AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)
	data, _ := req.parseJson()
	resp := NewResponse()

	if data["method"] == "/api/user" {
		resp.Payload = "{\"id\":1,\"name\":\"tom\"}"
	} else if data["method"] == "/api/order" {
		resp.Payload = "{\"id\":1,\"price\":100}"
	} else {
		resp.Payload = "123"
	}

	conn.sendResp(resp)
}

// 可以用这个网站测试 http://www.websocket-test.com/
func Test_Server(t *testing.T) {
	h := NewTestHandler()
	s := NewServer("localhost", 9603, h)
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
}
