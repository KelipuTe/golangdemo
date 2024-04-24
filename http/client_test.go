package http

import (
	"log"
	"testing"
	"time"
)

func Test_Client_Get(t *testing.T) {
	c := NewClient()
	req := NewRequest()
	req.Method = MethodGet
	req.Uri = "/api/user"
	req.Query["id"] = "1"
	resp, err := c.Send(req)
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}

func Test_Client_Post(t *testing.T) {
	c := NewClient()
	req := NewRequest()
	req.Method = MethodPost
	req.Uri = "/api/order"
	req.Body = "{\"id\":1}"
	resp, err := c.Send(req)
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.Body)
}

func Test_Client_KeepAliveOn(t *testing.T) {
	c := NewClient()

	req := NewRequest()
	req.Method = MethodGet
	req.Uri = "/api/user"
	req.Query["id"] = "1"
	req.KeepAliveOn()
	resp, err := c.Send(req)
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.Body)

	time.Sleep(time.Second)

	req2 := NewRequest()
	req2.Method = MethodPost
	req2.Uri = "/api/order"
	req2.Body = "{\"id\":1}"
	resp2, err2 := c.Send(req2)
	if err2 != nil {
		t.Error(err2)
	}
	log.Println(resp2.Body)
}

func Test_Client_KeepAliveOff(t *testing.T) {
	c := NewClient()

	req := NewRequest()
	req.Method = MethodGet
	req.Uri = "/api/user"
	req.Query["id"] = "1"
	resp, err := c.Send(req)
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.Body)

	time.Sleep(time.Second)

	req2 := NewRequest()
	req2.Method = MethodPost
	req2.Uri = "/api/order"
	req2.Body = "{\"id\":1}"
	resp2, err2 := c.Send(req2)
	if err2 != nil {
		t.Error(err2)
	}
	log.Println(resp2.Body)
}
