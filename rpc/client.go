package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"demo-golang/rpc/serialize"
	"net"
)

// ClientI9 可以发起 RPC 调用的客户端
type ClientI9 interface {
	// GetProtocol 获取RPC协议实现
	GetProtocol() protocol.ProtocolI9
	// GetSerialize 获取请求参数的序列化方式
	GetSerialize() serialize.SerializeI9
	// SendRPC 发起 RPC 调用
	SendRPC(context.Context, *protocol.Request) (*protocol.Response, error)
}

type Client struct {
	protocol  protocol.ProtocolI9
	serialize serialize.SerializeI9
}

func NewClient() *Client {
	c := &Client{}
	c.protocol = protocol.NewStream()
	c.serialize = serialize.F8NewS6Json()
	return c
}

func (t *Client) GetSerialize() serialize.SerializeI9 {
	return t.serialize
}

func (t *Client) GetProtocol() protocol.ProtocolI9 {
	return t.protocol
}

func (t *Client) SetProtocol(p protocol.ProtocolI9) {
	t.protocol = p
}

func (t *Client) SendRPC(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	conn, _ := net.Dial("tcp", "127.0.0.1:9602")

	p := t.GetProtocol()
	reqEn, err := p.EncodeReq(req)
	if err != nil {
		return nil, err
	}

	_, _ = conn.Write(reqEn)
	respEn, err := p.ReadRespMsg(conn)
	if err != nil {
		return nil, err
	}

	resp, err := p.DecodeResp(respEn)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
