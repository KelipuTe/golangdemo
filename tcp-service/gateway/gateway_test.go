package gateway

import (
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/service"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"log"
	"testing"
)

func Test_Gateway(t *testing.T) {
	gateway := &Gateway{
		name:              defaultName,
		debugStatus:       config.DebugStatusOff,
		mapInnerConnPool:  make(map[string][]*service.TCPConnection),
		mapInnerConnCount: make(map[string]uint64),
		mapConnToPing:     make(map[string]*service.TCPConnection),
		mapOpenConn:       make(map[string]*service.TCPConnection),
	}

	log.Println("version: ", config.Version)

	gateway.SetDebugStatusOn()

	innerService := service.NewTCPService(config.StreamStr, "127.0.0.1", 9501)
	innerService.SetName(fmt.Sprintf("%s-service-gateway", config.StreamStr))
	innerService.OpenDebug()

	innerService.OnServiceStart = func(p1service *service.TCPService) {
		if p1service.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1service.GetName()))
		}
		gateway.SetInnerService(p1service)
		go gateway.StartPingConn()
	}

	innerService.OnConnRequest = func(p1conn *service.TCPConnection) {
		if innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnGetRequest", innerService.GetName()))
		}
		gateway.DispatchInnerRequest(p1conn)
	}

	innerService.OnConnClose = func(p1conn *service.TCPConnection) {
		if innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.AfterConnClose", innerService.GetName()))
		}
		gateway.DeleteServiceProvider(p1conn)
	}

	go innerService.Start()

	openService := service.NewTCPService(config.HTTPStr, "127.0.0.1", 9502)
	openService.SetName(fmt.Sprintf("%s-service-gateway", config.HTTPStr))
	openService.OpenDebug()

	openService.OnConnRequest = func(p1conn *service.TCPConnection) {
		if innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", innerService.GetName()))
		}
		gateway.DispatchOpenRequest(p1conn)
	}

	go openService.Start()

	signal.WaitForShutdown()
}
