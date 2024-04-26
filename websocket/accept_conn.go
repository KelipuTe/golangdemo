package websocket

import (
	"crypto/sha1"
	"demo-golang/http"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net"
	"strings"
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
	hasHandshake  bool     //是否握手
}

func NewAcceptConn(s *Server, c net.Conn) *AcceptConn {
	return &AcceptConn{
		server:        s,
		conn:          c,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
		hasHandshake:  false,
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
			if t.hasHandshake {
				//握好手了，用websocket解析
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

				t.server.handler.HandleMsg(req, t)

			} else {
				//没有握手，用http解析
				req := http.NewRequest()
				resp := http.NewResponse()
				err := t.checkHandshakeReq(req, resp)
				t.sendHandshakeResp(resp)
				if err != nil {
					t.close()
					return
				}
				t.hasHandshake = true
			}
		}
	}
}

func (t *AcceptConn) checkHandshakeReq(req *http.Request, resp *http.Response) error {
	copyBuffer := t.readBuffer[0:t.readBufferLen]

	req.Addr = t.conn.RemoteAddr().String()
	err := req.Decode(copyBuffer, t.readBufferLen)
	if err != nil {
		return err
	}

	t.readBuffer = t.readBuffer[req.MsgLen:]
	t.readBufferLen -= req.MsgLen

	//检查握手信息
	if v, ok := req.Header["Connection"]; !ok ||
		strings.Index(v, "Upgrade") < 0 {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	if v, ok := req.Header["Upgrade"]; !ok ||
		strings.Index(v, "websocket") < 0 {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	if _, ok := req.Header["Sec-WebSocket-Key"]; !ok {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	//计算 Sec-WebSocket-Accept
	secKey := req.Header["Sec-WebSocket-Key"]
	//将 Sec-WebSocket-Key 跟 "258EAFA5-E914-47DA-95CA-C5AB0DC85B11" 拼接
	secAccept := secKey + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	saSHA1 := sha1.Sum([]byte(secAccept))                    //SHA1 计算摘要
	saBase64 := base64.StdEncoding.EncodeToString(saSHA1[:]) //转成 base64

	//测试使用的是 JavaScript 的 WebSocket 工具
	//在上述条件下，这几个字段都要有，少一个都跑不通
	resp.Status = http.StatusSwitchingProtocols
	resp.Header["Connection"] = "Upgrade"
	resp.Header["Upgrade"] = "websocket"
	resp.Header["Sec-WebSocket-Version"] = "13"
	resp.Header["Sec-WebSocket-Accept"] = saBase64

	return nil
}

func (t *AcceptConn) sendHandshakeResp(resp *http.Response) {
	writeBuffer, err := resp.Encode()
	if err != nil {
		return
	}
	_, _ = t.conn.Write(writeBuffer)
	return
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
