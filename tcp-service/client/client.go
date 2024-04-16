package client

import (
	"demo-golang/tcp-service/config"
	"net"
	"strconv"
	"sync"

	pkgErrors "github.com/pkg/errors"
)

// TCPClient tcp客户端，负责接收外部请求并转发给gateway
type TCPClient struct {
	name         string             //客户端名称
	runStatus    config.RunStatus   //运行状态
	debugStatus  config.DebugStatus //debug 开关状态
	protocolName string             //使用的协议
	address      string             //ip地址
	port         uint16             //端口号

	conn          *TCPConnection          //tcp连接
	OnClientStart func(*TCPClient)        //客户端启动事件回调
	OnClientError func(*TCPClient, error) //客户端错误事件回调
	OnConnConnect func(*TCPConnection)    //TCP 连接，连接事件回调
	OnConnRequest func(*TCPConnection)    //TCP 连接，请求事件回调
	OnConnClose   func(*TCPConnection)    //TCP 连接，关闭事件回调
}

func NewTCPClient(protocolName string, address string, port uint16) *TCPClient {
	return &TCPClient{
		name:         defaultName,
		runStatus:    config.RunStatusOn,
		debugStatus:  config.DebugStatusOff,
		protocolName: protocolName,
		address:      address,
		port:         port,

		OnClientStart: defaultOnClientStart,
		OnClientError: defaultOnClientError,
		OnConnConnect: defaultOnConnConnect,
		OnConnRequest: defaultOnConnRequest,
		OnConnClose:   defaultOnConnClose,
	}
}

// SetName 设置客户端名称
func (p1this *TCPClient) SetName(name string) {
	p1this.name = name
}

// GetName 获取客户端名称
func (p1this *TCPClient) GetName() string {
	return p1this.name
}

// SetDebugStatusOn 打开 debug
func (p1this *TCPClient) SetDebugStatusOn() {
	p1this.debugStatus = config.DebugStatusOn
}

// IsDebug 是否是 debug 模式
func (p1this *TCPClient) IsDebug() bool {
	return p1this.debugStatus == config.DebugStatusOn
}

// GetTCPConn 获取 TCP 客户端内部的 TCP 连接
func (p1this *TCPClient) GetTCPConn() *TCPConnection {
	return p1this.conn
}

func (p1this *TCPClient) Start() {
	p1this.OnClientStart(p1this)

	p1conn, err := net.Dial("tcp4", p1this.address+":"+strconv.Itoa(int(p1this.port)))
	if nil != err {
		p1this.OnClientError(p1this, pkgErrors.WithMessage(err, "TCPClient.StartListen"))
		return
	}

	p1this.conn = NewTCPConnection(p1this, p1conn)
	p1this.OnConnConnect(p1this.conn)

	var wg sync.WaitGroup
	wg.Add(1)
	go p1this.conn.HandleConnection(wg.Done)
	wg.Wait()
}
