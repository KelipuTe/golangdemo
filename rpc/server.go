package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"demo-golang/rpc/serialize"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
)

// Server 可以处理 RPC 调用的服务端
type Server struct {
	m3i9Serialize map[uint8]serialize.SerializeI9
	i9Protocol    protocol.ProtocolI9
	// 本地服务
	m3service map[string]*ServiceReflect

	i9listener net.Listener
}

// Option 设计模式
type F8S6RPCServerOption func(*Server)

// #### func ####

func F8NewS6RPCServer(s5Option ...F8S6RPCServerOption) *Server {
	p7s6server := &Server{
		m3service: make(map[string]*ServiceReflect, 4),
	}
	for _, t4value := range s5Option {
		t4value(p7s6server)
	}
	if nil == p7s6server.m3i9Serialize {
		p7s6server.m3i9Serialize = make(map[uint8]serialize.SerializeI9, 2)
		p7s6JsonSerialize := serialize.F8NewS6Json()
		p7s6server.m3i9Serialize[p7s6JsonSerialize.F8GetCode()] = p7s6JsonSerialize
	}
	if nil == p7s6server.i9Protocol {
		p7s6server.i9Protocol = protocol.NewStream()
	}
	return p7s6server
}

func F8SetS6RPCServerSerialize(i9Serializer serialize.SerializeI9) F8S6RPCServerOption {
	return func(p7this *Server) {
		if nil == p7this.m3i9Serialize {
			p7this.m3i9Serialize = make(map[uint8]serialize.SerializeI9, 2)
			p7this.m3i9Serialize[i9Serializer.F8GetCode()] = i9Serializer
		}
	}
}

func F8SetS6RPCServerProtocol(i9Protocol protocol.ProtocolI9) F8S6RPCServerOption {
	return func(p7this *Server) {
		p7this.i9Protocol = i9Protocol
	}
}

// #### struct func ####

// 注册本地服务
func (p7this *Server) F8RegisterService(i9RPCService ServiceI9) {
	// 这里用本地服务对应的 RPC 服务的服务名作为 key
	// 这样就可以通过 RPC 客户端发过来的 RPC 调用里的服务名，找到对应的本地服务
	p7this.m3service[i9RPCService.GetServiceName()] = &ServiceReflect{
		service: i9RPCService,
		reflect: reflect.ValueOf(i9RPCService),
	}
}

// 处理 rpc
func (p7this *Server) F8HandleRPC(i9ctx context.Context, p7s6req *protocol.Request) ([]byte, error) {
	p7s6service, ok := p7this.m3service[p7s6req.ServiceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] not found", p7s6req.ServiceName)
	}
	i9Serialize, ok := p7this.m3i9Serialize[p7s6req.SerializeCode]
	if !ok {
		return nil, fmt.Errorf("serialize code [%d] not found", p7s6req.SerializeCode)
	}
	functionOutputDataEncode, err := p7s6service.handleRPC(i9ctx, p7s6req.FuncName, p7s6req.FuncInput, i9Serialize)
	if nil != err {
		return nil, err
	}
	return functionOutputDataEncode, nil
}

func (p7this *Server) f8HandleTCP(i9conn net.Conn) {
	for {
		s5ReqMsg, err := p7this.i9Protocol.ReadReqMsg(i9conn)
		if err != nil {
			// 一旦从 TCP 读取数据发生异常，这个链接最好就是断掉，字节流的异常处理太麻烦了
			log.Printf("f8HandleTCP ReadReqMsg with: %s", err)
			_ = i9conn.Close()
			return
		}
		p7s6req, err := p7this.i9Protocol.DecodeReq(s5ReqMsg)
		p7s6resp := &protocol.Response{
			Error:         errors.New("OK"),
			SerializeCode: p7s6req.SerializeCode,
		}
		if err != nil {
			log.Printf("f8HandleTCP DecodeReq with: %s", err)
			p7s6resp.Error = err
		}
		i9ctx := context.Background()
		flowId, ok := p7s6req.MetaData["flowId"]
		if ok {
			i9ctx = context.WithValue(i9ctx, "flowId", flowId)
		}
		functionOutputDataEncode, err := p7this.F8HandleRPC(i9ctx, p7s6req)
		if err != nil {
			log.Printf("f8HandleTCP F8HandleRPC with: %s", err)
			p7s6resp.Error = err
		} else {
			p7s6resp.FuncOutput = functionOutputDataEncode
		}
		s5RespMsg, err := p7this.i9Protocol.EncodeResp(p7s6resp)
		if err != nil {
			log.Printf("f8HandleTCP EncodeResp with: %s", err)
		}
		_, err = i9conn.Write(s5RespMsg)
		if err != nil {
			log.Printf("f8HandleTCP Write with: %s", err)
		}
	}
}

func (p7this *Server) F8Start(address string) error {
	i9listener, err := net.Listen("tcp", address)
	if nil != err {
		return err
	}
	log.Printf("rpc server start at [%s]", address)
	p7this.i9listener = i9listener
	for {
		i9conn, err2 := i9listener.Accept()
		if nil != err2 {
			log.Printf("F8Start Accept with : %s", err2)
		}
		go p7this.f8HandleTCP(i9conn)
	}
}
