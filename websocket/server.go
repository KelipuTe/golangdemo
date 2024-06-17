package websocket

import (
	"demo-golang/http"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

// ServerHandler 处理消息的接口，需要外部实现
type ServerHandler interface {
	HandleMsg(req *Msg, conn *AcceptConn)
}

// Server 服务端
type Server struct {
	ip   string
	port int

	listener     net.Listener
	connPool     map[string]*AcceptConn //连接上来的tcp
	connPoolNum  int                    //当前tcp连接数
	connPoolLock *sync.Mutex            //connPool的锁

	handler ServerHandler //websocket处理接口

	httpHandler http.Handler //http处理接口

	onConn func(*AcceptConn) //ws连接事件

	isRun    bool          //是不是再跑
	needPing bool          //需不需要ping
	interval time.Duration //心跳间隔
}

func NewServer(port int, h ServerHandler) *Server {
	return &Server{
		ip:           "localhost",
		port:         port,
		connPool:     make(map[string]*AcceptConn, connPoolNumMax),
		connPoolNum:  0,
		connPoolLock: &sync.Mutex{},
		handler:      h,
		httpHandler:  nil,
		onConn:       nil,
		isRun:        true,
		needPing:     true,
	}
}

func (t *Server) SetHTTPHandler(h http.Handler) {
	t.httpHandler = h
}

func (t *Server) SetOnConn(f func(*AcceptConn)) {
	t.onConn = f
}

func (t *Server) SetNeedPing(b bool) {
	t.needPing = b
}

func (t *Server) SetInterval(i time.Duration) {
	t.interval = i
}

// Start 启动服务
func (t *Server) Start() error {
	addr := t.ip + ":" + strconv.Itoa(t.port)
	netListener, err := net.Listen("tcp4", addr) //开始监听
	if err != nil {
		return err
	}
	t.listener = netListener

	for t.isRun {
		netConn, err := t.listener.Accept() //等待客户端连接上来
		if err != nil {
			return err
		}

		acceptConn := t.connAccept(netConn)
		go acceptConn.handleMsg() //可以并发处理每个连接
	}

	return nil
}

// connAccept 连接建立
func (t *Server) connAccept(netConn net.Conn) *AcceptConn {
	log.Println("conn accept")

	acceptConn := NewAcceptConn(t, netConn)

	t.connPoolNum++
	addr := acceptConn.conn.RemoteAddr().String()
	t.connPool[addr] = acceptConn //这里没有并发

	return acceptConn
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

// Close 关闭服务
func (t *Server) Close() {
	_ = t.listener.Close()
	t.isRun = false
	for _, v := range t.connPool {
		v.close()
	}
}
