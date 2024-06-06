package gateway

import (
	"demo-golang/websocket"
)

type Gateway struct {
	name string

	//给内部服务注册用
	innerServer *websocket.Server
	innerPool   map[string]*websocket.AcceptConn

	//给外部调接口用
	openServer *websocket.Server
	openPool   map[string]*websocket.AcceptConn
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
