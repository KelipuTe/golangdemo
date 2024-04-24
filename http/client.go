package http

import (
	"log"
	"net"
	"strconv"
)

type Client struct {
	ip        string
	port      int
	conn      *DialConn
	keepAlive bool
}

func NewClient() *Client {
	return &Client{
		keepAlive: false,
	}
}

func (t *Client) Send(req *Request) (*Response, error) {
	err := t.parseDNS(req)
	if err != nil {
		return nil, err
	}
	t.keepAlive = req.isKeepAlive()

	if t.conn == nil {
		addr := t.ip + ":" + strconv.Itoa(t.port)
		netConn, err := net.Dial("tcp4", addr)
		if err != nil {
			return nil, err
		}
		log.Println("conn dial")
		t.dialConn(netConn)
	}

	t.conn.SendReq(req)

	resp := NewResponse()
	t.conn.waitResp(resp)

	if !t.keepAlive {
		t.CloseConn()
	}

	return resp, nil
}

func (t *Client) parseDNS(req *Request) error {
	t.ip = "localhost"
	t.port = 9601
	return nil
}

func (t *Client) dialConn(netConn net.Conn) *DialConn {
	httpConn := NewDialConn(t, netConn)
	t.conn = httpConn
	return httpConn
}

func (t *Client) CloseConn() {
	if t.conn != nil {
		t.conn.close()
		t.conn = nil
	}
}
