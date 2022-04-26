package service

import (
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/protocol/http"
  "demo_golang/tcp_service_v1/internal/protocol/stream"
  "demo_golang/tcp_service_v1/internal/protocol/websocket"
  "errors"
  "fmt"
  "io"
  "net"
)

const (
  // RecvBufferMax 接收缓冲区最大大小，1MB == 2^20 == 1048576。
  // uint32，最大 2^32-1，差不多 4GB，理论上应该够用了，uint64 只会更大。
  // 这里用 uint64 是因为 WebSocket 协议最大负荷可以是 2^64-1 个字节
  RecvBufferMax uint64 = 10 * 1048576
)

// TCPConnection TCP 连接
type TCPConnection struct {
  // 连接状态，详见 RunStatus 开头的常量
  runStatus uint8

  // TCP 连接所属 TCP 服务端
  p1service *TCPService

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
func NewTCPConnection(p1service *TCPService, p1conn net.Conn) *TCPConnection {
  p1connection := &TCPConnection{
    runStatus:      RunStatusOn,
    p1service:      p1service,
    protocolName:   "",
    p1protocol:     nil,
    p1conn:         p1conn,
    sli1recvBuffer: make([]byte, RecvBufferMax),
    recvBufferMax:  RecvBufferMax,
    recvBufferNow:  0,
  }

  p1connection.protocolName = p1service.protocolName

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

// TCPService.IsDebug
func (p1this *TCPConnection) IsDebug() bool {
  return p1this.p1service.IsDebug()
}

// HandleConnection 处理连接
func (p1this *TCPConnection) HandleConnection() {
  for p1this.IsRun() {
    // net.Conn.Read，系统调用，从 socket 读取数据
    byteNum, err := p1this.p1conn.Read(p1this.sli1recvBuffer[p1this.recvBufferNow:])

    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.byteNum: %d", p1this.p1service.name, byteNum))
    }

    if nil != err {
      if err == io.EOF {
        // 对端关闭了连接
        p1this.CloseConnection()
        return
      }
      p1this.p1service.OnServiceError(p1this.p1service, err)
      return
    }

    p1this.recvBufferNow += uint64(byteNum)

    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.recvBufferNow: %d", p1this.p1service.name, p1this.recvBufferNow))
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleConnection.sli1recvBuffer:", p1this.p1service.name))
      fmt.Println(string(p1this.sli1recvBuffer[0:p1this.recvBufferNow]))
    }

    p1this.HandleBuffer()
  }
}

