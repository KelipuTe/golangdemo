package gateway

import (
  "demo_golang/tcp_service_v2/internal/service"
)

var P1gateway *Gateway

type Gateway struct {
  p1service *service.TCPService

  // mapConnPool 不同服务提供者的 TCP 连接池
  mapConnPool map[string][]*service.TCPConnection
  mapConnOpen map[string]*service.TCPConnection
}

func init() {
  P1gateway = &Gateway{
    mapConnPool: make(map[string][]*service.TCPConnection),
    mapConnOpen: make(map[string]*service.TCPConnection),
  }
}

func (p1this *Gateway) SetService(p1service *service.TCPService) {
  p1this.p1service = p1service
}

func (p1this *Gateway) GetConn(name string) *service.TCPConnection {
  return p1this.mapConnPool[name][0]
}
