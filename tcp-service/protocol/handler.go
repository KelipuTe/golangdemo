package protocol

import (
	"demo-golang/tcp-service/config"
	"sync"
)

var supportProtocol map[string]bool //支持的协议
var supportProtocolLock sync.RWMutex

func init() {
	supportProtocol = map[string]bool{
		config.TCPStr:       true,
		config.HTTPStr:      true,
		config.StreamStr:    true,
		config.WebSocketStr: true,
	}
}

// IsSupported 判断是否支持某个协议
func IsSupported(name string) bool {
	supportProtocolLock.RLock()
	defer supportProtocolLock.RUnlock()
	_, ok := supportProtocol[name]
	return ok
}

// HandlerI9 协议处理器
type HandlerI9 interface {
	// FirstMsgLen 计算缓冲区中第1个完整的报文的长度
	FirstMsgLen(buffer []byte) (uint64, error)
	// Decode 报文解码
	Decode(msg []byte) error
	// Encode 报文编码
	Encode() ([]byte, error)
}
