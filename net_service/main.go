package main

import (
  "demo_golang/net_service/config"
  "demo_golang/net_service/service"
  "fmt"
)

func main() {
  netSvc := service.NetService{
    ServiceRunning: service.SERVICE_RUNNING_ON,
    ProtocolName:   config.STR_HTTP,
    Address:        config.ADDRESS,
    Port:           config.PORT,

    MapTcpCnctPool: make(map[string]*service.TcpConnection),
    NowTcpCnctNum:  0,
    MaxTcpCnctNum:  config.TCP_CONNECTION_MAX_NUM,

    OnStart:   MyOnStart,
    OnError:   MyOnError,
    OnConnect: MyOnConnect,
    OnRequest: MyOnRequest,
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

func MyOnRequest(p1TcpCnct *service.TcpConnection) {
  fmt.Println("NetService.MyOnRequest")
}

func MyOnClose(p1TcpCnct *service.TcpConnection) {
  fmt.Println("NetService.MyOnClose")
}
