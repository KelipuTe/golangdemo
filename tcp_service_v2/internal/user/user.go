package user

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/client"
  "demo_golang/tcp_service_v2/internal/protocol"
  "demo_golang/tcp_service_v2/internal/protocol/stream"
  "encoding/json"
  "fmt"
)

var P1UserService *UserService

func init() {
  t1mapRoute := map[string]HandlerFunc{
    "/api/user_info": GetUserInfo,
  }
  t1sliRoute := []string{"/api/user_info"}

  P1UserService = &UserService{
    mapRoute:  t1mapRoute,
    sli1Route: t1sliRoute,
  }
}

// HandlerFunc 路由对应的处理方法
type HandlerFunc func()

type UserService struct {
  p1connection *client.TCPConnection
  sli1Route    []string
  mapRoute     map[string]HandlerFunc
}

func GetUserInfo() {
  fmt.Println("GetUserInfo")
}

func (p1this *UserService) SetClient(p1connection *client.TCPConnection) {
  p1this.p1connection = p1connection
}

func (p1this *UserService) RegisteServiceProvider() {
  apiPackage := &api.APIPackage{}
  apiPackage.Type = api.TypeRequest
  apiPackage.Action = "registe"
  t1mapData := map[string]interface{}{
    "name":  "user_service",
    "route": p1this.sli1Route,
  }
  msg, _ := json.Marshal(t1mapData)
  apiPackage.Data = string(msg)
  t1apiPackage, _ := json.Marshal(apiPackage)

  // 发送服务注册数据
  protocolName := p1this.p1connection.GetProtocolName()
  switch protocolName {
  case protocol.StreamStr:
    t1p1protocol := p1this.p1connection.GetProtocol().(*stream.Stream)
    t1p1protocol.SetDecodeMsg(string(t1apiPackage))
    p1this.p1connection.SendMsg([]byte{})
  }
}
