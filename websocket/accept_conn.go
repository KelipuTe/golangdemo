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

// AcceptConn 服务端封装的tcp连接
type AcceptConn struct {
	server *Server

	conn          net.Conn //tcp连接本体
	readBuffer    []byte   //接收缓冲区
	readBufferLen int      //接收缓冲区长度

	hasHandshake bool //是否握手

	isRunning bool //是否运行
}

func NewAcceptConn(s *Server, c net.Conn) *AcceptConn {
	return &AcceptConn{
		server:        s,
		conn:          c,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
		hasHandshake:  false,
		isRunning:     true,
	}
}

// handleMsg 处理消息
func (t *AcceptConn) handleMsg() {
	for t.isRunning {
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

		var reqMerge *Msg = nil
		for t.readBufferLen > 0 {
			if t.hasHandshake {
				//握好手了，用websocket解析
				copyBuffer := t.readBuffer[0:t.readBufferLen]

				req := NewMaskTestMsg()
				req.Addr = t.conn.RemoteAddr().String()
				err := req.decode(copyBuffer, t.readBufferLen)
				if err != nil {
					t.close()
					return
				}

				t.readBuffer = t.readBuffer[req.MsgLen:]
				t.readBufferLen -= req.MsgLen

				switch req.Opcode {
				case opcodeText:
					//看看需不需要合并
					if req.Fin == fin1 {
						if reqMerge == nil {
							t.server.handler.HandleMsg(req, t)
						} else {
							reqMerge.Payload = reqMerge.Payload + req.Payload
							t.server.handler.HandleMsg(req, t)
							reqMerge = nil
						}
					} else {
						if reqMerge == nil {
							reqMerge = req
						} else {
							reqMerge.Payload = reqMerge.Payload + req.Payload
						}
					}
				case opcodePing:
					log.Println("get ping from", req.Addr)
					resp := NewPongMsg()
					err := t.sendMsg(resp)
					if err != nil {
						t.close()
						return
					}
					continue
				case opcodePong:
					log.Println("get pong from", req.Addr)
					continue
				default:
				}
			} else {
				//没有握手，用http解析
				if t.isHandshakeReq(req) {
					//有http升级websocket的字段，走握手逻辑
					req := http.NewRequest()
					resp := http.NewResponse()
					err := t.checkHandshakeReq(req, resp)
					if err != nil {
						t.close()
						return
					}
					err = t.sendHttpResp(resp)
					if err != nil {
						t.close()
						return
					}
					t.hasHandshake = true
					go t.sendPing()
				} else {
					//没有http升级websocket的字段，当http请求处理
					req := http.NewRequest()
					resp := http.NewResponse()
					t.server.httpHandler.HandleMsg(req, resp)
					t.sendHttpResp(resp)
				}
			}
		}
	}
}

func (t *AcceptConn) parseReq(req *http.Request) error {
	copyBuffer := t.readBuffer[0:t.readBufferLen]

	req.Addr = t.conn.RemoteAddr().String()
	err := req.Decode(copyBuffer, t.readBufferLen)
	if err != nil {
		return err
	}

	t.readBuffer = t.readBuffer[req.MsgLen:]
	t.readBufferLen -= req.MsgLen

	return nil
}

func (t *AcceptConn) isHandshakeReq(req *http.Request) bool {

}

// checkHandshakeReq 检查握手请求
func (t *AcceptConn) checkHandshakeReq(req *http.Request, resp *http.Response) error {

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

// sendHttpResp 发送http响应
func (t *AcceptConn) sendHttpResp(resp *http.Response) error {
	writeBuffer, err := resp.Encode()
	if err != nil {
		return err
	}
	_, err = t.conn.Write(writeBuffer)
	if err != nil {
		return err
	}
	return nil
}

// sendPing 发送心跳
func (t *AcceptConn) sendPing() {
	for t.isRunning {
		req := NewPingMsg()
		err := t.sendMsg(req)
		if err != nil {
			log.Println("send ping error:", err)
			t.close()
			return
		}
		time.Sleep(10 * time.Second)
	}
}

// sendMsg 发送消息
func (t *AcceptConn) sendMsg(req *Msg) error {
	writeBuffer, err := req.encode()
	if err != nil {
		return err
	}
	_, err = t.conn.Write(writeBuffer)
	if err != nil {
		return err
	}
	return nil
}

// close 关闭连接
func (t *AcceptConn) close() {
	log.Println("conn close")
	_ = t.conn.Close()
	t.isRunning = false
	t.server.connClose(t) //通知server
}
