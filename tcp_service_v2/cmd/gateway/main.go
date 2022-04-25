package main

import (
  "demo_golang/tcp_service_v2"
  "demo_golang/tcp_service_v2/internal/gateway"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/service"
  "demo_golang/tcp_service_v2/internal/tool/debug"
  "fmt"
  "log"
)

var p1service *service.TCPService

var p1linkservice *service.TCPService

func main() {
  log.Println("version: ", tcp_service_v2.Version)

  p1service = service.NewTCPService(protocol.StrStream, "127.0.0.1", 9501)
  p1service.OnStart = MyOnStart
  p1service.OnError = MyOnError
  p1service.OnConnect = MyOnConnect
  p1service.OnRequest = MyOnRequest
  p1service.OnClose = MyOnClose

  go p1service.Start()

  p1linkservice = service.NewTCPService(protocol.StrHTTP, "127.0.0.1", 9502)
  p1linkservice.OnStart = MyOnStart2
  p1linkservice.OnError = MyOnError
  p1linkservice.OnConnect = MyOnConnect
  p1linkservice.OnRequest = MyOnRequest2
  p1linkservice.OnClose = MyOnClose

  p1linkservice.Start()
}

func MyOnStart(p1service *service.TCPService) {
  debug.Println(p1service.IsDebug(), "TCPService.OnStart")
  gateway.P1gateway.SetService(p1service)
}

func MyOnStart2(p1service *service.TCPService) {
  debug.Println(p1service.IsDebug(), "TCPService.OnStart")
}

func MyOnError(p1service *service.TCPService, err error) {
  debug.Println(p1service.IsDebug(), "TCPService.OnError")
  fmt.Println(fmt.Sprintf("%s", err))
}

func MyOnConnect(p1connection *service.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnConnect")
}

func MyOnRequest(p1connection *service.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnRequest")
  gateway.P1gateway.RegisteService(p1connection)
}

func MyOnRequest2(p1connection *service.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnRequest")
  gateway.P1gateway.GetConn("user_service").SendMsg(string("aaa"))
}

func MyOnClose(p1connection *service.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnClose")
}
