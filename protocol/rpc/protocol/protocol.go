package protocol

import "net"

// ProtocolI9 协议接口
type ProtocolI9 interface {
	// EncodeReq 请求编码成报文
	EncodeReq(*Request) ([]byte, error)
	// DecodeReq 报文解码成请求
	DecodeReq([]byte) (*Request, error)

	// EncodeResp 响应编码成报文
	EncodeResp(*Response) ([]byte, error)
	// DecodeResp 报文解码成响应
	DecodeResp([]byte) (*Response, error)

	// ReadReqMsg 从连接中读取一条请求报文
	ReadReqMsg(net.Conn) ([]byte, error)
	// ReadRespMsg 从连接中读取一条响应报文
	ReadRespMsg(net.Conn) ([]byte, error)
}
