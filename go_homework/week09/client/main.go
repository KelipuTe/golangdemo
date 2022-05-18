package main

import (
  "encoding/binary"
  "fmt"
  "net"
)

func main() {
  p1conn, err := net.Dial("tcp4", "127.0.0.1:9502")
  if nil != err {
    fmt.Println("net.Dial err=", err.Error())
    return
  }
  // 这里发完一次就关闭
  defer p1conn.Close()

  // ppt里没看到 goim 详细的协议说明，这里就假定是都是大端字节序

  var sli1msg []byte
  t1sli1packageLength := make([]byte, 4)

  t1sli1headerLength := make([]byte, 2)
  binary.BigEndian.PutUint16(t1sli1headerLength, 16)
  sli1msg = append(sli1msg, t1sli1headerLength...)

  t1sli1protocolVersion := make([]byte, 2)
  binary.BigEndian.PutUint16(t1sli1protocolVersion, 1)
  sli1msg = append(sli1msg, t1sli1protocolVersion...)

  t1sli1operation := make([]byte, 4)
  binary.BigEndian.PutUint32(t1sli1operation, 10)
  sli1msg = append(sli1msg, t1sli1operation...)

  t1sli1sequenceId := make([]byte, 4)
  binary.BigEndian.PutUint32(t1sli1sequenceId, 100)
  sli1msg = append(sli1msg, t1sli1sequenceId...)

  body := "{\"goim\":\"msg\"}"
  sli1Body := []byte(body)

  binary.BigEndian.PutUint32(t1sli1packageLength, uint32(len(sli1Body)+16))
  sli1msg = append(t1sli1packageLength, sli1msg...)
  sli1msg = append(sli1msg, body...)

  // 复制一遍，发两次，模拟粘包
  sli1msg = append(sli1msg, sli1msg...)

  byteNum, err := p1conn.Write(sli1msg)
  fmt.Println("net.Conn.Write byteNum=", byteNum)
}
