package websocket

import "errors"

var (
	ErrParseFailed = errors.New("解析 WebSocket 消息报文失败")
)
