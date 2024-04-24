package stream

import (
	"log"
	"net"
	"strconv"
	"sync"
)

const (
	connPoolNumMax = 1024 //最大tcp连接数
)

// Handler 处理请求的接口，需要外部实现
type Handler interface {
	HandleMsg(req *Request, resp *Response)
}

// Server 服务端
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

// Start 启动服务
func (t *Server) Start() error {
	addr := t.ip + ":" + strconv.Itoa(t.port)
	netListener, err := net.Listen("tcp4", addr) //开始监听
	if err != nil {
		return err
	}
	t.listener = netListener

	for {
		netConn, err := t.listener.Accept() //等待连接
		if err != nil {
			return err
		}

		httpConn := t.connAccept(netConn)
		go httpConn.handleMsg() //可以并发处理每个连接
	}
}

// connAccept 连接接受
func (t *Server) connAccept(netConn net.Conn) *AcceptConn {
	log.Println("conn accept")

	httpConn := NewAcceptConn(t, netConn)

	t.connPoolNum++
	addr := httpConn.conn.RemoteAddr().String()
	t.connPool[addr] = httpConn //这里没有并发

	return httpConn
}

// connClose 连接关闭
func (t *Server) connClose(c *AcceptConn) {
	t.connPoolLock.Lock()
	defer t.connPoolLock.Unlock()

	addr := c.conn.RemoteAddr().String()
	if _, ok := t.connPool[addr]; !ok {
		return
	}
	delete(t.connPool, addr) //这里有并发
	t.connPoolNum--
}
