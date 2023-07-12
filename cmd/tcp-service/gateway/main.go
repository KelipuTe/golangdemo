package main

import (
	tcp_service "demo-golang/tcp-service"
	"demo-golang/tcp-service/gateway"
	"demo-golang/tcp-service/protocol"
	"demo-golang/tcp-service/service"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"log"
)

var p1innerService *service.TCPService
var p1openService *service.TCPService

func main() {
	log.Println("version: ", tcp_service.Version)

	gateway.P1gateway.SetDebugStatusOn()

	p1innerService := service.NewTCPService(protocol.StreamStr, "127.0.0.1", 9501)
	p1innerService.SetName(fmt.Sprintf("%s-service-gateway", protocol.StreamStr))
	p1innerService.SetDebugStatusOn()

	p1innerService.OnServiceStart = func(p1service *service.TCPService) {
		if p1service.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1service.GetName()))
		}
		gateway.P1gateway.SetInnerService(p1service)
		go gateway.P1gateway.StartPingConn()
	}

	p1innerService.OnConnRequest = func(p1conn *service.TCPConnection) {
		if p1innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnRequest", p1innerService.GetName()))
		}
		gateway.P1gateway.DispatchInnerRequest(p1conn)
	}

	p1innerService.OnConnClose = func(p1conn *service.TCPConnection) {
		if p1innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnClose", p1innerService.GetName()))
		}
		gateway.P1gateway.DeleteServiceProvider(p1conn)
	}

	go p1innerService.Start()

	p1openService := service.NewTCPService(protocol.HTTPStr, "127.0.0.1", 9502)
	p1openService.SetName(fmt.Sprintf("%s-service-gateway", protocol.HTTPStr))
	p1openService.SetDebugStatusOn()

	p1openService.OnConnRequest = func(p1conn *service.TCPConnection) {
		if p1innerService.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1innerService.GetName()))
		}
		gateway.P1gateway.DispatchOpenRequest(p1conn)
	}

	go p1openService.Start()

	signal.WaitForShutdown()
}
