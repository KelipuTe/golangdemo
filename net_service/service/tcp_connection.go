package service

import (
  "demo_golang/net_service/protocol"
  "fmt"
  "net"
)

const (
  RecvBufferMaxLength = 1048576 // 接收缓冲区最大大小，1024*1024
)

// TCP连接结构体
type TcpConnection struct {
  NetService     *NetService               // 服务结构体实例地址
  Protocol       protocol.Protocol         // 协议接口实例
  Conn           net.Conn                  // net.Conn实例
  Arr1RecvBuffer [RecvBufferMaxLength]byte // 接收缓冲区
  RecvNowLength  int                       // 接收缓冲区当前大小
}

// 构造TCP连接结构体
func MakeTcpConnection(p1NetSvc *NetService, conn net.Conn) (p1TcpCnct *TcpConnection, err error) {
  var ptc protocol.Protocol
  switch p1NetSvc.ProtocolName {
  case "http":
    ptc = &protocol.Http{}
    break
  }
  p1TcpCnct = &TcpConnection{
    NetService:    p1NetSvc,
    Protocol:      ptc,
    Conn:          conn,
    RecvNowLength: 0,
  }
  return
}

// 处理Tcp连接
func (p1this *TcpConnection) HandleTcpConnection() {
  for 1 == p1this.NetService.ServiceRunning {
    recvByteNum, err := p1this.Conn.Read(p1this.Arr1RecvBuffer[p1this.RecvNowLength:])
    fmt.Println("recvByteNum=", recvByteNum)
    p1this.RecvNowLength += recvByteNum
    fmt.Println("p1this.Arr1RecvLength=", p1this.RecvNowLength)
    if nil != err {
      p1this.NetService.OnError("HandleMessage()," + err.Error())
      return
    }
    fmt.Println("p1this.Arr1RecvBuffer=", p1this.Arr1RecvBuffer[0:p1this.RecvNowLength])
    fmt.Println("p1this.Arr1RecvBuffer=", string(p1this.Arr1RecvBuffer[0:p1this.RecvNowLength]))
  }
}
