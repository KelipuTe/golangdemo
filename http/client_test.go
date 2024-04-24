package http

import (
	"log"
	"testing"
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
