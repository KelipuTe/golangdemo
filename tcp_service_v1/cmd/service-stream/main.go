package main

import (
  "demo_golang/tcp_service_v1"
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/service"
  "demo_golang/tcp_service_v1/internal/tool/debug"
  "fmt"
  "log"
)

func main() {
  log.Println("version: ", tcp_service_v1.Version)

  p1service := service.NewTCPService(protocol.StrStream, "127.0.0.1", 9501)
  p1service.OnStart = MyOnStart
  p1service.OnError = MyOnError
  p1service.OnConnect = MyOnConnect
  p1service.OnRequest = MyOnRequest
  p1service.OnClose = MyOnClose

  p1service.Start()
}

func MyOnStart(p1service *service.TCPService) {
  debug.Println(p1service.IsDebug(), "TCPService.OnStart")
}

func MyOnError(p1service *service.TCPService, err error) {
  debug.Println(p1service.IsDebug(), "TCPService.OnError")
  fmt.Println(fmt.Sprintf("%s", err))
}

func MyOnConnect(p1service *service.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPService.MyOnConnect")
}

func MyOnRequest(p1service *service.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPService.MyOnRequest")
}

func MyOnClose(p1service *service.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPService.MyOnClose")
}
