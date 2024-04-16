package abs

// HandlerI9 协议处理器
type HandlerI9 interface {
	// FirstMsgLen 计算缓冲区中第1个完整的报文的长度
	FirstMsgLen(buffer []byte) (uint64, error)
	// Decode 报文解码
	Decode(msg []byte) error
	// Encode 报文编码
	Encode() ([]byte, error)
}
