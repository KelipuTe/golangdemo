package websocket

import (
	"demo-golang/http"
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

func (t *DialConn) sendHandshakeReq(req *http.Request) {
	writeBuffer, err := req.Encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
}

// sendReq 发送请求
func (t *DialConn) sendReq(req *Request) {
	writeBuffer, err := req.encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
}

func (t *DialConn) waitHandshakeResp(resp *http.Response) {
	num, err := t.conn.Read(t.readBuffer[t.readBufferLen:])
	if err != nil {
		if err == io.EOF {
			t.close()
			return
		}
		log.Println("conn read error:", err)
		t.close()
		return
	}

	t.readBufferLen += num

	if t.readBufferLen > 0 {
		copyBuffer := t.readBuffer[0:t.readBufferLen]

		err := resp.Decode(copyBuffer, t.readBufferLen)
		if err != nil {
			t.close()
			return
		}

		t.readBuffer = t.readBuffer[resp.MsgLen:]
		t.readBufferLen -= resp.MsgLen

		return
	}
}

// waitResp 等待响应
func (t *DialConn) waitResp(resp *Response) {
	num, err := t.conn.Read(t.readBuffer[t.readBufferLen:])

	if err != nil {
		if err == io.EOF {
			t.close()
			return
		}
		log.Println("conn read error:", err)
		t.close()
		return
	}

	t.readBufferLen += num

	if t.readBufferLen > 0 {
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
	t.client.connClose()
}
