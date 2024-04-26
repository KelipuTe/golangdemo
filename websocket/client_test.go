package websocket

import (
	"log"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	c := NewClient("localhost", 9603)
	defer c.CloseConn()

	//for {
	req := NewRequest()
	msg := "123"
	req.Payload = msg
	resp, err := c.Send(req)
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.Payload)

	time.Sleep(time.Second)
	//}
}
