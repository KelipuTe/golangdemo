package http

import (
	"io"
	"log"
	"net"
)

const (
	readBufferMaxLen = 1048576 // 1048576 == 2^20 == 1MB。
)

// AcceptConn 服务端封装的tcp连接
type AcceptConn struct {
	server        *Server
	conn          net.Conn //tcp连接本体
	readBuffer    []byte   //接收缓冲区
	readBufferLen int      //接收缓冲区长度
	keepAlive     bool
}

func NewAcceptConn(s *Server, c net.Conn) *AcceptConn {
	return &AcceptConn{
		server:        s,
		conn:          c,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
		keepAlive:     false,
	}
}

// handleMsg 处理消息
func (t *AcceptConn) handleMsg() {
	for {
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

			req := NewRequest()
			req.Addr = t.conn.RemoteAddr().String()
			err := req.decode(copyBuffer, t.readBufferLen)
			if err != nil {
				t.close()
				return
			}

			t.readBuffer = t.readBuffer[req.MsgLen:]
			t.readBufferLen -= req.MsgLen
			t.keepAlive = req.isKeepAlive()

			resp := NewResponse()
			t.server.handler.HandleHTTP(req, resp)
			t.sendResp(resp)
		}

		// 如果不是长连接，则关闭连接
		if !t.keepAlive {
			t.close()
			return
		}
	}
}

// sendResp 发送响应
func (t *AcceptConn) sendResp(resp *Response) {
	writeBuffer, err := resp.encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
}

// close 关闭连接
func (t *AcceptConn) close() {
	log.Println("conn close")
	_ = t.conn.Close()
	t.server.connClose(t) //通知server
}
