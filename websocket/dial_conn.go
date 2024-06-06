package websocket

import (
	"crypto/md5"
	"crypto/sha1"
	"demo-golang/http"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// DialConn 客户端封装的tcp连接
type DialConn struct {
	client *Client

	conn          net.Conn //tcp连接本体
	readBuffer    []byte   //接收缓冲区
	readBufferLen int      //接收缓冲区长度

	hasHandshake bool   //是否握手
	secKey       string //握手用的key

	isRunning bool //是否运行
}

func NewDialConn(c *Client, n net.Conn) *DialConn {
	return &DialConn{
		client:        c,
		conn:          n,
		readBuffer:    make([]byte, readBufferMaxLen),
		readBufferLen: 0,
		isRunning:     true,
	}
}

func (t *DialConn) handshake() error {
	hsReq := http.NewRequest()
	t.makeHandshakeReq(hsReq)
	err := t.sendHandshakeReq(hsReq)
	if err != nil {
		return err
	}

	hsResp := http.NewResponse()
	t.waitHandshakeResp(hsResp)
	err = t.checkHandshakeResp(hsResp)
	if err != nil {
		return err
	}

	t.hasHandshake = true
	go t.sendPing()

	return nil
}

// makeHandshakeReq 构造握手请求
func (t *DialConn) makeHandshakeReq(req *http.Request) {
	secKey := strconv.Itoa(int(time.Now().Unix()))
	secKeyMD5 := md5.Sum([]byte(secKey))
	t.secKey = base64.StdEncoding.EncodeToString(secKeyMD5[:])

	req.Method = "GET"
	req.Uri = "/chat"
	req.Header["Connection"] = "Upgrade"
	req.Header["Upgrade"] = "websocket"
	req.Header["Sec-WebSocket-Version"] = "13"
	req.Header["Sec-WebSocket-Key"] = t.secKey
}

// sendHandshakeReq 发送握手请求
func (t *DialConn) sendHandshakeReq(req *http.Request) error {
	writeBuffer, err := req.Encode()
	if err != nil {
		return err
	}
	_, err = t.conn.Write(writeBuffer)
	if err != nil {
		return err
	}
	return nil
}

// waitHandshakeResp 等待握手响应
func (t *DialConn) waitHandshakeResp(resp *http.Response) {
	num, err := t.conn.Read(t.readBuffer[t.readBufferLen:])
	if err != nil {
		if err == io.EOF {
			log.Println("conn read error:", err)
			t.close()
			return
		}
		log.Println("conn read error:", err)
		t.close()
		return
	}

	t.readBufferLen += num

	if t.readBufferLen <= 0 {
		return
	}

	copyBuffer := t.readBuffer[0:t.readBufferLen]

	err = resp.Decode(copyBuffer, t.readBufferLen)
	if err != nil {
		log.Println("handshake error:", err)
		t.close()
		return
	}

	t.readBuffer = t.readBuffer[resp.MsgLen:]
	t.readBufferLen -= resp.MsgLen

	return
}

// checkHandshakeResp 检查握手响应
func (t *DialConn) checkHandshakeResp(resp *http.Response) error {
	//检查握手信息
	if v, ok := resp.Header["Connection"]; !ok ||
		strings.Index(v, "Upgrade") < 0 {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	if v, ok := resp.Header["Upgrade"]; !ok ||
		strings.Index(v, "websocket") < 0 {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	if _, ok := resp.Header["Sec-WebSocket-Accept"]; !ok {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	// 校验 Sec-WebSocket-Accept
	secAccept := t.secKey + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	saSHA1 := sha1.Sum([]byte(secAccept))
	saBase64 := base64.StdEncoding.EncodeToString(saSHA1[:])

	if saBase64 != resp.Header["Sec-WebSocket-Accept"] {
		resp.Status = 400
		resp.Body = "Handshake failed"
		return errors.New("handshake failed")
	}

	return nil
}

// handleMsg 处理消息
func (t *DialConn) handleMsg() {
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

			copyBuffer := t.readBuffer[0:t.readBufferLen]

			req := NewMaskTextMsg()
			req.Addr = t.conn.RemoteAddr().String()
			err := req.Decode(copyBuffer, t.readBufferLen)
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
						t.client.handler.HandleMsg(req, t)
					} else {
						reqMerge.Payload = reqMerge.Payload + req.Payload
						t.client.handler.HandleMsg(req, t)
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
				err := t.SendMsg(resp)
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
		}
	}
}

func (t *DialConn) sendPing() {
	for t.isRunning {
		req := NewPingMsg()
		err := t.SendMsg(req)
		if err != nil {
			log.Println("send ping error:", err)
			t.close()
			return
		}
		time.Sleep(10 * time.Second)
	}
}

// SendMsg 发送消息
func (t *DialConn) SendMsg(req *Msg) error {
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
func (t *DialConn) close() {
	log.Println("conn close")
	_ = t.conn.Close()
	t.client.connClose() //通知client
}
