package main

import (
	"context"
	"demo-golang/rpc"
	"demo-golang/rpc/protocol"
	"fmt"
)

func main() {
	p7s6client := rpc.F8NewS6RPCClient(
		rpc.F8SetS6RPCClientProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6RPCService := &rpc.S6UserRPCService{}
	rpc.F8CoverWithRPC(p7s6client, p7s6RPCService)

	i9ctx := context.Background()
	i9ctx = context.WithValue(i9ctx, "flowId", "flowId12345678")
	resp, err := p7s6RPCService.F8GetUserById(i9ctx, &rpc.S6F8GetUserByIdRequest{UserId: 33})
	fmt.Println(resp, err)
}
