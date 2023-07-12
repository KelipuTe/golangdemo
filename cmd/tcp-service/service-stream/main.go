package main

import (
	tcp_service "demo-golang/tcp-service"
	"demo-golang/tcp-service/protocol"
	"demo-golang/tcp-service/service"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"log"
)

func main() {
	log.Println("version: ", tcp_service.Version)

	p1service := service.NewTCPService(protocol.StreamStr, "127.0.0.1", 9501)
	p1service.SetName(fmt.Sprintf("%s-service", protocol.StreamStr))
	p1service.SetDebugStatusOn()
	p1service.Start()

	signal.WaitForShutdown()
}
