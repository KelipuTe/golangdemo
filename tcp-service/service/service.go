package service

import (
	"demo-golang/tcp-service/config"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// TCPService tcp服务端
type TCPService struct {
	name         string             //名称
	runStatus    config.RunStatus   //运行状态
	debugStatus  config.DebugStatus //debug开关
	protocolName string             //使用的协议
	address      string             //ip地址
	port         uint16             //端口号

	listener     net.Listener              //socket本体
	connPool     map[string]*TCPConnection //连接上来的tcp
	connPoolLock *sync.Mutex               //连接专用的锁
	maxConnNum   uint32                    //最大tcp连接数
	nowConnNum   uint32                    //当前tcp连接数

	OnServiceStart func(*TCPService)        //服务端，启动事件回调
	OnServiceError func(*TCPService, error) //服务端，错误事件回调
	OnConnConnect  func(*TCPConnection)     //tcp连接，连接事件回调
	OnConnRequest  func(*TCPConnection)     //tcp连接，请求事件回调
	OnConnClose    func(*TCPConnection)     //tcp连接，关闭事件回调
}

func NewTCPService(protocolName string, address string, port uint16) *TCPService {
	return &TCPService{
		name:         defaultName,
		runStatus:    config.RunStatusOff,
		debugStatus:  config.DebugStatusOff,
		protocolName: protocolName,
		address:      address,
		port:         port,

		connPool:     make(map[string]*TCPConnection),
		maxConnNum:   1024,
		nowConnNum:   0,
		connPoolLock: &sync.Mutex{},

		OnServiceStart: defaultOnServiceStart,
		OnServiceError: defaultOnServiceError,
		OnConnConnect:  defaultOnConnConnect,
		OnConnRequest:  defaultOnConnRequest,
		OnConnClose:    defaultOnConnClose,
	}
}

// SetName 设置服务名称
func (s *TCPService) SetName(name string) {
	s.name = name
}

// GetName 获取服务名称
func (s *TCPService) GetName() string {
	return s.name
}

// IsRun 是不是运行中
func (s *TCPService) IsRun() bool {
	return s.runStatus == config.RunStatusOn
}

// OpenDebug 打开debug
func (s *TCPService) OpenDebug() {
	s.debugStatus = config.DebugStatusOn
}

// IsDebug 是否是debug模式
func (s *TCPService) IsDebug() bool {
	return s.debugStatus == config.DebugStatusOn
}

// Start 服务启动
func (s *TCPService) Start() {
	s.StartInfo()

	address := s.address + ":" + strconv.Itoa(int(s.port))
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		errStr := fmt.Sprintf("service [%s] Start() with err:%s", s.name, err.Error())
		s.OnServiceError(s, errors.New(errStr))
		return
	}

	s.listener = listener
	defer s.listener.Close()

	s.runStatus = config.RunStatusOn
	s.OnServiceStart(s)
	s.StartListen()
}

// StartInfo 输出服务的配置和环境参数
func (s *TCPService) StartInfo() {
	log.Println("version: ", config.Version)
	log.Println("runtime.GOOS=", runtime.GOOS)
	log.Println("runtime.NumCPU()=", runtime.NumCPU())
	log.Println("runtime.Version()=", runtime.Version())
	log.Println("os.Getpid()=", os.Getpid())
}

// StartListen 开始监听
func (s *TCPService) StartListen() {
	for s.IsRun() {
		// net.Listener.Accept，系统调用，获取连接上来的tcp
		conn, err := s.listener.Accept()
		if err != nil {
			errStr := fmt.Sprintf("service [%s] StartListen() with err:%s", s.name, err.Error())
			s.OnServiceError(s, errors.New(errStr))
			return
		}

		// 判断当前tcp连接数是否超过最大tcp连接数
		if s.nowConnNum >= s.maxConnNum {
			errStr := fmt.Sprintf("service [%s] nowConnNum >= maxConnNum", s.name)
			s.OnServiceError(s, errors.New(errStr))
		}

		tcpConn := NewTCPConnection(s, conn)
		s.AddConnection(tcpConn)
		tcpConn.runStatus = config.RunStatusOn
		s.OnConnConnect(tcpConn)
		go tcpConn.HandleConnection() //tcp连接丢出去自己执行
	}
}

// AddConnection 添加连接
func (s *TCPService) AddConnection(tcpConn *TCPConnection) {
	//用c编码时，可以通过socket的文件描述符区分tcp连接
	//在go中也可以获得文件描述符，但是文件描述符不是唯一的，所以不能用于区分tcp连接
	if s.IsDebug() {
		fd, err := tcpConn.netConn.(*net.TCPConn).File()
		fmt.Println("net.TCPConn.File", fd.Fd(), err)
	}

	//这里不用处理并发问题，建立连接的时候，是单线程的
	addrStr := tcpConn.netConn.RemoteAddr().String()
	s.nowConnNum++
	s.connPool[addrStr] = tcpConn

	if s.IsDebug() {
		fmt.Println("net.TCPConn.RemoteAddr.String", addrStr)
	}
}

// DeleteConnection 移除连接
func (s *TCPService) DeleteConnection(tcpConn *TCPConnection) {
	//这里需要处理并发问题，关闭连接的时候，存在并发情况
	s.connPoolLock.Lock()
	defer s.connPoolLock.Unlock()

	addrStr := tcpConn.netConn.RemoteAddr().String()
	if _, ok := s.connPool[addrStr]; ok {
		delete(s.connPool, addrStr)
		s.nowConnNum--
	}
}
