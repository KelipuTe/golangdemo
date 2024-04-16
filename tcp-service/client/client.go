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

// TCPClient tcp客户端
type TCPClient struct {
	name         string             //客户端名称
	runStatus    config.RunStatus   //运行状态
	debugStatus  config.DebugStatus //debug 开关状态
	protocolName string             //使用的协议
	address      string             //ip地址
	port         uint16             //端口号

	conn          *TCPConnection          //tcp连接本体
	OnClientStart func(*TCPClient)        //客户端启动事件回调
	OnClientError func(*TCPClient, error) //客户端错误事件回调
	OnConnConnect func(*TCPConnection)    //tcp连接，连接事件回调
	OnConnRequest func(*TCPConnection)    //tcp连接，请求事件回调
	OnConnClose   func(*TCPConnection)    //tcp连接，关闭事件回调
}

func NewTCPClient(protocolName string, address string, port uint16) *TCPClient {
	return &TCPClient{
		name:         defaultName,
		runStatus:    config.RunStatusOff,
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

func (c *TCPClient) SetName(name string) {
	c.name = name
}

func (c *TCPClient) GetName() string {
	return c.name
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

	c.OnClientStart(c)

	conn, err := net.Dial("tcp4", c.address+":"+strconv.Itoa(int(c.port)))
	if err != nil {
		c.OnClientError(c, pkgErrors.WithMessage(err, "TCPClient.StartListen"))
		return
	}

	c.conn = NewTCPConnection(c, conn)
	c.runStatus = config.RunStatusOn
	c.OnConnConnect(c.conn)

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
