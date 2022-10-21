package main

import (
  "encoding/binary"
  "fmt"
  "io"
  "net"
)

func main() {
  // 启动监听
  p1listener, err := net.Listen("tcp4", "127.0.0.1:9502")
  if nil != err {
    fmt.Println("net.Listen err=", err.Error())
    return
  }
  defer p1listener.Close()

  // 简单的就无限循环等待连接
  for {
    // 获取连接
    p1conn, err := p1listener.Accept()
    if nil != err {
      fmt.Println("net.Listen.Accept err=", err.Error())
      return
    }
    fmt.Println("net.Conn.RemoteAddr.String", p1conn.RemoteAddr().String())

    // 丢出去处理，主线程只负责建立连接
    go handleConn(p1conn)
  }
}

func handleConn(p1conn net.Conn) {
  // 接收缓冲区
  var sli1recvBuffer []byte
  sli1recvBuffer = make([]byte, 10240)
  // 接收缓冲区当前大小
  var recvBufferNow uint32
  recvBufferNow = 0

  // 持续读取数据
  for {
    // 读数据到缓冲区
    byteNum, err := p1conn.Read(sli1recvBuffer[recvBufferNow:])
    recvBufferNow += uint32(byteNum)
    fmt.Println("net.Conn.Read byteNum=", byteNum)

    if nil != err {
      if err == io.EOF {
        // 对端关闭了连接
        p1conn.Close()
        break
      }
    }

    for {
      // 获取第一条消息
      if recvBufferNow < 4 {
        // 消息没接收全
        break
      }
      packageLength := binary.BigEndian.Uint32(sli1recvBuffer[0:4])
      if recvBufferNow < packageLength {
        // 消息没接收全
        break
      }

      // 解析消息
      headerLength := binary.BigEndian.Uint16(sli1recvBuffer[4:6])
      protocolVersion := binary.BigEndian.Uint16(sli1recvBuffer[6:8])
      operation := binary.BigEndian.Uint32(sli1recvBuffer[8:12])
      sequenceId := binary.BigEndian.Uint32(sli1recvBuffer[12:16])
      sli1Body := sli1recvBuffer[headerLength:packageLength]
      body := string(sli1Body)

      fmt.Println("decode:")
      fmt.Println("packageLength=", packageLength)
      fmt.Println("headerLength=", headerLength)
      fmt.Println("protocolVersion=", protocolVersion)
      fmt.Println("operation=", operation)
      fmt.Println("sequenceId=", sequenceId)
      fmt.Println("body=", body)

      // 清理缓冲区
      sli1recvBuffer = sli1recvBuffer[packageLength:]
      recvBufferNow -= packageLength
    }
  }
}
