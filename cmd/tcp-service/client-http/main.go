package main

import (
	tcp_service "demo-golang/tcp-service"
	"demo-golang/tcp-service/client"
	"demo-golang/tcp-service/protocol"
	"fmt"
	"log"
)

func main() {
	log.Println("version: ", tcp_service.Version)

	p1client := client.NewTCPClient(protocol.HTTPStr, "127.0.0.1", 9501)
	p1client.SetName(fmt.Sprintf("%s-client", protocol.HTTPStr))
	p1client.SetDebugStatusOn()
	p1client.Start()
}
