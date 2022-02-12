package service

import (
  "demo_golang/net_service/helper"
  "demo_golang/net_service/protocol"
  "fmt"
  "net"
)

const (
  RecvBufferMaxLength = 1048576 // 接收缓冲区最大大小，1024*1024
)

// tcp连接结构体
type TcpConnection struct {
  NetService        *NetService               // 服务结构体实例地址
  ConnectionRunning int                       // 连接运行状态，1=连接运行
  Protocol          protocol.Protocol         // 协议接口实例
  Conn              net.Conn                  // net.Conn实例
  Arr1RecvBuffer    [RecvBufferMaxLength]byte // 接收缓冲区
  RecvNowLength     int                       // 接收缓冲区当前大小
}

// 构造tcp连接结构体
func MakeTcpConnection(p1NetSvc *NetService, conn net.Conn) (p1TcpCnct *TcpConnection, err error) {
  var ptc protocol.Protocol
  switch p1NetSvc.ProtocolName {
  case "http":
    ptc = &protocol.Http{}
    break
  }
  p1TcpCnct = &TcpConnection{
    NetService:        p1NetSvc,
    ConnectionRunning: 1,
    Protocol:          ptc,
    Conn:              conn,
    RecvNowLength:     0,
  }
  return
}

// 处理tcp连接
func (p1this *TcpConnection) HandleTcpConnection() {
  for 1 == p1this.ConnectionRunning {
    recvByteNum, err := p1this.Conn.Read(p1this.Arr1RecvBuffer[p1this.RecvNowLength:])
    helper.PrintlnInDebug(p1this.NetService.AppDebug, "TcpConnection,recvByteNum=", recvByteNum)
    p1this.RecvNowLength += recvByteNum
    helper.PrintlnInDebug(p1this.NetService.AppDebug, "TcpConnection,Arr1RecvLength=", p1this.RecvNowLength)
    if nil != err {
      p1this.NetService.OnError("TcpConnection," + err.Error())
      return
    }
    // helper.PrintlnInDebug(p1this.NetService.AppDebug, "TcpConnection,Arr1RecvBuffer=", p1this.Arr1RecvBuffer[0:p1this.RecvNowLength])
    // helper.PrintlnInDebug(p1this.NetService.AppDebug, "TcpConnection,string(Arr1RecvBuffer)=", string(p1this.Arr1RecvBuffer[0:p1this.RecvNowLength]))

    if "tcp" == p1this.NetService.ProtocolName {
      fmt.Println("read from tcp:")
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
      if "http" == p1this.NetService.ProtocolName {
        // 接口转实例
        ptcHttp := p1this.Protocol.(*protocol.Http)
        switch ptcHttp.Status {
        case protocol.HTTP_STATUS_NO_DATA:
        case protocol.HTTP_STATUS_NOT_HTTP:
        case protocol.HTTP_STATUS_NOT_FINISH:
          // 继续接收
          break
        case protocol.HTTP_STATUS_WRONG_DATA:
        case protocol.HTTP_STATUS_TOO_LONG:
        case protocol.HTTP_STATUS_ATOI_ERR:
          // 明显出错
          p1this.DelTcpCnct()
          break
        }
      }
      // 跳出for
      break
    }
    // 取出第1条完整的报文
    firstMsg := p1this.Arr1RecvBuffer[0:msgLen]
    switch p1this.NetService.ProtocolName {
    case "http":
      ptcHttp := p1this.Protocol.(*protocol.Http)
      ptcHttp.DataDecode(firstMsg)
      fmt.Println(ptcHttp)
    }

    // 计算剩余的数据
    arr1CopyBuffer = arr1CopyBuffer[msgLen:]
    p1this.RecvNowLength -= msgLen
    if p1this.RecvNowLength <= 0 {
      break
    }
  }
}

func (p1this *TcpConnection) DelTcpCnct() {
  p1this.ConnectionRunning = 0
  p1this.RecvNowLength = 0
  p1this.Conn.Close()
  p1this.NetService.OnClose(p1this)
  p1this.NetService.DelTcpCnct(p1this)
}
