package gateway

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/protocol/http"
  "demo_golang/tcp_service_v2/internal/service"
)

func (p1this *Gateway) DispatchOpenRequest(p1connection *service.TCPConnection) {
  msg := p1connection.GetProtocol().(*http.HTTP)

  msgId := p1connection.GetNetConnRemoteAddr()
  p1this.mapConnOpen[msgId] = p1connection

  p1apipkg := &api.APIPackage{}
  p1apipkg.Id = msgId
  p1apipkg.Type = api.TypeRequest
  p1apipkg.Action = msg.Uri
  p1apipkg.Data = "{\"id\":1}"

  t1p1conn := p1this.mapConnPool[msg.Uri][0]
  p1this.SendInnerResponse(t1p1conn, p1apipkg)
}
