package websocket

import (
	"log"
	"testing"
	"time"
)

type TestClientHandler struct {
}

func NewTestClientHandler() *TestClientHandler {
	return &TestClientHandler{}
}

func (t *TestClientHandler) HandleMsg(req *Msg, conn *DialConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)

	data, _ := req.parseJson()
	if data["method"] == "/api/msg_only" {
		log.Println(data)
	}
}

func Test_Client(t *testing.T) {
	h := NewTestClientHandler()
	c := NewClient("localhost", 9601, h)
	err := c.Start()
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	for {
		req := NewMaskTextMsg()
		req.Payload = "{\"method\":\"/api/msg_only\",\"msg\":\"client\"}"
		err := c.Send(req)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(5 * time.Second)

		req2 := NewMaskTextMsg()
		req2.Payload = "{\"method\":\"/api/need_resp\",\"msg\":\"client\"}"
		err2 := c.Send(req2)
		if err2 != nil {
			t.Error(err2)
		}

		time.Sleep(5 * time.Second)
	}
}
