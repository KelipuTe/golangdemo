package client

import (
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/protocol/websocket"
  "net"
  "strconv"
  "sync"

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

// TCPClient TCP 客户端
type TCPClient struct {
  // runningStatus 服务状态，详见 RunningStatus 开头的常量
  runningStatus uint8
  // debugStatus debug 开关，详见 DebugStatus 开头的常量
  debugStatus uint8

  // protocolName 协议名称
  protocolName string

  // address IP 地址
  address string
  // port 端口号
  port uint16

  // 连接
  p1connection *TCPConnection

  // OnStart 服务启动事件回调
  OnStart func(*TCPClient)
  // OnError 服务错误事件回调
  OnError func(*TCPClient, error)
  // OnConnect TCP 连接事件回调
  OnConnect func(*TCPConnection)
  // OnRequest TCP 响应事件回调
  OnRequest func(*TCPConnection)
  // OnClose TCP 关闭事件回调
  OnClose func(*TCPConnection)
}

// NewTCPClient 创建默认的 TCPClient
func NewTCPClient(protocolName string, address string, port uint16) *TCPClient {
  return &TCPClient{
    runningStatus: RunningStatusOn,
    debugStatus:   DebugStatusOn,
    protocolName:  protocolName,
    address:       address,
    port:          port,
  }
}

// IsDebug 是否是 debug 模式
func (p1this *TCPClient) IsDebug() bool {
  return DebugStatusOn == p1this.debugStatus
}

func (p1this *TCPClient) Start() {
  p1conn, err := net.Dial("tcp4", p1this.address+":"+strconv.Itoa(int(p1this.port)))
  if nil != err {
    p1this.OnError(p1this, pkgErrors.WithMessage(err, "TCPClient.StartListen"))
    return
  }
  p1this.OnStart(p1this)

  p1this.p1connection = NewTCPConnection(p1this, p1conn)

  if protocol.StrWebSocket == p1this.protocolName {
    t1p1protocol := p1this.p1connection.p1protocol.(*websocket.WebSocket)
    if t1p1protocol.IsHandshakeStatusNo() {
      handshake, _ := t1p1protocol.HandShakeClient()
      err = p1this.p1connection.WriteData(handshake)
    }
  }

  p1this.OnConnect(p1this.p1connection)

  var wg sync.WaitGroup
  wg.Add(1)
  go p1this.p1connection.HandleConnection(wg.Done)
  wg.Wait()
}
