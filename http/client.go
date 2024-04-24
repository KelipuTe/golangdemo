package http

import (
	"net"
	"strconv"
)

type Client struct {
	ip   string
	port int
	conn *DialConn
}

func NewClient() *Client {
	return &Client{}
}

func (t *Client) Send(req *Request) (*Response, error) {
	err := t.parseDNS(req)
	if err != nil {
		return nil, err
	}

	addr := t.ip + ":" + strconv.Itoa(t.port)
	netConn, err := net.Dial("tcp4", addr)
	if err != nil {
		return nil, err
	}

	httpConn := t.dialConn(netConn)
	httpConn.SendReq(req)

	resp := NewResponse()
	httpConn.waitResp(resp)

	return resp, nil
}

func (t *Client) parseDNS(req *Request) error {
	t.ip = "localhost"
	t.port = 9601
	return nil
}

func (t *Client) dialConn(netConn net.Conn) *DialConn {
	httpConn := NewDialConn(t, netConn)
	return httpConn
}
