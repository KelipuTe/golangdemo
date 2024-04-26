package http

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
	log.Println(req.Method)
	log.Println(req.Uri)
	log.Println(req.Version)
	if req.Method == MethodGet {
		log.Println(req.Query)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"name\":\"tom\"}"
	} else if req.Method == MethodPost {
		log.Println(req.Body)
		resp.Status = 200
		resp.Body = "{\"id\":1,\"price\":100}"
	}
}

func Test_Server(t *testing.T) {
	h := NewTestHandler()
	s := NewServer("localhost", 9601, h)
	err := s.Start()
	if err != nil {
		t.Error(err)
	}
}
