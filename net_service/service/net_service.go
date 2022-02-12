package service

import (
  "fmt"
  "net"
  "os"
  "runtime"
  "strconv"
)

// 服务结构体
type NetService struct {
  AppDebug       int // 调试模式，1=开启调试
  ServiceRunning int // 服务运行状态，1=服务运行

  ProtocolName string       // 协议类型
  Address      string       // 地址
  Port         int          // 端口
  Listener     net.Listener // net.Listener实例

  MapTcpCnctPool map[string]*TcpConnection // Tcp连接池
  NowTcpCnctNum  int                       // 当前连接数
  MaxTcpCnctNum  int                       // 最大连接数

  OnStart   func(p1NetSvc *NetService)     // 服务启动事件回调
  OnError   func(errStr string)            // 服务错误事件回调
  OnConnect func(p1TcpCnct *TcpConnection) // tcp连接事件回调
  OnClose   func(p1TcpCnct *TcpConnection) // tcp关闭事件回调
}

// 拼接地址和端口
func (p1this *NetService) JoinAddrAndPort() string {
  return p1this.Address + ":" + strconv.Itoa(p1this.Port)
}

// 服务启动
func (p1this *NetService) Start() {
  listener, err := net.Listen("tcp4", p1this.JoinAddrAndPort())
  if nil != err {
    p1this.OnError("Start()," + err.Error())
    return
  }
  p1this.Listener = listener
  defer p1this.Listener.Close()
  p1this.StartInfo()
  p1this.OnStart(p1this)
  p1this.StartListen()
}

// 输出服务配置和环境参数
func (p1this *NetService) StartInfo() {
  fmt.Println("NetService.AppDebug=", p1this.AppDebug)
  fmt.Println("NetService.ServiceRunning=", p1this.ServiceRunning)
  fmt.Println("NetService.ProtocolName=", p1this.ProtocolName)
  fmt.Println("NetService.Address=", p1this.Address)
  fmt.Println("NetService.Port=", p1this.Port)
  fmt.Println("runtime.GOOS=", runtime.GOOS)
  fmt.Println("runtime.NumCPU()=", runtime.NumCPU())
  fmt.Println("runtime.Version()=", runtime.Version())
  fmt.Println("os.Getpid()=", os.Getpid())
}

// 开始监听
func (p1this *NetService) StartListen() {
  for 1 == p1this.ServiceRunning {
    conn, err := p1this.Listener.Accept()
    if nil != err {
      p1this.OnError("StartListen(),Accept()," + err.Error())
      return
    }
    if p1this.NowTcpCnctNum >= p1this.MaxTcpCnctNum {
      p1this.OnError("StartListen(),p1this.NowCNCTNum >= p1this.MaxCNCTNum")
    }
    p1TcpCnct, err := MakeTcpConnection(p1this, conn)
    if nil != err {
      p1this.OnError("StartListen(),MakeTcpConnection()," + err.Error())
      return
    }
    p1this.AddTcpCnct(p1TcpCnct)
    p1this.OnConnect(p1TcpCnct)
    go p1TcpCnct.HandleTcpConnection()
  }
}

// 添加tcp连接
func (p1this *NetService) AddTcpCnct(p1TcpCnct *TcpConnection) {
  p1this.MapTcpCnctPool[p1TcpCnct.Conn.RemoteAddr().String()] = p1TcpCnct
  p1this.NowTcpCnctNum++
}

// 移除tcp连接
func (p1this *NetService) DelTcpCnct(p1TcpCnct *TcpConnection) {

}
