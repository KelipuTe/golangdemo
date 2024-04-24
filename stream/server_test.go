package stream

import (
	"log"
	"testing"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) HandleMsg(req *Request, resp *Response) {
	log.Println(req.MsgLen, req.Body)
	data, _ := req.parseJson()
	if data["method"] == "/api/user" {
		resp.Body = "{\"id\":1,\"name\":\"tom\"}"
	} else if data["method"] == "/api/order" {
		resp.Body = "{\"id\":1,\"price\":100}"
	}
}

func Test_Server(t *testing.T) {
	h := NewTestHandler()
	s := NewServer("localhost", 9602, h)
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
}
