package client

import (
	"demo_golang/tcp_service_v1/config"
	"demo_golang/tcp_service_v1/protocol"
	"demo_golang/tcp_service_v1/tool"
	"errors"
	"net"
)

const (
  CONNECTION_RUNNING_OFF = 0
  CONNECTION_RUNNING_ON  = 1
  RECV_BUFFER_MAX_LENGTH = 1048576 // 接收缓冲区最大大小，1024*1024
)

// tcp连接结构体
type TcpConnection struct {
  NetClient         *NetClient                   // 客户端结构体实例地址
  ConnectionRunning int                          // 连接运行状态，1=连接运行
  ProtocolName      string                       // 协议名称
  Protocol          protocol.Protocol            // 协议接口实例
  Conn              net.Conn                     // net.Conn实例
  Arr1RecvBuffer    [RECV_BUFFER_MAX_LENGTH]byte // 接收缓冲区
  RecvMaxLength     int                          // 接收缓冲区最大大小
  RecvNowLength     int                          // 接收缓冲区当前大小
}

// 构造tcp连接结构体
func MakeTcpConnection(p1NetClt *NetClient, conn net.Conn) (p1TcpCnct *TcpConnection) {
  var ptc protocol.Protocol
  switch p1NetClt.ProtocolName {
  case config.STR_HTTP:
    ptc = &protocol.Http{}
    break
  case config.STR_STREAM:
    ptc = &protocol.Stream{}
    break
  }
  p1TcpCnct = &TcpConnection{
    NetClient:         p1NetClt,
    ConnectionRunning: CONNECTION_RUNNING_ON,
    ProtocolName:      p1NetClt.ProtocolName,
    Protocol:          ptc,
    Conn:              conn,
    RecvMaxLength:     RECV_BUFFER_MAX_LENGTH,
    RecvNowLength:     0,
  }
  return
}

// 处理tcp连接
func (p1this *TcpConnection) HandleTcpConnection(deferFunc func()) {
  defer func() {
    deferFunc()
  }()

  if config.STR_STREAM == p1this.ProtocolName {
    p1this.Send("stream")
  }

  for CONNECTION_RUNNING_ON == p1this.ConnectionRunning {
    p1this.HandleMessageWithProtocol()
  }
}

// 用协议处理接收到的消息
func (p1this *TcpConnection) HandleMessageWithProtocol() {

}

// 关闭tcp连接
func (p1this *TcpConnection) CloseTcpCnct() {
  p1this.ConnectionRunning = CONNECTION_RUNNING_OFF
  p1this.RecvNowLength = 0
  p1this.Conn.Close()
  p1this.NetClient.OnClose(p1this)
}

func (p1this *TcpConnection) Send(data string) {
  switch p1this.ProtocolName {
  case config.STR_TCP, config.STR_HTTP:
    p1this.WriteData([]byte(data))
    break
  case config.STR_STREAM, config.STR_WEBSOCKET:
    encodeData, _ := p1this.Protocol.DataEncode([]byte(data))
    p1this.WriteData(encodeData)
    break
  }
}

// 发送数据
func (p1this *TcpConnection) WriteData(data []byte) (err error) {
  tool.DebugPrintln("data=", data)
  dataLen := len(data)
  writeByteNum, errWrite := p1this.Conn.Write(data)
  tool.DebugPrintln("writeByteNum=", writeByteNum)
  if nil != errWrite {
    p1this.NetClient.OnError("WriteData(),Write()," + err.Error())
    p1this.CloseTcpCnct()
  }
  if writeByteNum != dataLen {
    err = errors.New("WRITE_BYTE_NOT_EQUAL")
  }
  return
}
