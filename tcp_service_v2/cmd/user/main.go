package main

import (
  "demo_golang/tcp_service_v2"
  "demo_golang/tcp_service_v2/internal/client"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/tool/debug"
  "demo_golang/tcp_service_v2/internal/user"
  "fmt"
  "log"
)

var p1client *client.TCPClient

func main() {
  log.Println("version: ", tcp_service_v2.Version)

  p1client = client.NewTCPClient(protocol.StrStream, "127.0.0.1", 9501)
  p1client.OnStart = MyOnStart
  p1client.OnError = MyOnError
  p1client.OnConnect = MyOnConnect
  p1client.OnRequest = MyOnRequest
  p1client.OnClose = MyOnClose

  p1client.Start()
}

func MyOnStart(p1service *client.TCPClient) {
  debug.Println(p1service.IsDebug(), "TCPService.OnStart")
}

func MyOnError(p1service *client.TCPClient, err error) {
  debug.Println(p1service.IsDebug(), "TCPService.OnError")
  fmt.Println(fmt.Sprintf("%s", err))
}

func MyOnConnect(p1connection *client.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnConnect")
  user.P1UserService.SetClient(p1connection)
  user.P1UserService.RegisterService()
}

func MyOnRequest(p1connection *client.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnRequest")
}

func MyOnClose(p1connection *client.TCPConnection) {
  debug.Println(p1connection.IsDebug(), "TCPConnection.MyOnClose")
}
