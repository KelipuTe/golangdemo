package http

import (
	"log"
	"net"
	"strconv"
)

// Client 客户端
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

// Send 发送请求
func (t *Client) Send(req *Request) (*Response, error) {
	err := t.parseDNS(req)
	if err != nil {
		return nil, err
	}
	t.keepAlive = req.isKeepAlive()

	if t.conn == nil {
		addr := t.ip + ":" + strconv.Itoa(t.port)
		netConn, err := net.Dial("tcp4", addr) //发起连接
		if err != nil {
			return nil, err
		}
		log.Println("conn dial")
		t.connDial(netConn)
	}

	t.conn.SendReq(req)

	resp := NewResponse()
	t.conn.waitResp(resp)

	// 如果不是长连接，则关闭连接
	if !t.keepAlive {
		t.CloseConn()
	}

	return resp, nil
}

// parseDNS 解析DNS
func (t *Client) parseDNS(req *Request) error {
	t.ip = "localhost"
	t.port = 9601
	return nil
}

// connDial 连接建立
func (t *Client) connDial(netConn net.Conn) *DialConn {
	httpConn := NewDialConn(t, netConn)
	t.conn = httpConn
	return httpConn
}

// CloseConn 关闭连接
func (t *Client) CloseConn() {
	if t.conn != nil {
		t.conn.close()
		t.conn = nil
	}
}
