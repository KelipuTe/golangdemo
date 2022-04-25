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
  // 接收缓冲区最大大小
  RecvBufferMax uint64 = 10 * 1048576
)

// TCPConnection TCP 连接
type TCPConnection struct {
  // 连接状态，详见 RunStatus 开头的常量
  runStatus uint8

  // TCP 连接所属 TCP 客户端
  p1client *TCPClient

  // 协议名称
  protocolName string
  // protocol.Protocol
  p1protocol protocol.Protocol

  // net.Conn
  p1conn net.Conn
  // 接收缓冲区
  sli1recvBuffer []byte
  // 接收缓冲区最大大小
  recvBufferMax uint64
  // 接收缓冲区当前大小
  recvBufferNow uint64
}

// NewTCPConnection 创建 TCPConnection
func NewTCPConnection(p1client *TCPClient, p1conn net.Conn) *TCPConnection {
  p1connection := &TCPConnection{
    runStatus:      RunStatusOn,
    p1client:       p1client,
    protocolName:   "",
    p1protocol:     nil,
    p1conn:         p1conn,
    sli1recvBuffer: make([]byte, RecvBufferMax),
    recvBufferMax:  RecvBufferMax,
    recvBufferNow:  0,
  }

  p1connection.protocolName = p1client.protocolName

  switch p1connection.protocolName {
  case protocol.HTTPStr:
    p1connection.p1protocol = http.NewHTTP()
  case protocol.StreamStr:
    p1connection.p1protocol = stream.NewStream()
  case protocol.WebSocketStr:
    p1connection.p1protocol = websocket.NewWebSocket()
  }

  return p1connection
}

// IsRun TCP 连接是不是正在运行
func (p1this *TCPConnection) IsRun() bool {
  return RunStatusOn == p1this.runStatus
}

// TCPClient.IsDebug
func (p1this *TCPConnection) IsDebug() bool {
  return p1this.p1client.IsDebug()
}

// HandleConnection 处理连接
func (p1this *TCPConnection) HandleConnection(deferFunc func()) {
  defer func() {
    deferFunc()
  }()

  switch p1this.protocolName {
  case protocol.HTTPStr:
    // 处理 HTTP 消息
    resp := http.NewResponse()
    resp.SetStatusCode(http.StatusOk)
    respStr := resp.MakeResponse(fmt.Sprintf("this is %s.", p1this.p1client.name))
    p1this.SendMsg([]byte(respStr))
  case protocol.StreamStr:
    // 处理自定义字节流消息
    t1p1protocol := p1this.p1protocol.(*stream.Stream)
    t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", p1this.p1client.name))
    p1this.SendMsg([]byte{})
  case protocol.WebSocketStr:
    // t1p1protocol := p1this.p1conn.p1protocol.(*websocket.WebSocket)
    // if t1p1protocol.IsHandshakeStatusNo() {
    //   handshake, _ := t1p1protocol.HandShakeClient()
    //   err = p1this.p1conn.WriteData(handshake)
    // }
  }

  for p1this.IsRun() {
    byteNum, err := p1this.p1conn.Read(p1this.sli1recvBuffer[p1this.recvBufferNow:])

    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.byteNum: %d", p1this.p1client.name, byteNum))
    }

    if nil != err {
      if err == io.EOF {
        // 对端关闭了连接
        p1this.CloseConnection()
        return
      }
      p1this.p1client.OnClientError(p1this.p1client, err)
      return
    }

    p1this.recvBufferNow += uint64(byteNum)

    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.recvBufferNow: %d", p1this.p1client.name, p1this.recvBufferNow))
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.sli1recvBuffer:", p1this.p1client.name))
      fmt.Println(string(p1this.sli1recvBuffer[0:p1this.recvBufferNow]))
    }

    p1this.HandleBuffer()
  }
}

func (p1this *TCPConnection) HandleBuffer() {
  sli1Copy := p1this.sli1recvBuffer[0:p1this.recvBufferNow]
  for p1this.recvBufferNow > 0 {
    firstMsgLength, err := p1this.p1protocol.FirstMsgLength(sli1Copy)

    sli1firstMsg := p1this.sli1recvBuffer[0:firstMsgLength]
    p1this.p1client.OnConnRequest(p1this)

    switch p1this.protocolName {
    case protocol.StreamStr:
      t1p1protocol := p1this.p1protocol.(*stream.Stream)
      t1p1protocol.Decode(sli1firstMsg)

      if p1this.IsDebug() {
        fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleBuffer.StreamStr.Decode: ", p1this.p1client.name))
      }
    case protocol.WebSocketStr:
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
    }

    p1this.sli1recvBuffer = p1this.sli1recvBuffer[firstMsgLength:]
    p1this.recvBufferNow -= firstMsgLength
    if p1this.recvBufferNow <= 0 {
      p1this.recvBufferNow = 0
      break
    }
  }
}

// SendMsg 发送数据
func (p1this *TCPConnection) SendMsg(sli1msg []byte) {
  switch p1this.protocolName {
  case protocol.TCPStr, protocol.HTTPStr, protocol.WebSocketStr:
    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", p1this.p1client.name))
      fmt.Println(string(sli1msg))
    }
    p1this.WriteData(sli1msg)
  case protocol.StreamStr:
    t1sli1msg, _ := p1this.p1protocol.Encode()
    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", p1this.p1client.name))
      fmt.Println(string(t1sli1msg))
    }
    p1this.WriteData(t1sli1msg)
  }
}

// WriteData 发送数据
func (p1this *TCPConnection) WriteData(sli1data []byte) (err error) {
  byteNum, err := p1this.p1conn.Write(sli1data)

  if p1this.IsDebug() {
    fmt.Println(fmt.Sprintf("%s.TCPConnection.WriteData.byteNum: %d", p1this.p1client.name, byteNum))
  }

  if nil != err {
    p1this.p1client.OnClientError(p1this.p1client, err)
    p1this.CloseConnection()
  }

  if byteNum != len(sli1data) {
    return errors.New("write byte != data length")
  }
  return nil
}

// CloseConnection 关闭连接
func (p1this *TCPConnection) CloseConnection() {
  p1this.runStatus = RunStatusOff
  p1this.recvBufferNow = 0
  p1this.p1conn.Close()
  p1this.p1client.OnConnClose(p1this)
}
