package gateway

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/protocol/http"
  "demo_golang/tcp_service_v2/internal/protocol/stream"
  "demo_golang/tcp_service_v2/internal/service"
  "encoding/json"
)

// DispatchInnerRequest
func (p1this *Gateway) DispatchInnerRequest(p1connection *service.TCPConnection) {
  p1apipkg := &api.APIPackage{}
  msg := p1connection.GetProtocol().(*stream.Stream).GetDecodeMsg()
  json.Unmarshal([]byte(msg), p1apipkg)

  switch p1apipkg.Type {
  case api.TypeRequest:
    switch p1apipkg.Action {
    case "registe_service_provider":
      // 接收服务提供者的注册信息
      p1this.RegisteServiceProvider(p1connection, p1apipkg)

      p1apipkg.Type = api.TypeResponse
      p1apipkg.Data = "registe_service_provider success."
      p1this.SendInnerResponse(p1connection, p1apipkg)
    }
  case api.TypeResponse:
    resp := http.NewResponse()
    resp.SetStatusCode(http.StatusOk)
    respStr := resp.MakeResponse(p1apipkg.Data)

    t1p1connection := p1this.mapConnOpen[p1apipkg.Id]
    t1p1connection.SendMsg([]byte(respStr))
    t1p1connection.CloseConnection()
  }
}

func (p1this *Gateway) SendInnerResponse(p1connection *service.TCPConnection, p1apipkg *api.APIPackage) {
  p1apipkgJson, _ := json.Marshal(p1apipkg)

  t1p1protocol := p1connection.GetProtocol().(*stream.Stream)
  t1p1protocol.SetDecodeMsg(string(p1apipkgJson))
  p1connection.SendMsg([]byte{})
}

func (p1this *Gateway) RegisteServiceProvider(p1connection *service.TCPConnection, p1apipkg *api.APIPackage) {
  p1req := &api.ReqInRegisteServiceProvider{}
  json.Unmarshal([]byte(p1apipkg.Data), p1req)

  // 注册 api
  for _, api := range p1req.Sli1Route {
    sli1bizService, ok := p1this.mapConnPool[api]
    if ok {
      sli1bizService = append(sli1bizService, p1connection)
    } else {
      sli1bizService = []*service.TCPConnection{p1connection}
    }
    p1this.mapConnPool[api] = sli1bizService
  }
}
