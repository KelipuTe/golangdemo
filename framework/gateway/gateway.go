package gateway

import (
	"demo-golang/websocket"
)

type Gateway struct {
	name string

	//给内部注册用的服务
	innerServer *websocket.Server
	//k=服务名，v=连接
	innerPool map[string]*websocket.AcceptConn

	//给外部调接口用的服务
	openServer *websocket.Server
	//k=ip+端口，v=连接
	openPool map[string]*websocket.AcceptConn
}

func NewGateway(name string, innerPort, openPort int) *Gateway {
	ih := NewInnerHandler()
	inner := websocket.NewServer(innerPort, ih)

	oh := NewOpenHandler()
	open := websocket.NewServer(openPort, oh)

	g := &Gateway{
		name:        name,
		innerServer: inner,
		innerPool:   make(map[string]*websocket.AcceptConn),
		openServer:  open,
		openPool:    make(map[string]*websocket.AcceptConn),
	}

	ih.gateway = g
	oh.gateway = g

	return g
}

func (t *Gateway) Start() {
	go t.innerServer.Start()
	go t.openServer.Start()
}