// HandleBuffer 处理缓冲区
func (p1this *TCPConnection) HandleBuffer() {
  sli1Copy := p1this.sli1recvBuffer[0:p1this.recvBufferNow]
  for p1this.recvBufferNow > 0 {
    firstMsgLength, err := p1this.p1protocol.FirstMsgLength(sli1Copy)
    if nil != err {
      if protocol.HTTPStr == p1this.protocolName {
        // 处理 HTTP 解析异常
        p1http := p1this.p1protocol.(*http.HTTP)
        switch p1http.ParseStatus {
        case http.ParseStatusRecvBufferEmpty,
          http.ParseStatusNotHTTP,
          http.ParseStatusIncomplete:
          // 继续接收
        case http.ParseStatusParseErr:
          // 明显出错
          p1this.CloseConnection()
        }
      }
      break
    }
    // 取出第 1 条完整的报文
    sli1firstMsg := p1this.sli1recvBuffer[0:firstMsgLength]
    p1this.p1service.OnConnRequest(p1this)

    switch p1this.protocolName {
    case protocol.HTTPStr:
      p1this.HandleHTTPMsg(sli1firstMsg)
      return
    case protocol.StreamStr:
      p1this.HandleStreamMsg(sli1firstMsg)
      return
    case protocol.WebSocketStr:
      err := p1this.HandleWebSocketMsg(sli1firstMsg)
      if nil != err {
        p1this.CloseConnection()
      }
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

// HandelHTTPMsg 处理 HTTP 消息
func (p1this *TCPConnection) HandleHTTPMsg(sli1firstMsg []byte) {
  t1p1protocol := p1this.p1protocol.(*http.HTTP)
  t1p1protocol.Decode(sli1firstMsg)

  if p1this.IsDebug() {
    fmt.Println(fmt.Sprintf("%s.TCPConnection.HandelHTTPMsg.Decode: ", p1this.p1service.name))
    fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
  }

  // 返回响应数据
  resp := http.NewResponse()
  resp.SetStatusCode(http.StatusOk)
  respStr := resp.MakeResponse(fmt.Sprintf("this is %s.", p1this.p1service.name))
  p1this.SendMsg([]byte(respStr))

  p1this.CloseConnection()
}

// HandleStreamMsg 处理自定义字节流消息
func (p1this *TCPConnection) HandleStreamMsg(sli1firstMsg []byte) {
  t1p1protocol := p1this.p1protocol.(*stream.Stream)
  t1p1protocol.Decode(sli1firstMsg)

  if p1this.IsDebug() {
    fmt.Println(fmt.Sprintf("%s.TCPConnection.HandelStreamMsg.Decode: ", p1this.p1service.name))
    fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
  }

  t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", p1this.p1service.name))
  p1this.SendMsg([]byte{})

  p1this.CloseConnection()
}

// HandleWebSocketMsg 处理 WebSocket 消息
func (p1this *TCPConnection) HandleWebSocketMsg(sli1firstMsg []byte) error {
  t1p1protocol := p1this.p1protocol.(*websocket.WebSocket)
  t1p1protocol.Decode(sli1firstMsg)

  if p1this.IsDebug() {
    fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleWebSocketMsg.Decode: ", p1this.p1service.name))
    fmt.Println(fmt.Sprintf("%+v", t1p1protocol))
  }

  if t1p1protocol.IsHandshakeStatusNo() {
    sli1respMsg, err := t1p1protocol.CheckHandshakeReq()

    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.HandleWebSocketMsg.CheckHandshakeReq: ", p1this.p1service.name))
      fmt.Println(fmt.Sprintf("%+v", string(sli1respMsg)))
    }

    if nil != err {
      // 发送 400 给客户端，并且关闭连接
      resp := http.NewResponse()
      resp.SetStatusCode(http.StatusBadRequest)
      respStr := resp.MakeResponse(fmt.Sprintf("this is %s. handshake err: %s", p1this.p1service.name, err))
      p1this.WriteData([]byte(respStr))

      return err
    } else {
      // SendMsg 方法会编码，这里走 WriteData 方法直接发
      err = p1this.WriteData(sli1respMsg)
      if nil == err {
        t1p1protocol.SetHandshakeStatusYes()
      }
    }
  } else {
    // 返回响应数据
    t1p1protocol.SetDecodeMsg(fmt.Sprintf("this is %s.", p1this.p1service.name))
    p1this.SendMsg([]byte{})
  }

  return nil
}

// SendMsg 发送数据
func (p1this *TCPConnection) SendMsg(sli1msg []byte) {
  switch p1this.protocolName {
  case protocol.TCPStr, protocol.HTTPStr:
    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", p1this.p1service.name))
      fmt.Println(string(sli1msg))
    }
    p1this.WriteData(sli1msg)
  case protocol.StreamStr, protocol.WebSocketStr:
    t1sli1msg, _ := p1this.p1protocol.Encode()
    if p1this.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.TCPConnection.SendMsg: ", p1this.p1service.name))
      fmt.Println(string(t1sli1msg))
    }
    p1this.WriteData(t1sli1msg)
  }
}

// WriteData 发送数据
func (p1this *TCPConnection) WriteData(sli1data []byte) error {
  // net.Conn.Write，系统调用，用 socket 发送数据
  byteNum, err := p1this.p1conn.Write(sli1data)

  if p1this.IsDebug() {
    fmt.Println(fmt.Sprintf("%s.TCPConnection.WriteData.byteNum: %d", p1this.p1service.name, byteNum))
  }

  if nil != err {
    p1this.p1service.OnServiceError(p1this.p1service, err)
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
  p1this.p1service.DeleteConnection(p1this)
  p1this.p1service.OnConnClose(p1this)
}
