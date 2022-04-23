package websocket

import (
  "crypto/md5"
  "crypto/sha1"
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/protocol/http"
  "encoding/base64"
  "errors"
  "fmt"
  "strings"
)

const (
  handshakeStatusNo  uint8 = iota // 没有握手
  handshakeStatusYes              // 已经握手
)
const (
  opcodeText   uint8 = 0x01 // 文本帧
  opcodeBinary uint8 = 0x02 // 二进制帧
  opcodeClose  uint8 = 0x08 // 连接断开
  opcodePing   uint8 = 0x09 // ping
  opcodePong   uint8 = 0x0A // pong
)

var (
  ErrDataIncomplete     = errors.New("websocket sli1data incomplete.")
  ErrConnectionIsClosed = errors.New("websocket connection is closed.")
)

var _ protocol.Protocol = &WebSocket{}

// WebSocket 协议
// https://www.rfc-editor.org/rfc/rfc6455
type WebSocket struct {
  // 握手阶段要用 HTTP 协议
  p1HttpInner *http.HTTP

  // 握手状态，详见 handshakeStatus 开头的常量
  handshakeStatus uint8

  // FIN，1 bit
  // 0（不是消息的最后一个分片）；1（这是消息的最后一个分片）；
  fin bool
  // opcode，4 bit
  opcode uint8
  // MASK，1 bit
  // 0（没有 Masking-key）；1（有 Masking-key）；
  mask bool
  // Payload len，7 bit
  payloadLen8 uint8
  // Extended payload length，16 bit，if payload len==126
  payloadLen16 uint16
  // Extended payload length，64 bit，if payload len==127
  payloadLen64 uint64
  // Masking-key，4 byte
  arr1MaskingKey [4]byte

  // 头部长度
  headerLength uint8
  // 消息体长度
  bodyLength uint64

  // 请求报文
  Sli1Msg []byte
  // 解析后的数据
  DecodeMsg string

  SecWebSocketKey string
}

func NewWebSocket() *WebSocket {
  return &WebSocket{
    p1HttpInner:     http.NewHTTP(),
    handshakeStatus: handshakeStatusNo,
  }
}

