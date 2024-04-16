package service

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"testing"
)

func Test_Service_HTTP(t *testing.T) {
	service := NewTCPService(config.HTTPStr, "127.0.0.1", 9501)
	service.SetName(fmt.Sprintf("%s-service", config.HTTPStr))
	service.OpenDebug()
	service.Start()
	signal.WaitForShutdown()
}

func Test_Service_Stream(t *testing.T) {
	service := NewTCPService(config.StreamStr, "127.0.0.1", 9502)
	service.SetName(fmt.Sprintf("%s-service", config.StreamStr))
	service.OpenDebug()
	service.Start()
	signal.WaitForShutdown()
}

// 可以用 EasySwoole-WebSocket在线测试工具
// http://www.easyswoole.com/wstool.html
// 也可以直接用 JavaScript 的 WebSocket 工具
func Test_Service_WebSocket(t *testing.T) {
	service := NewTCPService(config.WebSocketStr, "127.0.0.1", 9503)
	service.SetName(fmt.Sprintf("%s-service", config.WebSocketStr))
	service.OpenDebug()
	service.Start()
	signal.WaitForShutdown()
}
