package client

import (
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/protocol/http"
  "demo_golang/tcp_service_v1/internal/protocol/stream"
  "demo_golang/tcp_service_v1/internal/protocol/websocket"
  "demo_golang/tcp_service_v1/internal/tool/debug"
  "errors"
  "fmt"
  "io"
  "net"
)

const (
  // RecvBufferMax 接收缓冲区最大大小，1MB==2^20==1048576。
  RecvBufferMax uint64 = 10 * 1048576
)

// TCPConnection TCP 连接结
type TCPConnection struct {
  // 连接状态，详见 RunningStatus 开头的常量
  runningStatus uint8

  // TCP 连接所属 TCP 客户端
  p1client *TCPClient

  // 协议名称
  protocolName string
  // protocol.Protocol
  p1protocol protocol.Protocol

  // net.Conn
  p1Conn net.Conn
  // 接收缓冲区
  sli1recvBuffer []byte
  // 接收缓冲区最大大小
  recvBufferMax uint64
  // 接收缓冲区当前大小
  recvBufferNow uint64
}

func NewTCPConnection(p1client *TCPClient, p1conn net.Conn) *TCPConnection {
  p1connection := &TCPConnection{
    runningStatus:  RunningStatusOn,
    p1client:       p1client,
    protocolName:   "",
    p1protocol:     nil,
    p1Conn:         p1conn,
    sli1recvBuffer: make([]byte, RecvBufferMax),
    recvBufferMax:  RecvBufferMax,
    recvBufferNow:  0,
  }

  p1connection.protocolName = p1client.protocolName

  switch p1connection.protocolName {
  case protocol.StrHTTP:
    p1connection.p1protocol = http.NewHTTP()
  case protocol.StrStream:
    p1connection.p1protocol = stream.NewStream()
  case protocol.StrWebSocket:
    p1connection.p1protocol = websocket.NewWebSocket()
  }

  return p1connection
}

// TCPClient.IsDebug
func (p1this *TCPConnection) IsDebug() bool {
  return p1this.p1client.IsDebug()
}

func (p1this *TCPConnection) HandleConnection(deferFunc func()) {
  defer func() {
    deferFunc()
  }()

  if protocol.StrStream == p1this.protocolName {
    t1p1protocol := p1this.p1protocol.(*stream.Stream)
    t1p1protocol.DecodeMsg = "client stream."
    sli1msg, _ := p1this.p1protocol.Encode()
    p1this.Send(string(sli1msg))
  } else if protocol.StrWebSocket == p1this.protocolName {
    for p1this.runningStatus == RunningStatusOn {

      byteNum, err := p1this.p1Conn.Read(p1this.sli1recvBuffer[p1this.recvBufferNow:])
      debug.Println(p1this.IsDebug(), "TCPConnection.HandleConnection.byteNum: ", byteNum)
      if nil != err {
        if err == io.EOF {
          // 对端关闭了连接
          p1this.CloseConnection()
          return
        }
        p1this.p1client.OnError(p1this.p1client, err)
        return
      }

      p1this.recvBufferNow += uint64(byteNum)

      sli1Copy := p1this.sli1recvBuffer[0:p1this.recvBufferNow]
      for p1this.recvBufferNow > 0 {
        firstMsgLength, err := p1this.p1protocol.FirstMsgLength(sli1Copy)
        // 取出第 1 条完整的报文
        sli1firstMsg := p1this.sli1recvBuffer[0:firstMsgLength]
        p1this.p1client.OnRequest(p1this)

        _ = p1this.p1protocol.Decode(sli1firstMsg)
        t1p1protocol := p1this.p1protocol.(*websocket.WebSocket)
        if t1p1protocol.IsHandshakeStatusNo() {
          err = t1p1protocol.VerifyShakeHand()
          if err == nil {
            t1p1protocol.SetHandshakeStatusYes()
            resp := websocket.NewResponse()
            sli1resp := resp.MakeRequest("client webserver")
            debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.MakeRequest: ")
            debug.Println(p1this.IsDebug(), sli1resp)
            p1this.WriteData(sli1resp)
          } else {
            p1this.CloseConnection()
          }
        } else if t1p1protocol.IsHandshakeStatusYes() {

          debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.Decode: ")
          debug.Println(p1this.IsDebug(), fmt.Sprintf("%+v", t1p1protocol))
        }

        // 处理接收缓冲区中剩余的数据
        p1this.sli1recvBuffer = p1this.sli1recvBuffer[firstMsgLength:]
        p1this.recvBufferNow -= firstMsgLength
        if p1this.recvBufferNow <= 0 {
          p1this.recvBufferNow = 0
          break
        }
      }

    }
  }
}

// CloseConnection 关闭连接
func (p1this *TCPConnection) CloseConnection() {
  p1this.runningStatus = RunningStatusOff
  p1this.recvBufferNow = 0
  p1this.p1Conn.Close()
  p1this.p1client.OnClose(p1this)
}

func (p1this *TCPConnection) Send(msg string) {
  p1this.WriteData([]byte(msg))
}

// WriteData 发送数据
func (p1this *TCPConnection) WriteData(sli1data []byte) (err error) {
  // 系统调用，用 socket 发送数据
  byteNum, err := p1this.p1Conn.Write(sli1data)
  debug.Println(p1this.IsDebug(), "TCPConnection.WriteData.byteNum: ", byteNum)
  if nil != err {
    p1this.CloseConnection()
    p1this.p1client.OnError(p1this.p1client, err)
  }
  if byteNum != len(sli1data) {
    return errors.New("write byte != data length")
  }
  return nil
}
