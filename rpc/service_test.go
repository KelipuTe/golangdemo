package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"demo-golang/rpc/serialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

type s6MockI9RPCClient struct {
	p7s6t       *testing.T
	i9serialize serialize.SerializeI9
	i9protocol  protocol.ProtocolI9
	p7s6req     *protocol.Request
	p7s6resp    *protocol.Response
	err         error
}

func (p7this *s6MockI9RPCClient) GetSerialize() serialize.SerializeI9 {
	return p7this.i9serialize
}

func (p7this *s6MockI9RPCClient) GetProtocol() protocol.ProtocolI9 {
	return p7this.i9protocol
}

func (p7this *s6MockI9RPCClient) SendRPC(i9ctx context.Context, p7s6req *protocol.Request) (*protocol.Response, error) {
	assert.Equal(p7this.p7s6t, p7this.p7s6req, p7s6req)
	return p7this.p7s6resp, p7this.err
}

type s6MockI9RPCService struct {
	i9RPCService ServiceI9
	f8SendRPC    func() (any, error)
}

func TestF8CoverWithRPC(p7s6t *testing.T) {
	s5s6case := []struct {
		name            string
		p7s6MockClient  *s6MockI9RPCClient
		p7s6MockService *s6MockI9RPCService
		wantResp        *S6F8GetUserByIdResponse
		wantErr         error
	}{
		{
			name: "user_rpc_service_client",
			p7s6MockClient: &s6MockI9RPCClient{
				p7s6t:       p7s6t,
				i9serialize: serialize.F8NewS6Json(),
				p7s6req: &protocol.Request{
					ServiceName: "user-rpc-service",
					FuncName:    "F8GetUserById",
					FuncInput:   []byte(`{"userId":11}`),
				},
				p7s6resp: &protocol.Response{
					FuncOutput: []byte(`{"userId":11,"userName":"aa"}`),
				},
			},
			p7s6MockService: func() *s6MockI9RPCService {
				p7s6RPCService := &S6UserRPCService{}
				return &s6MockI9RPCService{
					i9RPCService: p7s6RPCService,
					f8SendRPC: func() (any, error) {
						return p7s6RPCService.F8GetUserById(context.Background(), &S6F8GetUserByIdRequest{UserId: 11})
					},
				}
			}(),
			wantResp: &S6F8GetUserByIdResponse{
				UserId:   11,
				UserName: "aa",
			},
		},
	}

	for _, s6case := range s5s6case {
		p7s6t.Run(s6case.name, func(p7s6t2 *testing.T) {
			CoverWithRPC(s6case.p7s6MockClient, s6case.p7s6MockService.i9RPCService)
			resp, err := s6case.p7s6MockService.f8SendRPC()
			assert.Equal(p7s6t2, s6case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, s6case.wantResp, resp)
		})
	}
}
