package client

import (
	"demo-golang/tcp-service/config"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"

	pkgErrors "github.com/pkg/errors"
)

// TCPClient 客户端
type TCPClient struct {
	runStatus   config.RunStatus   //运行状态
	debugStatus config.DebugStatus //debug 开关状态

	name         string //客户端名称
	address      string //ip地址
	port         uint16 //端口号
	protocolName string //使用的协议

	OnClientError func(*TCPClient, error) //客户端，发生错误时的事件回调

	conn *TCPConnection //tcp连接本体

	AfterConnConnect func(*TCPConnection) //tcp连接，连接上之后的事件回调
	OnConnGetRequest func(*TCPConnection) //tcp连接，接到请求时的事件回调
	AfterConnClose   func(*TCPConnection) //tcp连接，连接关闭后的事件回调
}

func NewTCPClient(address string, port uint16, protocolName string) *TCPClient {
	return &TCPClient{
		runStatus:   config.RunStatusOff,
		debugStatus: config.DebugStatusOff,

		name:         defaultName,
		address:      address,
		port:         port,
		protocolName: protocolName,

		OnClientError: defaultOnClientError,

		conn: nil,

		AfterConnConnect: defaultAfterConnConnect,
		OnConnGetRequest: defaultOnConnGetRequest,
		AfterConnClose:   defaultAfterConnClose,
	}
}

func (c *TCPClient) GetName() string {
	return c.name
}

func (c *TCPClient) SetName(name string) {
	c.name = name
}

func (c *TCPClient) OpenDebug() {
	c.debugStatus = config.DebugStatusOn
}

func (c *TCPClient) IsDebug() bool {
	return c.debugStatus == config.DebugStatusOn
}

// GetTCPConn 获取tcp连接本体
func (c *TCPClient) GetTCPConn() *TCPConnection {
	return c.conn
}

func (c *TCPClient) Start() {
	c.StartInfo()

	conn, err := net.Dial("tcp4", c.address+":"+strconv.Itoa(int(c.port)))
	if err != nil {
		c.OnClientError(c, pkgErrors.WithMessage(err, "TCPClient.Start"))
		return
	}

	c.conn = NewTCPConnection(c, conn)
	c.runStatus = config.RunStatusOn
	c.AfterConnConnect(c.conn)

	var wg sync.WaitGroup
	wg.Add(1)
	go c.conn.HandleConnection(wg.Done)
	wg.Wait()
}

// StartInfo 输出服务的配置和环境参数
func (c *TCPClient) StartInfo() {
	log.Println("version: ", config.Version)
	log.Println("runtime.GOOS=", runtime.GOOS)
	log.Println("runtime.NumCPU()=", runtime.NumCPU())
	log.Println("runtime.Version()=", runtime.Version())
	log.Println("os.Getpid()=", os.Getpid())
}
