package http

import (
	"io"
	"log"
	"net"
)

const (
	readBufferMaxLen = 1048576 // 1048576 == 2^20 == 1MB。
)

type AcceptConn struct {
	server        *Server
	conn          net.Conn
	readBuffer    []byte //接收缓冲区
	readBufferLen int    //接收缓冲区长度
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

func (t *AcceptConn) handleConn() {
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

		if !t.keepAlive {
			t.close()
			return
		}
	}
}

func (t *AcceptConn) sendResp(resp *Response) {
	writeBuffer, err := resp.encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
}

func (t *AcceptConn) close() {
	log.Println("conn close")
	_ = t.conn.Close()
	t.server.closeConn(t)
}
