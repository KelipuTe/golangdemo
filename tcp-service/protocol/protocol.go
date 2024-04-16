package protocol

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/protocol/abs"
	"demo-golang/tcp-service/protocol/http"
	"demo-golang/tcp-service/protocol/stream"
	"demo-golang/tcp-service/protocol/websocket"
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

func NewHandler(protocolName string) abs.HandlerI9 {
	switch protocolName {
	case config.HTTPStr:
		return http.NewHandler()
	case config.StreamStr:
		return stream.NewStream()
	case config.WebSocketStr:
		return websocket.NewWebSocket()
	}
	return nil
}
