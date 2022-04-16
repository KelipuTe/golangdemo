package service

import (
  "demo_golang/net_service/config"
  "demo_golang/net_service/http"
  "demo_golang/net_service/protocol"
  "demo_golang/net_service/tool"
  "errors"
  "fmt"
  "net"
)

const (
  CONNECTION_RUNNING_OFF = 0
  CONNECTION_RUNNING_ON  = 1
  RECV_BUFFER_MAX_LENGTH = 1048576 // 接收缓冲区最大大小，1024*1024
)

// tcp连接结构体
type TcpConnection struct {
  NetService        *NetService                  // 服务端结构体实例地址
  ConnectionRunning int                          // 连接运行状态，1=连接运行
  ProtocolName      string                       // 协议名称
  Protocol          protocol.Protocol            // 协议接口实例
  Conn              net.Conn                     // net.Conn实例
  Arr1RecvBuffer    [RECV_BUFFER_MAX_LENGTH]byte // 接收缓冲区
  RecvMaxLength     int                          // 接收缓冲区最大大小
  RecvNowLength     int                          // 接收缓冲区当前大小
}

// 构造tcp连接结构体
func MakeTcpConnection(p1NetSvc *NetService, conn net.Conn) (p1TcpCnct *TcpConnection) {
  var ptc protocol.Protocol
  switch p1NetSvc.ProtocolName {
  case config.STR_HTTP:
    ptc = &protocol.Http{}
    break
  case config.STR_STREAM:
    ptc = &protocol.Stream{}
    break
  }
  p1TcpCnct = &TcpConnection{
    NetService:        p1NetSvc,
    ConnectionRunning: CONNECTION_RUNNING_ON,
    ProtocolName:      p1NetSvc.ProtocolName,
    Protocol:          ptc,
    Conn:              conn,
    RecvMaxLength:     RECV_BUFFER_MAX_LENGTH,
    RecvNowLength:     0,
  }
  return
}

// 处理tcp连接
func (p1this *TcpConnection) HandleTcpConnection() {
  for CONNECTION_RUNNING_ON == p1this.ConnectionRunning {
    recvByteNum, err := p1this.Conn.Read(p1this.Arr1RecvBuffer[p1this.RecvNowLength:])
    p1this.RecvNowLength += recvByteNum
    tool.DebugPrintln("TcpConnection")
    tool.DebugPrintln("recvByteNum=", recvByteNum)
    tool.DebugPrintln("RecvNowLength=", p1this.RecvNowLength)
    if nil != err {
      p1this.NetService.OnError("TcpConnection," + err.Error())
      return
    }
    tool.DebugPrintln(string(p1this.Arr1RecvBuffer[0:p1this.RecvNowLength]))

    if config.STR_TCP == p1this.NetService.ProtocolName {
      fmt.Println("service protocol name is tcp")
      fmt.Println(string(p1this.Arr1RecvBuffer[0:p1this.RecvNowLength]))
    }
    p1this.HandleMessageWithProtocol()
  }
}

// 用协议处理接收到的消息
func (p1this *TcpConnection) HandleMessageWithProtocol() {
  arr1CopyBuffer := p1this.Arr1RecvBuffer[0:p1this.RecvNowLength]
  for p1this.RecvNowLength > 0 {
    msgLen, err := p1this.Protocol.DataLength(arr1CopyBuffer)
    if nil != err {
      if config.STR_HTTP == p1this.ProtocolName {
        // 接口转实例
        ptcHttp := p1this.Protocol.(*protocol.Http)
        switch ptcHttp.ParseStatus {
        case protocol.HTTP_STATUS_NO_DATA,
          protocol.HTTP_STATUS_NOT_HTTP,
          protocol.HTTP_STATUS_NOT_FINISH:
          // 继续接收
          break
        case protocol.HTTP_STATUS_WRONG_DATA,
          protocol.HTTP_STATUS_TOO_LONG,
          protocol.HTTP_STATUS_ATOI_ERR:
          // 明显出错
          p1this.CloseTcpCnct()
          break
        }
      }
      // 处理完错误后，跳出for
      break
    }
    // 取出第1条完整的报文
    firstMsg := p1this.Arr1RecvBuffer[0:msgLen]
    switch p1this.ProtocolName {
    case config.STR_HTTP:
      ptcHttp := p1this.Protocol.(*protocol.Http)
      ptcHttp.DataDecode(firstMsg)
      tool.DebugPrintln(ptcHttp)
      p1this.NetService.OnRequest(p1this)
      resp := &http.Response{}
      resp.HandInit()
      respStr := resp.MakeData(200, "hello, world")
      tool.DebugPrintln("respStr=", respStr)
      p1this.Send(respStr)
      break
    case config.STR_STREAM:
      ptcStream := p1this.Protocol.(*protocol.Stream)
      ptcStream.DataDecode(firstMsg)
      tool.DebugPrintln(ptcStream)
      break
    }

    // 计算剩余的数据
    arr1CopyBuffer = arr1CopyBuffer[msgLen:]
    p1this.RecvNowLength -= msgLen
    if p1this.RecvNowLength <= 0 {
      break
    }
  }
}

// 关闭tcp连接
func (p1this *TcpConnection) CloseTcpCnct() {
  p1this.ConnectionRunning = CONNECTION_RUNNING_OFF
  p1this.RecvNowLength = 0
  p1this.Conn.Close()
  p1this.NetService.OnClose(p1this)
  p1this.NetService.DelTcpCnct(p1this)
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
    p1this.NetService.OnError("WriteData(),Write()," + err.Error())
    p1this.CloseTcpCnct()
  }
  if writeByteNum != dataLen {
    err = errors.New("WRITE_BYTE_NOT_EQUAL")
  }
  return
}
