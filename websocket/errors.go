package websocket

import "errors"

var (
	ErrParseFailed = errors.New("WebSocket 报文解析失败")
)
