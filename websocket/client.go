package websocket

import (
	"log"
	"net"
	"strconv"
)

// ClientHandler 处理消息的接口，需要外部实现
type ClientHandler interface {
	HandleMsg(req *Msg, conn *DialConn)
}

// Client 客户端
type Client struct {
	ip   string
	port int

	conn *DialConn

	handler ClientHandler
}

func NewClient(ip string, port int, h ClientHandler) *Client {
	client := &Client{
		ip:      ip,
		port:    port,
		handler: h,
	}

	return client
}

// Start 启动客户端
func (t *Client) Start() error {
	addr := t.ip + ":" + strconv.Itoa(t.port)
	netConn, err := net.Dial("tcp4", addr) //发起连接
	if err != nil {
		return err
	}

	t.connDial(netConn)

	err = t.conn.handshake()
	if err != nil {
		return err
	}

	go t.conn.handleMsg()

	return nil
}

// connDial 连接建立
func (t *Client) connDial(netConn net.Conn) *DialConn {
	log.Println("conn dial")

	httpConn := NewDialConn(t, netConn)

	t.conn = httpConn

	return httpConn
}

// connClose 连接关闭
func (t *Client) connClose() {
	t.conn = nil
}

// Send 发送消息
func (t *Client) Send(req *Msg) error {
	return t.conn.sendMsg(req)
}

// Close 关闭连接
func (t *Client) Close() {
	if t.conn != nil {
		t.conn.close()
		t.conn = nil
	}
}
