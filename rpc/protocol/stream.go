package protocol

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// Stream stream 格式的 RPC 协议
// 报文 = 报文长度(8字节) + json编码的报文内容
type Stream struct {
	msgLen int //报文长度
}

func NewStream() *Stream {
	return &Stream{
		msgLen: 8,
	}
}

func (t Stream) EncodeReq(req *Request) ([]byte, error) {
	s5MsgBody, err := json.Marshal(req)
	if nil != err {
		return nil, err
	}
	log.Printf("EncodeReq:%s,%+v", string(s5MsgBody), req)
	s5ReqMsg := make([]byte, t.msgLen+len(s5MsgBody))
	binary.BigEndian.PutUint64(s5ReqMsg[:t.msgLen], uint64(len(s5MsgBody)))
	copy(s5ReqMsg[8:], s5MsgBody)
	return s5ReqMsg, nil
}

func (t Stream) DecodeReq(s5ReqMsg []byte) (*Request, error) {
	p7s6req := &Request{}
	err := json.Unmarshal(s5ReqMsg, p7s6req)
	if nil != err {
		return nil, err
	}
	log.Printf("DecodeReq:%s,%+v", string(s5ReqMsg), p7s6req)
	return p7s6req, nil
}

func (t Stream) EncodeResp(p7s6resp *Response) ([]byte, error) {
	s5MsgBody, err := json.Marshal(p7s6resp)
	if nil != err {
		return nil, err
	}
	log.Printf("EncodeReq:%s,%+v", string(s5MsgBody), p7s6resp)
	s5RespMsg := make([]byte, t.msgLen+len(s5MsgBody))
	binary.BigEndian.PutUint64(s5RespMsg[:t.msgLen], uint64(len(s5MsgBody)))
	copy(s5RespMsg[8:], s5MsgBody)
	return s5RespMsg, nil
}

func (t Stream) DecodeResp(s5RespMsg []byte) (*Response, error) {
	p7s6resp := &Response{}
	err := json.Unmarshal(s5RespMsg, p7s6resp)
	if nil != err {
		return nil, err
	}
	log.Printf("DecodeResp:%s,%+v", string(s5RespMsg), p7s6resp)
	return p7s6resp, nil
}

func (t Stream) ReadReqMsg(i9conn net.Conn) (s5ReqMsg []byte, err error) {
	defer func() {
		if err2 := recover(); nil != err2 {
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()

	s5ReqMsgLen := make([]byte, t.msgLen)
	readByteNum, err := i9conn.Read(s5ReqMsgLen)
	if nil != err {
		return nil, err
	}
	if t.msgLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqMsgLen := binary.BigEndian.Uint64(s5ReqMsgLen)
	s5ReqMsg = make([]byte, reqMsgLen)
	_, err = io.ReadFull(i9conn, s5ReqMsg)
	return s5ReqMsg, err
}

func (t Stream) ReadRespMsg(i9conn net.Conn) (s5RespMsg []byte, err error) {
	defer func() {
		if err2 := recover(); nil != err2 {
			// 因为这个地方要返回异常，所以返回值要用命名的
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()

	s5RespMsgLen := make([]byte, t.msgLen)
	readByteNum, err := i9conn.Read(s5RespMsgLen)
	if nil != err {
		return nil, err
	}
	if t.msgLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	respMsgLen := binary.BigEndian.Uint64(s5RespMsgLen)
	s5RespMsg = make([]byte, respMsgLen)
	_, err = io.ReadFull(i9conn, s5RespMsg)
	return s5RespMsg, err
}
