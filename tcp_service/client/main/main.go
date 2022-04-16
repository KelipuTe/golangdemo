package main

import (
  "demo_golang/net_service/client"
  "demo_golang/net_service/config"
  "fmt"
)

func main() {
  netClt := client.NetClient{
    ProtocolName: config.STR_STREAM,
    Address:      config.ADDRESS,
    Port:         config.PORT,

    OnStart:   MyOnStart,
    OnError:   MyOnError,
    OnConnect: MyOnConnect,
    OnRequest: MyOnRequest,
    OnClose:   MyOnClose,
  }
  netClt.Start()
}

func MyOnStart(p1NetClt *client.NetClient) {
  fmt.Println("NetService.OnStart")
}

func MyOnError(errStr string) {
  fmt.Println("NetService.OnError")
  fmt.Println("errStr=", errStr)
}

func MyOnConnect(p1TcpCnct *client.TcpConnection) {
  fmt.Println("NetService.MyOnConnect")
}

func MyOnRequest(p1TcpCnct *client.TcpConnection) {
  fmt.Println("NetService.MyOnRequest")
}

func MyOnClose(p1TcpCnct *client.TcpConnection) {
  fmt.Println("NetService.MyOnClose")
}
