package client

import (
  "net"
  "strconv"
  "sync"
)

// 客户端结构体
type NetClient struct {
  ProtocolName string // 协议名称，[tcp,stream,http,websocket]
  Address      string // 地址
  Port         int    // 端口

  TcpCnct *TcpConnection // tcp连接

  OnStart   func(p1NetClt *NetClient)      // 服务启动事件回调
  OnError   func(errStr string)            // 服务错误事件回调
  OnConnect func(p1TcpCnct *TcpConnection) // tcp连接事件回调
  OnRequest func(p1TcpCnct *TcpConnection) // tcp响应事件回调
  OnClose   func(p1TcpCnct *TcpConnection) // tcp关闭事件回调
}

// 拼接地址和端口
func (p1this *NetClient) JoinAddrAndPort() string {
  return p1this.Address + ":" + strconv.Itoa(p1this.Port)
}

func (p1this *NetClient) Start() {
  conn, err := net.Dial("tcp4", p1this.JoinAddrAndPort())
  if nil != err {
    p1this.OnError("Start()," + err.Error())
    return
  }
  p1this.OnStart(p1this)

  p1TcpCnct := MakeTcpConnection(p1this, conn)
  p1this.TcpCnct = p1TcpCnct
  p1this.OnConnect(p1TcpCnct)

  var wg sync.WaitGroup
  wg.Add(1)
  go p1TcpCnct.HandleTcpConnection(wg.Done)
  wg.Wait()
}