func (p1this *WebSocket) FirstMsgLength(sli1recv []byte) (uint64, error) {
  if handshakeStatusNo == p1this.handshakeStatus {
    // 没有握手
    return p1this.p1HttpInner.FirstMsgLength(sli1recv)
  } else if handshakeStatusYes == p1this.handshakeStatus {
    // 已经握手
    recvLen := len(sli1recv)
    if recvLen < 2 {
      // 至少 2 个字节才能解析
      return 0, ErrDataIncomplete
    }

    // 取 FIN，第 1 个字节的第 1 位
    t1fin := sli1recv[0] & 0b10000000
    if t1fin == 0b10000000 {
      p1this.fin = true
    }

    // 取 opcode，第 1 个字节的后 4 位
    p1this.opcode = sli1recv[0] & 0b00001111
    if opcodeClose == p1this.opcode {
      return 0, ErrConnectionIsClosed
    }

    // 头部长度至少 2 字节
    p1this.headerLength = 2

    // 取 MASK，第 2 个字节的第 1 位
    mask := sli1recv[1] & 0b10000000
    if mask == 0b10000000 {
      p1this.mask = true
      // 有 Masking-key，头部长度 +4 字节
      p1this.headerLength += 4
    } else {
      p1this.mask = false
    }

    // 取 Payload len，第 2 个字节的后 7 位
    p1this.payloadLen8 = sli1recv[1] & 0b0111111
    if 126 == p1this.payloadLen8 {
      p1this.headerLength += 2
    } else if 127 == p1this.payloadLen8 {
      p1this.headerLength += 8
    }

    if int(p1this.headerLength) > recvLen {
      // 计算出来的报文长度大于接收缓冲区中数据长度
      return 0, ErrDataIncomplete
    }

    var msgLen uint64 = 0

    // 计算报文长度
    if 126 == p1this.payloadLen8 {
      // Payload len 为 126，需要扩展 2 个字节
      p1this.payloadLen16 = 0
      p1this.payloadLen16 |= uint16(sli1recv[2]) << 8
      p1this.payloadLen16 |= uint16(sli1recv[3]) << 0
      msgLen = uint64(p1this.headerLength) + uint64(p1this.payloadLen16)
    } else if 127 == p1this.payloadLen8 {
      // Payload len 为 127，需要扩展 8 个字节
      p1this.payloadLen64 |= uint64(sli1recv[2]) << 56
      p1this.payloadLen64 |= uint64(sli1recv[3]) << 48
      p1this.payloadLen64 |= uint64(sli1recv[4]) << 40
      p1this.payloadLen64 |= uint64(sli1recv[5]) << 32
      p1this.payloadLen64 |= uint64(sli1recv[6]) << 24
      p1this.payloadLen64 |= uint64(sli1recv[7]) << 16
      p1this.payloadLen64 |= uint64(sli1recv[8]) << 8
      p1this.payloadLen64 |= uint64(sli1recv[9]) << 0
      msgLen = uint64(p1this.headerLength) + uint64(p1this.payloadLen64)
    } else {
      msgLen = uint64(p1this.headerLength) + uint64(p1this.payloadLen8)
    }

    p1this.bodyLength = msgLen
    if msgLen > uint64(recvLen) {
      // 计算出来的报文长度大于接收缓冲区中数据长度
      return 0, ErrDataIncomplete
    }

    // 获取 MaskingKey
    if p1this.mask {
      p1this.arr1MaskingKey[0] = sli1recv[p1this.headerLength-4]
      p1this.arr1MaskingKey[1] = sli1recv[p1this.headerLength-3]
      p1this.arr1MaskingKey[2] = sli1recv[p1this.headerLength-2]
      p1this.arr1MaskingKey[3] = sli1recv[p1this.headerLength-1]
    }

    return msgLen, nil
  }

  return 0, nil
}

func (p1this *WebSocket) Decode(sli1msg []byte) error {
  if handshakeStatusNo == p1this.handshakeStatus {
    // 没有握手
    return p1this.p1HttpInner.Decode(sli1msg)
  } else if handshakeStatusYes == p1this.handshakeStatus {
    // 已经握手
    p1this.Sli1Msg = sli1msg
    msgLen := uint64(len(sli1msg))
    t1sli1msg := make([]byte, msgLen-uint64(p1this.headerLength))
    // 头部不需要解析，只解析数据部分
    // 解析的时候，4 个 Masking-key 轮着用
    // 第 1 个字节和第 1 个 Masking-key 异或
    // 第 2 个字节和第 2 个 Masking-key 异或
    // 第 3 个字节和第 3 个 Masking-key 异或
    // 第 4 个字节和第 4 个 Masking-key 异或
    // 第 5 个字节和第 1 个 Masking-key 异或
    var i, j uint64 = 0, uint64(p1this.headerLength)
    for j < uint64(msgLen) {
      t1sli1msg[i] = p1this.Sli1Msg[j] ^ p1this.arr1MaskingKey[i&0b00000011]
      i++
      j++
    }
    p1this.DecodeMsg = string(t1sli1msg)
  }
  return nil
}

func (p1this *WebSocket) Encode() ([]byte, error) {
  return []byte{}, nil
}

func (p1this *WebSocket) IsHandshakeStatusNo() bool {
  return handshakeStatusNo == p1this.handshakeStatus
}

func (p1this *WebSocket) IsHandshakeStatusYes() bool {
  return handshakeStatusYes == p1this.handshakeStatus
}

func (p1this *WebSocket) SetHandshakeStatusYes() {
  p1this.handshakeStatus = handshakeStatusYes
}

