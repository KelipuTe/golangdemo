package stream

import (
	"log"
	"net"
	"testing"
	"time"
)

func Test_Dial_Close(t *testing.T) {
	addr := "localhost:9602"
	netConn, _ := net.Dial("tcp4", addr) //发起连接
	_ = netConn.Close()
}

func Test_Client(t *testing.T) {
	c := NewClient("localhost", 9602)
	defer c.CloseConn()

	for {
		req := NewRequest()
		msg := "{\"method\":\"/api/user\",\"id\":1}"
		req.Body = msg
		resp, err := c.Send(req)
		if err != nil {
			t.Error(err)
		}
		log.Println(resp.Body)

		time.Sleep(time.Second)

		req2 := NewRequest()
		msg2 := "{\"method\":\"/api/order\",\"id\":1}"
		req2.Body = msg2
		resp2, err2 := c.Send(req2)
		if err2 != nil {
			t.Error(err2)
		}
		log.Println(resp2.Body)

		time.Sleep(time.Second)
	}
}
