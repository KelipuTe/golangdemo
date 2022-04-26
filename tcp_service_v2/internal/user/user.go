package user

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/client"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/protocol/stream"
  "encoding/json"
)

var P1UserService *UserService

func init() {
  P1UserService = &UserService{}
}

// HandlerFunc 路由对应的处理方法
type HandlerFunc func(*api.APIPackage)

type UserService struct {
  p1connection *client.TCPConnection
  sli1Route    []string
  mapRoute     map[string]HandlerFunc
}

func (p1this *UserService) SetClient(p1connection *client.TCPConnection) {
  p1this.p1connection = p1connection
}

// RegisteServiceProvider 向 gateway 发送服务提供者的注册信息
func (p1this *UserService) RegisteServiceProvider() {

  p1this.mapRoute = map[string]HandlerFunc{
    "/api/user_info": p1this.GetUserInfo,
  }
  p1this.sli1Route = []string{"/api/user_info"}

  // 拼装数据
  p1apipkg := &api.APIPackage{}
  p1apipkg.Type = api.TypeRequest
  p1apipkg.Action = "registe_service_provider"
  t1data := &api.ReqInRegisteServiceProvider{
    Name:      "user_service",
    Sli1Route: p1this.sli1Route,
  }
  t1dataJson, _ := json.Marshal(t1data)
  p1apipkg.Data = string(t1dataJson)
  p1apipkgJson, _ := json.Marshal(p1apipkg)

  // 发送数据
  protocolName := p1this.p1connection.GetProtocolName()
  switch protocolName {
  case protocol.StreamStr:
    t1p1protocol := p1this.p1connection.GetProtocol().(*stream.Stream)
    t1p1protocol.SetDecodeMsg(string(p1apipkgJson))
    p1this.p1connection.SendMsg([]byte{})
  }
}

func (p1this *UserService) DispatchRequest(p1connection *client.TCPConnection) {
  p1apipkg := &api.APIPackage{}
  msg := p1connection.GetProtocol().(*stream.Stream).GetDecodeMsg()
  json.Unmarshal([]byte(msg), p1apipkg)

  switch p1apipkg.Type {
  case api.TypeRequest:
    t1func := p1this.mapRoute[p1apipkg.Action]
    t1func(p1apipkg)
  case api.TypeResponse:
  }
}
