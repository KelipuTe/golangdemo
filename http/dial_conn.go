package http

import (
	"io"
	"log"
	"net"
)

// DialConn 客户端封装的tcp连接
type DialConn struct {
	client        *Client
	conn          net.Conn //tcp连接本体
	readBuffer    []byte   //接收缓冲区
	readBufferLen int      //接收缓冲区长度
}

func NewDialConn(c *Client, n net.Conn) *DialConn {
	return &DialConn{
		client:        c,
		conn:          n,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
	}
}

// SendReq 发送请求
func (t *DialConn) SendReq(req *Request) {
	writeBuffer, err := req.encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
}

// waitResp 等待响应
func (t *DialConn) waitResp(resp *Response) {
	num, err := t.conn.Read(t.readBuffer[t.readBufferLen:])

	if err != nil {
		if err == io.EOF {
			t.close()
			return
		}
		return
	}

	t.readBufferLen += num

	for t.readBufferLen > 0 {
		copyBuffer := t.readBuffer[0:t.readBufferLen]

		err := resp.decode(copyBuffer, t.readBufferLen)
		if err != nil {
			t.close()
			return
		}

		t.readBuffer = t.readBuffer[resp.MsgLen:]
		t.readBufferLen -= resp.MsgLen
	}
}

// close 关闭连接
func (t *DialConn) close() {
	log.Println("conn close")
	_ = t.conn.Close()
}
