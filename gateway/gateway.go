package gateway

import (
	"demo-golang/websocket"
	"log"
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

type InnerHandler struct {
	gateway *Gateway
}

func NewInnerHandler() *InnerHandler {
	return &InnerHandler{}
}

func (t *InnerHandler) HandleMsg(req *websocket.Msg, conn *websocket.AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)
}

type OpenHandler struct {
}

func NewOpenHandler() *OpenHandler {
	return &OpenHandler{}
}

func (t *OpenHandler) HandleMsg(req *websocket.Msg, conn *websocket.AcceptConn) {
	log.Println(req.MsgLen, req.Fin, req.Opcode, req.Payload)
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

	inner.OnConn(func(conn *websocket.AcceptConn) {
		addr := conn.GetRemoteAddr()
		ih.gateway.innerPool[addr] = conn
	})

	return g
}

func (t *Gateway) Start() {
	go t.innerServer.Start()
	go t.openServer.Start()
}
