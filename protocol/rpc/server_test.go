package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Server(p7s6t *testing.T) {
	p7s6server := F8NewS6RPCServer(
		F8SetS6RPCServerProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6UserService := &CaseUserService{}
	p7s6server.F8RegisterService(p7s6UserService)

	_ = p7s6server.F8Start("127.0.0.1:9602")
}

func TestF8HandleRPC(p7s6t *testing.T) {
	s5s6case := []struct {
		name           string
		p7s6RPCServer  *Server
		p7s6RPCRequest *protocol.Request
		wantResp       *protocol.Response
		wantErr        error
	}{
		{
			name: "user_rpc_service_client",
			p7s6RPCServer: func() *Server {
				p7s6server := F8NewS6RPCServer()
				p7s6UserService := &CaseUserService{}
				p7s6server.F8RegisterService(p7s6UserService)
				return p7s6server
			}(),
			p7s6RPCRequest: &protocol.Request{
				ServiceName: "user-rpc-service",
				FuncName:    "GetUserByID",
				FuncInput:   []byte(`{"userId":22}`),
			},
			wantResp: &protocol.Response{
				FuncOutput: []byte(`{"userId":22,"userName":"bb"}`),
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			resp := &protocol.Response{}
			functionOutputDataEncode, err := t4value.p7s6RPCServer.F8HandleRPC(context.Background(), t4value.p7s6RPCRequest)
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			resp.FuncOutput = functionOutputDataEncode
			assert.Equal(p7s6t2, t4value.wantResp, resp)
		})
	}
}
