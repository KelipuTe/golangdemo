package main

import (
	"demo-golang/rpc"
	"demo-golang/rpc/protocol"
)

func main() {
	p7s6server := rpc.F8NewS6RPCServer(
		rpc.F8SetS6RPCServerProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6UserService := &rpc.S6UserService{}
	p7s6server.F8RegisterService(p7s6UserService)

	_ = p7s6server.F8Start("127.0.0.1:9602")
}
