package main

import (
  "demo_golang/tcp_service_v1"
  "demo_golang/tcp_service_v1/internal/client"
  "demo_golang/tcp_service_v1/internal/protocol"
  "demo_golang/tcp_service_v1/internal/tool/debug"
  "fmt"
  "log"
)

func main() {
  log.Println("version: ", tcp_service_v1.Version)

  p1client := client.NewTCPClient(protocol.StrStream, "127.0.0.1", 9501)
  p1client.OnStart = MyOnStart
  p1client.OnError = MyOnError
  p1client.OnConnect = MyOnConnect
  p1client.OnRequest = MyOnRequest
  p1client.OnClose = MyOnClose

  p1client.Start()
}

func MyOnStart(p1service *client.TCPClient) {
  debug.Println(p1service.IsDebug(), "TCPClient.OnStart")
}

func MyOnError(p1service *client.TCPClient, err error) {
  debug.Println(p1service.IsDebug(), "TCPClient.OnError")
  fmt.Println(fmt.Sprintf("%s", err))
}

func MyOnConnect(p1service *client.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPClient.MyOnConnect")
}

func MyOnRequest(p1service *client.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPClient.MyOnRequest")
}

func MyOnClose(p1service *client.TCPConnection) {
  debug.Println(p1service.IsDebug(), "TCPClient.MyOnClose")
}
