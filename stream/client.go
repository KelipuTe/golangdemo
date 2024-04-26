package stream

import (
	"log"
	"net"
	"strconv"
)

// Client 客户端
type Client struct {
	ip   string
	port int
	conn *DialConn
}

func NewClient(ip string, port int) *Client {
	return &Client{
		ip:   ip,
		port: port,
	}
}

// Send 发送请求
func (t *Client) Send(req *Request) (*Response, error) {
	if t.conn == nil {
		addr := t.ip + ":" + strconv.Itoa(t.port)
		netConn, err := net.Dial("tcp4", addr) //发起连接
		if err != nil {
			return nil, err
		}

		t.connDial(netConn)
	}

	t.conn.SendReq(req)

	resp := NewResponse()
	t.conn.waitResp(resp)

	return resp, nil
}

// connDial 连接建立
func (t *Client) connDial(netConn net.Conn) *DialConn {
	log.Println("conn dial")

	dialConn := NewDialConn(t, netConn)

	t.conn = dialConn

	return dialConn
}

func (t *Client) connClose() {
	t.conn = nil
}

// CloseConn 关闭连接
func (t *Client) CloseConn() {
	if t.conn != nil {
		t.conn.close()
		t.conn = nil
	}
}
