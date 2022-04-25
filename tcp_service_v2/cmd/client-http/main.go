package main

import (
  "demo_golang/tcp_service_v2"
  "demo_golang/tcp_service_v2/internal/client"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/tool/debug"
  "fmt"
  "log"
)

func main() {
  log.Println("version: ", tcp_service_v2.Version)

  p1client := client.NewTCPClient(protocol.StrHTTP, "127.0.0.1", 9501)
  p1client.OnStart = MyOnStart
  p1client.OnError = MyOnError
  p1client.OnConnect = MyOnConnect
  p1client.OnRequest = MyOnRequest
  p1client.OnClose = MyOnClose

  p1client.Start()
}

func MyOnStart(p1client *client.TCPClient) {
  debug.Println(p1client.IsDebug(), "TCPService.OnStart")
}

func MyOnError(p1client *client.TCPClient, err error) {
  debug.Println(p1client.IsDebug(), "TCPService.OnError")
  fmt.Println(fmt.Sprintf("%s", err))
}

func MyOnConnect(p1client *client.TCPConnection) {
  debug.Println(p1client.IsDebug(), "TCPService.MyOnConnect")
}

func MyOnRequest(p1client *client.TCPConnection) {
  debug.Println(p1client.IsDebug(), "TCPService.MyOnRequest")
}

func MyOnClose(p1client *client.TCPConnection) {
  debug.Println(p1client.IsDebug(), "TCPService.MyOnClose")
}
