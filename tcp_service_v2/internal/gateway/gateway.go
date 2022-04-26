package gateway

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/protocol/stream"
  "demo_golang/tcp_service_v2/internal/service"
  "encoding/json"
  "fmt"
)

var P1gateway *Gateway

type Gateway struct {
  p1service *service.TCPService

  // mapConnectionPool 不同服务的连接池
  mapConnectionPool map[string][]*service.TCPConnection
}

func init() {
  t1map := make(map[string][]*service.TCPConnection)
  P1gateway = &Gateway{mapConnectionPool: t1map}
}

func (p1this *Gateway) GetConn(name string) *service.TCPConnection {
  return p1this.mapConnectionPool[name][0]
}

func (p1this *Gateway) SetService(p1service *service.TCPService) {
  p1this.p1service = p1service
}

// RegisteServiceProvider 注册服务提供者
func (p1this *Gateway) RegisteServiceProvider(p1connection *service.TCPConnection) {
  p1apiPkg := &api.APIPackage{}
  msg := p1connection.GetProtocol().(*stream.Stream).GetDecodeMsg()
  json.Unmarshal([]byte(msg), p1apiPkg)

  switch p1apiPkg.Type {
  case api.TypeRequest:
    switch p1apiPkg.Action {
    case "registe":
      p1req := &api.ReqInRegiste{}
      json.Unmarshal([]byte(p1apiPkg.Data), p1req)
      p1this.Registe(p1connection, p1req)
      p1apiPkg.Type = api.TypeResponse
      p1apiPkg.Data = "服务注册成功"
      p1this.Response(p1connection, p1apiPkg)
    }
  case api.TypeResponse:
  }
}

func (p1this *Gateway) Response(p1connection *service.TCPConnection, p1apiData *api.APIPackage) {
  msg, _ := json.Marshal(p1apiData)
  p1connection.SendMsg(msg)
}

func (p1this *Gateway) Registe(p1connection *service.TCPConnection, p1req *api.ReqInRegiste) {
  sli1bizService, ok := p1this.mapConnectionPool[p1req.Name]
  if ok {
    sli1bizService = append(sli1bizService, p1connection)
  } else {
    sli1bizService = []*service.TCPConnection{p1connection}
  }
  p1this.mapConnectionPool[p1req.Name] = sli1bizService
  fmt.Println(p1req)
}
