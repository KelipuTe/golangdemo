package service

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"log"
	"testing"
)

func Test_Service_HTTP(t *testing.T) {
	log.Println("version: ", config.Version)

	p1service := NewTCPService(config.HTTPStr, "127.0.0.1", 9501)
	p1service.SetName(fmt.Sprintf("%s-service", config.HTTPStr))
	p1service.OpenDebug()
	p1service.Start()

	signal.WaitForShutdown()
}

func Test_Service_Stream(t *testing.T) {
	log.Println("version: ", config.Version)

	p1service := NewTCPService(config.StreamStr, "127.0.0.1", 9501)
	p1service.SetName(fmt.Sprintf("%s-service", config.StreamStr))
	p1service.OpenDebug()
	p1service.Start()

	signal.WaitForShutdown()
}

// 可以用 EasySwoole-WebSocket在线测试工具 测试
// http://www.easyswoole.com/wstool.html
// 也可以直接用 JavaScript 的 WebSocket 工具
func Test_Service_WebSocket(t *testing.T) {
	log.Println("version: ", config.Version)

	p1service := NewTCPService(config.WebSocketStr, "127.0.0.1", 9501)
	p1service.SetName(fmt.Sprintf("%s-service", config.WebSocketStr))
	p1service.OpenDebug()
	p1service.Start()

	signal.WaitForShutdown()
}
