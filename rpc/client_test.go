package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"fmt"
	"testing"
)

func Test_Client(p7s6t *testing.T) {
	p7s6client := F8NewS6RPCClient(
		F8SetS6RPCClientProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6RPCService := &S6UserRPCService{}
	F8CoverWithRPC(p7s6client, p7s6RPCService)

	i9ctx := context.Background()
	i9ctx = context.WithValue(i9ctx, "flowId", "flowId12345678")
	resp, err := p7s6RPCService.F8GetUserById(i9ctx, &S6F8GetUserByIdRequest{UserId: 33})
	fmt.Println(resp, err)
}
