package router

import (
	"fmt"
	"net/http"
)

type HTTPContext struct {
	I9writer    http.ResponseWriter
	P7request   *http.Request
	M3pathParam map[string]string

	p7routingNode *routingNode
}

func (this HTTPContext) GetRoutingInfo() string {
	return fmt.Sprintf("nodeType:%d\r\nrouting path:%s\r\n", this.p7routingNode.nodeType, this.p7routingNode.path)
}
