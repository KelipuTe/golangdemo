package user

import (
  "demo_golang/tcp_service_v2/internal/api"
  "demo_golang/tcp_service_v2/internal/protocol/stream"
  "encoding/json"
  "fmt"
)

func (p1this *UserService) GetUserInfo(p1apipkg *api.APIPackage) {
  p1req := &api.ReqInUserInfo{}
  json.Unmarshal([]byte(p1apipkg.Data), p1req)
  fmt.Printf("p1req: %+v\r\n", p1req)
  p1req.Name = "xxx"

  p1reqJson, _ := json.Marshal(p1req)
  p1apipkg.Type = api.TypeResponse
  p1apipkg.Data = string(p1reqJson)
  p1apipkgJson, _ := json.Marshal(p1apipkg)

  t1p1protocol := p1this.p1connection.GetProtocol().(*stream.Stream)
  t1p1protocol.SetDecodeMsg(string(p1apipkgJson))
  p1this.p1connection.SendMsg([]byte{})
}
