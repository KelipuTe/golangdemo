package gateway

import (
	"demo-golang/tcp-service/api"
	"demo-golang/tcp-service/protocol/http"
	"demo-golang/tcp-service/service"
	"fmt"
)

func (p1this *Gateway) DispatchOpenRequest(p1conn *service.TCPConnection) {
	msg := p1conn.GetProtocolHandler().(*http.Handler)

	t1p1conn := p1this.GetInnerConn(msg.Uri)
	// 如果找不到 api 对应的服务提供者，就直接报错给外部连接
	if nil == t1p1conn {
		resp := http.NewResponse()
		resp.SetStatusCode(http.StatusBadRequest)
		respStr := resp.MakeMsg("api not found.")

		p1conn.SendMsg([]byte(respStr))
		p1conn.CloseConnection()
		return
	}

	msgId := p1conn.GetNetConnRemoteAddr()
	p1this.mapOpenConn[msgId] = p1conn

	p1apipkg := &api.APIPackage{}
	p1apipkg.Id = msgId
	p1apipkg.Type = api.TypeRequest
	p1apipkg.Action = msg.Uri
	p1apipkg.Data = fmt.Sprintf("{\"id\":%s}", msg.MapQuery["id"])

	p1this.SendInnerResponse(t1p1conn, p1apipkg)
}
