package http

import (
	"log"
	"net"
	"strconv"
	"sync"
)

const (
	connPoolNumMax = 1024 //最大tcp连接数
)

type Handler interface {
	HandleHTTP(req *Request, resp *Response)
}

type Server struct {
	ip           string
	port         int
	listener     net.Listener
	connPool     map[string]*AcceptConn //连接上来的tcp
	connPoolNum  int                    //当前tcp连接数
	connPoolLock *sync.Mutex            //connPool的锁
	handler      Handler
}

func NewServer(ip string, port int, h Handler) *Server {
	return &Server{
		ip:           ip,
		port:         port,
		connPool:     make(map[string]*AcceptConn, connPoolNumMax),
		connPoolNum:  0,
		connPoolLock: &sync.Mutex{},
		handler:      h,
	}
}

func (t *Server) StartListen() error {
	addr := t.ip + ":" + strconv.Itoa(t.port)
	netListener, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	t.listener = netListener

	for {
		netConn, err := t.listener.Accept()
		if err != nil {
			return err
		}
		log.Println("conn accept")
		httpConn := t.acceptConn(netConn)

		go httpConn.handleConn()
	}
}

func (t *Server) acceptConn(netConn net.Conn) *AcceptConn {
	httpConn := NewAcceptConn(t, netConn)

	t.connPoolLock.Lock()
	defer t.connPoolLock.Unlock()

	t.connPoolNum++
	addr := httpConn.conn.RemoteAddr().String()
	t.connPool[addr] = httpConn

	return httpConn
}

func (t *Server) closeConn(c *AcceptConn) {
	t.connPoolLock.Lock()
	defer t.connPoolLock.Unlock()

	addr := c.conn.RemoteAddr().String()
	if _, ok := t.connPool[addr]; !ok {
		return
	}
	delete(t.connPool, addr)
	t.connPoolNum--
}
