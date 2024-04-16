package config

type RunStatus uint8

// 服务（tcp连接）状态
const (
	RunStatusOff RunStatus = iota //关闭
	RunStatusOn                   //运行中
)

type DebugStatus uint8

// debug开关
const (
	DebugStatusOff DebugStatus = iota //关
	DebugStatusOn                     //开
)

// RecvBufferMaxLen 接收缓冲区最大大小。1MB == 2^20 == 1048576。
// uint32，最大 2^32-1，差不多 4GB，理论上应该够用了，uint64 只会更大。
// 这里用 uint64 是因为 WebSocket 协议最大负荷可以是 2^64-1 个字节。
const RecvBufferMaxLen uint64 = 10 * 1048576

// 协议名称
const (
	TCPStr       string = "tcp"
	HTTPStr      string = "http"
	StreamStr    string = "stream"
	WebSocketStr string = "websocket"
)

// content-type
const (
	ContentTypeFormStr = "application/x-www-form-urlencoded"
	ContentTypeJsonStr = "application/json"
)