func (p1this *WebSocket) Handshake() ([]byte, error) {
  var sli1msg []byte = []byte{}

  // 判断请求头中 connection 和 upgrade 字段是否符合要求
  connection, ok := p1this.p1HttpInner.MapHeader["connection"]
  if !ok {
    return sli1msg, errors.New("请求报文connection字段不存在")
  }
  upgrade, ok := p1this.p1HttpInner.MapHeader["upgrade"]
  if !ok {
    return sli1msg, errors.New("请求报文upgrade字段不存在")
  }
  upgradeIndex := strings.Index(connection, "Upgrade")
  websocketIndex := strings.Index(upgrade, "websocket")
  if upgradeIndex < 0 || websocketIndex < 0 {
    return sli1msg, errors.New("握手失败请重试")
  }

  secWebSocketKey, ok := p1this.p1HttpInner.MapHeader["sec-websocket-key"]
  if !ok {
    return sli1msg, errors.New("请求报文sec-webSocket-key字段不存在")
  }

  // Sec-WebSocket-Accept
  // 将 Sec-WebSocket-Key 跟 258EAFA5-E914-47DA-95CA-C5AB0DC85B11 拼接
  acceptSumArg := secWebSocketKey + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
  // 通过 SHA1 计算摘要
  acceptKeySha := sha1.Sum([]byte(acceptSumArg))
  // 转成 base64 字符串
  acceptKeyStr := base64.StdEncoding.EncodeToString(acceptKeySha[:])

  // 这几个字段最好都有，少一个都跑不通
  // 测试使用的是 JavaScript 的 WebSocket 工具
  msg := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\n")
  msg += fmt.Sprintf("Connection: Upgrade\r\n")
  msg += fmt.Sprintf("Upgrade: websocket\r\n")
  msg += fmt.Sprintf("Sec-WebSocket-Accept: %v\r\n", acceptKeyStr)
  msg += fmt.Sprintf("Sec-WebSocket-Version: 13\r\n")
  msg += fmt.Sprintf("Server: tcp_server_v1\r\n\r\n")

  sli1msg = []byte(msg)

  return sli1msg, nil
}

func (p1this *WebSocket) HandShakeClient() (msg []byte, err error) {

  nonstr := "bf"
  key := md5.Sum([]byte(nonstr))
  secWebSocketKey := base64.StdEncoding.EncodeToString(key[:])
  p1this.SecWebSocketKey = secWebSocketKey

  text := fmt.Sprintf("GET /chat HTTP/1.1\r\n")
  text += fmt.Sprintf("Upgrade: websocket\r\n")
  text += fmt.Sprintf("Connection: Upgrade\r\n")
  text += fmt.Sprintf("Sec-WebSocket-Key: %v\r\n", secWebSocketKey)
  text += fmt.Sprintf("Sec-WebSocket-Version: 13\r\n\r\n")

  msg = []byte(text)

  return
}

func (this *WebSocket) VerifyShakeHand() (err error) {

  connection, ok := this.p1HttpInner.MapHeader["connection"]
  if !ok {
    err = errors.New("响应报文connection字段不存在")
    return
  }
  upgrade, ok := this.p1HttpInner.MapHeader["upgrade"]
  if !ok {
    err = errors.New("响应报文upgrade字段不存在")
    return
  }

  secWebsocketAccept, ok := this.p1HttpInner.MapHeader["sec-websocket-accept"]
  if !ok {
    err = errors.New("响应报文sec-websocket-accept字段不存在")
    return
  }

  if connection != "Upgrade" || upgrade != "websocket" {
    err = errors.New("握手失败请重试")
    return
  }

  acceptSumArg := this.SecWebSocketKey + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
  acceptKeySha := sha1.Sum([]byte(acceptSumArg))
  acceptKeyStr := base64.StdEncoding.EncodeToString(acceptKeySha[:])

  if acceptKeyStr != secWebsocketAccept {
    err = errors.New("握手失败请重试")
    return
  }

  return
}
