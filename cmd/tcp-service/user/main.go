package main

import (
	tcp_service "demo-golang/tcp-service"
	"demo-golang/tcp-service/client"
	"demo-golang/tcp-service/protocol"
	"demo-golang/tcp-service/tool/signal"
	"demo-golang/tcp-service/user"
	"fmt"
	"log"
)

var p1innerClient *client.TCPClient

func main() {
	log.Println("version: ", tcp_service.Version)

	p1innerClient := client.NewTCPClient(protocol.StreamStr, "127.0.0.1", 9501)
	p1innerClient.SetName(fmt.Sprintf("%s-client-user", protocol.StreamStr))
	p1innerClient.SetDebugStatusOn()

	p1innerClient.OnClientStart = func(p1client *client.TCPClient) {
		if p1client.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1client.GetName()))
		}
		user.P1UserService.SetInnerClient(p1client)
	}

	p1innerClient.OnConnConnect = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1innerClient.GetName()))
		}
		user.P1UserService.RegisteServiceProvider()
	}

	p1innerClient.OnConnRequest = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1innerClient.GetName()))
		}
		user.P1UserService.DispatchRequest(p1conn)
	}
	p1innerClient.Start()

	signal.WaitForShutdown()
}
