package service

import (
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/protocol/http"
  "demo_golang/tcp_service_v1/internal/protocol/stream"
  "demo_golang/tcp_service_v1/internal/protocol/websocket"
  "demo_golang/tcp_service_v1/internal/tool/debug"
  "errors"
  "fmt"
  "net"
)

const (
  // RecvBufferMax 接收缓冲区最大大小，1MB==2^20==1048576。
  // uint32，最大 2^32-1，差不多 4G，理论上应该够用了。uint64 只会更大。
  RecvBufferMax uint64 = 10 * 1048576
)

// TCPConnection TCP 连接
type TCPConnection struct {
  // 连接状态，详见 RunningStatus 开头的常量
  runningStatus uint8

  // TCP 连接所属 TCP 服务端
  p1service *TCPService

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

// NewTCPConnection 创建 TCPConnection
func NewTCPConnection(p1service *TCPService, p1conn net.Conn) *TCPConnection {
  p1connection := &TCPConnection{
    runningStatus:  RunningStatusOn,
    p1service:      p1service,
    protocolName:   "",
    p1protocol:     nil,
    p1Conn:         p1conn,
    sli1recvBuffer: make([]byte, RecvBufferMax),
    recvBufferMax:  RecvBufferMax,
    recvBufferNow:  0,
  }

  p1connection.protocolName = p1service.protocolName

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

// TCPService.IsDebug
func (p1this *TCPConnection) IsDebug() bool {
  return p1this.p1service.IsDebug()
}

// HandleConnection 处理连接
func (p1this *TCPConnection) HandleConnection() {
  for RunningStatusOn == p1this.runningStatus {
    // 系统调用，从 socket 读取数据
    byteNum, err := p1this.p1Conn.Read(p1this.sli1recvBuffer[p1this.recvBufferNow:])
    debug.Println(p1this.IsDebug(), "TCPConnection.HandleConnection.byteNum: ", byteNum)
    if nil != err {
      p1this.p1service.OnError(p1this.p1service, err)
      return
    }

    p1this.recvBufferNow += uint64(byteNum)
    debug.Println(p1this.IsDebug(), "TCPConnection.HandleConnection.recvBufferNow: ", p1this.recvBufferNow)
    debug.Println(p1this.IsDebug(), "TCPConnection.HandleConnection.sli1recvBuffer: ")
    debug.Println(p1this.IsDebug(), string(p1this.sli1recvBuffer[0:p1this.recvBufferNow]))

    p1this.HandleWithProtocol()
  }
}

// HandleWithProtocol 用协议处理接收到的消息
func (p1this *TCPConnection) HandleWithProtocol() {
  sli1Copy := p1this.sli1recvBuffer[0:p1this.recvBufferNow]
  for p1this.recvBufferNow > 0 {
    firstMsgLength, err := p1this.p1protocol.FirstMsgLength(sli1Copy)
    if nil != err {
      // if protocol.StrHTTP == p1this.protocolName {
      //   p1http := p1this.p1protocol.(*http.HTTP)
      //   switch p1http.ParseStatus {
      //   case protocol.HTTP_STATUS_NO_DATA,
      //     protocol.HTTP_STATUS_NOT_HTTP,
      //     protocol.HTTP_STATUS_NOT_FINISH:
      //     // 继续接收
      //   case protocol.HTTP_STATUS_WRONG_DATA,
      //     protocol.HTTP_STATUS_TOO_LONG,
      //     protocol.HTTP_STATUS_ATOI_ERR:
      //     // 明显出错
      //     p1this.CloseTcpCnct()
      //   }
      // }
      // 处理完错误后，跳出for
      break
    }
    // 取出第 1 条完整的报文
    sli1firstMsg := p1this.sli1recvBuffer[0:firstMsgLength]
    p1this.p1service.OnRequest(p1this)

    switch p1this.protocolName {
    case protocol.StrHTTP:
      // 处理 HTTP 请求
      t1p1protocol := p1this.p1protocol.(*http.HTTP)
      t1p1protocol.Decode(sli1firstMsg)
      debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.Decode: ")
      debug.Println(p1this.IsDebug(), fmt.Sprintf("%+v", t1p1protocol))

      // 返回响应数据
      resp := http.NewResponse()
      resp.SetStatusCode(http.StatusOk)
      respStr := resp.MakeResponse("this is service.")
      p1this.SendMsg(respStr)

      p1this.CloseConnection()
      return
    case protocol.StrStream:
      t1p1protocol := p1this.p1protocol.(*stream.Stream)
      t1p1protocol.Decode(sli1firstMsg)
      debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.Decode: ")
      debug.Println(p1this.IsDebug(), fmt.Sprintf("%+v", t1p1protocol))
    case protocol.StrWebSocket:
      // 处理 WebSocket 请求
      t1p1protocol := p1this.p1protocol.(*websocket.WebSocket)
      t1p1protocol.Decode(sli1firstMsg)
      debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.Decode: ")
      debug.Println(p1this.IsDebug(), fmt.Sprintf("%+v", t1p1protocol))

      if t1p1protocol.IsHandshakeStatusNo() {
        shakeHandMsg, err := t1p1protocol.Handshake()
        debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.Handshake")
        debug.Println(p1this.IsDebug(), string(shakeHandMsg))
        if err != nil {
          // 发送 400 给客户端，并且关闭连接
          return
        } else {
          err = p1this.WriteData(shakeHandMsg)
          if nil == err {
            t1p1protocol.SetHandshakeStatusYes()
          }
        }
      } else if t1p1protocol.IsHandshakeStatusYes() {
        // 返回响应数据
        resp := websocket.NewResponse()
        sli1resp := resp.MakeResponse("this is service.")
        debug.Println(p1this.IsDebug(), "TCPConnection.HandleWithProtocol.MakeResponse: ")
        debug.Println(p1this.IsDebug(), sli1resp)
        p1this.WriteData(sli1resp)
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

// CloseConnection 关闭连接
func (p1this *TCPConnection) CloseConnection() {
  p1this.runningStatus = RunningStatusOff
  p1this.recvBufferNow = 0
  p1this.p1Conn.Close()
  p1this.p1service.OnClose(p1this)
  p1this.p1service.DeleteConnection(p1this)
}

// SendMsg 发送数据
func (p1this *TCPConnection) SendMsg(msg string) {
  debug.Println(p1this.IsDebug(), "TCPConnection.SendMsg.msg: ")
  debug.Println(p1this.IsDebug(), msg)

  switch p1this.protocolName {
  case protocol.StrTCP, protocol.StrHTTP, protocol.StrWebSocket:
    p1this.WriteData([]byte(msg))
    // case protocol.StrStream:
    //   encodeData, _ := p1this.p1protocol.Encode()
    //   p1this.WriteData(encodeData)
    //   break
  }
}

// WriteData 发送数据
func (p1this *TCPConnection) WriteData(sli1data []byte) error {
  // 系统调用，用 socket 发送数据
  byteNum, err := p1this.p1Conn.Write(sli1data)
  debug.Println(p1this.IsDebug(), "TCPConnection.WriteData.byteNum: ", byteNum)
  if nil != err {
    p1this.CloseConnection()
    p1this.p1service.OnError(p1this.p1service, err)
  }
  if byteNum != len(sli1data) {
    return errors.New("write byte != data length")
  }
  return nil
}
