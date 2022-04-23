package service

import (
  "demo_golang/tcp_service_v1/internal/tool/debug"
  goErrors "errors"
  "log"
  "net"
  "os"
  "runtime"
  "strconv"

  pkgErrors "github.com/pkg/errors"
)

const (
  RunningStatusOff uint8 = iota // 服务（连接）关闭
  RunningStatusOn               // 服务（连接）运行
)

const (
  DebugStatusOff uint8 = iota // debug 关
  DebugStatusOn               // debug 开
)

// TCPService TCP 服务端
type TCPService struct {
  // runningStatus 服务状态，0（关闭）；1（运行）
  runningStatus uint8
  // debugStatus debug 开关，0（关）；1（开）
  debugStatus uint8

  // protocolName 协议名称
  protocolName string

  // address IP 地址
  address string
  // port 端口号
  port uint16
  // p1Listener net.Listener
  p1Listener net.Listener

  // mapConnectionPool 连接池
  mapConnectionPool map[string]*TCPConnection
  // maxConnectionNum 最大连接数
  maxConnectionNum uint32
  // nowConnectionNum 当前连接数
  nowConnectionNum uint32

  // OnStart 服务启动事件回调
  OnStart func(*TCPService)
  // OnError 服务错误事件回调
  OnError func(*TCPService, error)
  // OnConnect TCP 连接事件回调
  OnConnect func(*TCPConnection)
  // OnRequest TCP 响应事件回调
  OnRequest func(*TCPConnection)
  // OnClose TCP 关闭事件回调
  OnClose func(*TCPConnection)
}

// NewTCPService 创建默认的 TCPService
func NewTCPService(protocolName string, address string, port uint16) *TCPService {
  return &TCPService{
    runningStatus: RunningStatusOn,
    debugStatus:   DebugStatusOn,
    protocolName:  protocolName,
    address:       address,
    port:          port,

    mapConnectionPool: make(map[string]*TCPConnection),
    maxConnectionNum:  1024,
    nowConnectionNum:  0,
  }
}

// IsDebug 是否是 debug 模式
func (p1this *TCPService) IsDebug() bool {
  return DebugStatusOn == p1this.debugStatus
}

// Start 服务启动
func (p1this *TCPService) Start() {
  t1address := p1this.address + ":" + strconv.Itoa(int(p1this.port))
  listener, err := net.Listen("tcp4", t1address)
  if nil != err {
    p1this.OnError(p1this, pkgErrors.WithMessage(err, "TCPService.Start"))
    return
  }

  p1this.p1Listener = listener
  defer p1this.p1Listener.Close()

  p1this.StartInfo()
  p1this.OnStart(p1this)
  p1this.StartListen()
}

// StartInfo 输出服务配置和环境参数
func (p1this *TCPService) StartInfo() {
  log.Println("runtime.GOOS=", runtime.GOOS)
  log.Println("runtime.NumCPU()=", runtime.NumCPU())
  log.Println("runtime.Version()=", runtime.Version())
  log.Println("os.Getpid()=", os.Getpid())
}

// StartListen 开始监听
func (p1this *TCPService) StartListen() {
  for RunningStatusOn == p1this.runningStatus {
    // 系统调用，获取 TCP 连接
    p1conn, err := p1this.p1Listener.Accept()
    if nil != err {
      p1this.OnError(p1this, pkgErrors.WithMessage(err, "TCPService.StartListen"))
      return
    }
    // 判断连接数量是否超过最大数量限制
    if p1this.nowConnectionNum >= p1this.maxConnectionNum {
      p1this.OnError(p1this, pkgErrors.WithMessage(goErrors.New("NowConnectionNum >= MaxConnectionNum"), "TCPService.StartListen"))
    }

    p1connection := NewTCPConnection(p1this, p1conn)
    p1this.AddConnection(p1connection)
    p1this.OnConnect(p1connection)
    // 启动协程，处理连接
    go p1connection.HandleConnection()
  }
}

// AddConnection 添加连接
func (p1this *TCPService) AddConnection(p1connection *TCPConnection) {
  // 用 Linux C 编码时，可以通过 socket 的文件描述符区分 TCP 连接
  // 在 go 中也可以获得文件描述符，但是文件描述符不是唯一的，不能用于区分
  fd, err := p1connection.p1Conn.(*net.TCPConn).File()
  debug.Println(p1this.IsDebug(), "net.TCPConn.File", fd.Fd(), err)

  p1this.nowConnectionNum++
  addrStr := p1connection.p1Conn.RemoteAddr().String()
  p1this.mapConnectionPool[addrStr] = p1connection
  debug.Println(p1this.IsDebug(), "net.TCPConn.RemoteAddr.String", addrStr)
}

// DeleteConnection 移除连接
func (p1this *TCPService) DeleteConnection(p1connection *TCPConnection) {
  addrStr := p1connection.p1Conn.RemoteAddr().String()
  if _, ok := p1this.mapConnectionPool[addrStr]; ok {
    delete(p1this.mapConnectionPool, addrStr)
    p1this.nowConnectionNum--
  }
}
