package main

import (
  "demo_golang/tcp_service_v2"
  "demo_golang/tcp_service_v2/internal/gateway"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/service"
  "demo_golang/tcp_service_v2/internal/tool/signal"
  "fmt"
  "log"
)

var p1service *service.TCPService

var p1linkservice *service.TCPService

func main() {
  log.Println("version: ", tcp_service_v2.Version)

  p1service := service.NewTCPService(protocol.StreamStr, "127.0.0.1", 9501)
  p1service.SetName(fmt.Sprintf("%s-service-gateway", protocol.StreamStr))
  p1service.SetDebugStatusOn()

  p1service.OnServiceStart = func(p1service *service.TCPService) {
    if p1service.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1service.GetName()))
    }
    gateway.P1gateway.SetService(p1service)
  }

  p1service.OnConnRequest = func(p1connection *service.TCPConnection) {
    if p1service.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.OnConnRequest", p1service.GetName()))
    }
    gateway.P1gateway.DispatchInnerRequest(p1connection)
  }

  go p1service.Start()

  p1linkservice := service.NewTCPService(protocol.HTTPStr, "127.0.0.1", 9502)
  p1linkservice.SetName(fmt.Sprintf("%s-service-gateway", protocol.HTTPStr))
  p1linkservice.SetDebugStatusOn()

  p1linkservice.OnConnRequest = func(p1connection *service.TCPConnection) {
    if p1service.IsDebug() {
      fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1service.GetName()))
    }
    gateway.P1gateway.DispatchOpenRequest(p1connection)
  }

  go p1linkservice.Start()

  signal.WaitForShutdown()
}
