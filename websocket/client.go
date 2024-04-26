package websocket

import (
	"crypto/md5"
	"crypto/sha1"
	"demo-golang/http"
	"encoding/base64"
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Client 客户端
type Client struct {
	ip           string
	port         int
	conn         *DialConn
	secKey       string
	hasHandshake bool
}

func NewClient(ip string, port int) *Client {
	return &Client{
		ip:   ip,
		port: port,
	}
}

// Send 发送请求
func (t *Client) Send(req *Request) (*Response, error) {
	if t.conn == nil {
		addr := t.ip + ":" + strconv.Itoa(t.port)
		netConn, err := net.Dial("tcp4", addr) //发起连接
		if err != nil {
			return nil, err
		}

		t.connDial(netConn)

		hsReq := http.NewRequest()
		t.makeHandshakeReq(hsReq)
		t.conn.sendHandshakeReq(hsReq)

		hsResp := http.NewResponse()
		t.conn.waitHandshakeResp(hsResp)
		err = t.checkHandshakeResp(hsResp)
		if err != nil {
			return nil, err
		}
		t.hasHandshake = true
	}

	t.conn.sendReq(req)

	resp := NewResponse()
	t.conn.waitResp(resp)

	return resp, nil
}

func (t *Client) makeHandshakeReq(req *http.Request) {
	secKey := strconv.Itoa(int(time.Now().Unix()))
	secKeyMD5 := md5.Sum([]byte(secKey))
	t.secKey = base64.StdEncoding.EncodeToString(secKeyMD5[:])

	req.Method = "GET"
	req.Uri = "/chat"
	req.Header["Connection"] = "Upgrade"
	req.Header["Upgrade"] = "websocket"
	req.Header["Sec-WebSocket-Key"] = t.secKey
	req.Header["Sec-WebSocket-Version"] = "13"
}

func (t *Client) checkHandshakeResp(resp *http.Response) error {
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

// connDial 连接建立
func (t *Client) connDial(netConn net.Conn) *DialConn {
	log.Println("conn dial")

	httpConn := NewDialConn(t, netConn)

	t.conn = httpConn

	return httpConn
}

func (t *Client) connClose() {
	t.conn = nil
}

// CloseConn 关闭连接
func (t *Client) CloseConn() {
	if t.conn != nil {
		t.conn.close()
		t.conn = nil
	}
}
