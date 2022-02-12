package main

import (
  "demo_golang/net_service/service"
  "fmt"
)

func main() {
  netSvc := service.NetService{
    AppDebug:       1,
    ServiceRunning: 1,

    ProtocolName: "http",
    Address:      "127.0.0.1",
    Port:         9501,

    MapTcpCnctPool: make(map[string]*service.TcpConnection),
    NowTcpCnctNum:  0,
    MaxTcpCnctNum:  1024,

    OnStart:   MyOnStart,
    OnError:   MyOnError,
    OnConnect: MyOnConnect,
    OnClose:   MyOnClose,
  }
  netSvc.Start()
}

func MyOnStart(p1NetSvc *service.NetService) {
  fmt.Println("NetService.OnStart")
}

func MyOnError(errStr string) {
  fmt.Println("NetService.OnError")
  fmt.Println("errStr=", errStr)
}

func MyOnConnect(p1TcpCnct *service.TcpConnection) {
  fmt.Println("NetService.MyOnConnect")
}

func MyOnClose(p1TcpCnct *service.TcpConnection) {
  fmt.Println("NetService.MyOnClose")
}
