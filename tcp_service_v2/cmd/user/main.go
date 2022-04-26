package main

import (
	"demo_golang/tcp_service_v2"
	"demo_golang/tcp_service_v2/internal/client"
	"demo_golang/tcp_service_v2/internal/protocol"
	"demo_golang/tcp_service_v2/internal/tool/signal"
	"demo_golang/tcp_service_v2/internal/user"
	"fmt"
	"log"
)

var p1client *client.TCPClient

func main() {
  log.Println("version: ", tcp_service_v2.Version)

  p1client := client.NewTCPClient(protocol.StreamStr, "127.0.0.1", 9501)
  p1client.SetName(fmt.Sprintf("%s-client-user", protocol.StreamStr))
  p1client.SetDebugStatusOn()

  p1client.OnConnConnect = func(p1connection *client.TCPConnection) {
    if p1client.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1client.GetName()))
    }
    user.P1UserService.SetClient(p1connection)
    user.P1UserService.RegisteServiceProvider()
  }

  p1client.Start()

  signal.WaitForShutdown()
}
