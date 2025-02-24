package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"fmt"
	"testing"
)

func Test_Client(p7s6t *testing.T) {
	c := NewClient()
	c.SetProtocol(protocol.F8NewS6CustomRPC())

	svc := &CaseUserRPCService{}
	CoverWithRPC(c, svc)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "traceID", "1234567890")
	resp, err := svc.GetUserByID(ctx, &GetUserByIDReq{UserId: 33})
	fmt.Println(resp, err)
}
