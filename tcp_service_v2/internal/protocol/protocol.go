package protocol

const (
  StrTCP       = "tcp"
  StrHTTP      = "http"
  StrStream    = "stream"
  StrWebSocket = "websocket"
)

// Protocol 协议
type Protocol interface {
  // FirstMsgLength 计算接收缓冲区中第 1 个完整的报文的长度
  FirstMsgLength(sli1recv []byte) (uint64, error)
  // Decode 报文解码
  Decode(sli1msg []byte) error
  GetDecodeMsg() string
  // Encode 报文编码
  SetDecodeMsg(msg string)
  Encode() ([]byte, error)
}
