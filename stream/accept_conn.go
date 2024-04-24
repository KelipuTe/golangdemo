package stream

import (
	"io"
	"log"
	"net"
	"time"
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
}

func NewAcceptConn(s *Server, c net.Conn) *AcceptConn {
	return &AcceptConn{
		server:        s,
		conn:          c,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
	}
}

// handleMsg 处理消息
func (t *AcceptConn) handleMsg() {
	for {
		err := t.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.close()
			return
		}

		num, err := t.conn.Read(t.readBuffer[t.readBufferLen:])
		if err != nil {
			if err == io.EOF {
				t.close()
				return
			}
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("conn read timeout")
				continue //超时，可以回去，继续等待
			}
			log.Println("conn read error:", err)
			t.close()
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

			resp := NewResponse()
			t.server.handler.HandleMsg(req, resp)
			t.sendResp(resp)
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
