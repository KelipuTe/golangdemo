package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	VersionX1Y1 = "HTTP/1.1"

	MethodGet  = "GET"
	MethodPost = "POST"
)

var (
	ErrParseFailed = errors.New("解析失败")
)

type Request struct {
	HeaderLen  int
	ContentLen int
	MsgLen     int
	Msg        string

	Addr    string            //请求IP和端口
	Method  string            //请求方法
	Uri     string            //请求路由
	Query   map[string]string //解析后的查询参数
	Version string            //版本
	Header  map[string]string //解析后的请求头
	Body    string            //请求体
}

func NewRequest() *Request {
	return &Request{
		Query:   make(map[string]string),
		Version: VersionX1Y1,
		Header:  make(map[string]string),
	}
}

// 结构体->http请求报文
func (t *Request) encode() ([]byte, error) {
	msg := fmt.Sprintf("%s %s", t.Method, t.Uri)
	if len(t.Query) > 0 {
		msg += "?"
		i := 0
		for k, v := range t.Query {
			i++
			msg += fmt.Sprintf("%s=%s", k, v)
			if i < len(t.Query) {
				msg += "&"
			}
		}
	}
	msg += fmt.Sprintf(" %s\r\n", t.Version)

	for k, v := range t.Header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += fmt.Sprintf("Content-Length: %v\r\n\r\n%s", len(t.Body), t.Body)

	return []byte(msg), nil
}

// http请求报文->结构体
func (t *Request) decode(buffer []byte, bufferLen int) error {
	bufferStr := string(buffer)

	//找到 \r\n\r\n 的位置，用这个位置可以分隔请求头和请求体
	rnrnIndex := strings.Index(bufferStr, "\r\n\r\n")
	if rnrnIndex <= 0 {
		return ErrParseFailed
	}
	//请求头长度等于 \r\n\r\n 的位置下标加上 \r\n\r\n 的长度
	t.HeaderLen = rnrnIndex + 4

	clIndex := strings.Index(bufferStr, "Content-Length: ")
	if clIndex > 0 {
		clStr := bufferStr[clIndex+16:]         //"Content-Length: " 16 个字节
		rnIndex := strings.Index(clStr, "\r\n") //找到这一行请求头的 \r\n 的位置
		clStr = clStr[0:rnIndex]                //截取 Content-Length 的字符串值
		cl, err := strconv.Atoi(clStr)          //把字符串值转换成整数值
		if err != nil {
			return ErrParseFailed
		}
		t.ContentLen = cl
	}

	t.MsgLen = t.HeaderLen + t.ContentLen
	if t.MsgLen > bufferLen {
		// 计算出来的报文长度大于接收缓冲区中数据的长度
		return ErrParseFailed
	}
	t.Msg = bufferStr[0:t.MsgLen]

	err := t.parseHeader()
	if err != nil {
		return err
	}

	t.Body = t.Msg[t.HeaderLen:]

	return nil
}

func (t *Request) parseHeader() error {
	header := t.Msg[:t.HeaderLen]
	headerSplit := strings.Split(header, "\r\n")

	//请求行
	url := headerSplit[0]
	urlSplit := strings.Split(url, " ")
	t.Method = urlSplit[0]
	t.Uri = urlSplit[1]
	err := t.parseQuery()
	if err != nil {
		return err
	}
	t.Version = urlSplit[2]

	//请求头
	for _, v := range headerSplit[1:] {
		vSplit := strings.Split(v, ": ") //用 ": " 切成键值
		if len(vSplit) == 2 {
			t.Header[strings.ToLower(vSplit[0])] = vSplit[1] //键名全部转成小写
		}
	}

	return nil
}

func (t *Request) parseQuery() error {
	index := strings.Index(t.Uri, "?")
	if index < 0 {
		// ? 不存在
		return nil
	} else if index == 0 {
		// ? 在第一个字符的位置，不合法
		return ErrParseFailed
	}
	query := t.Uri[index+1:] //查询参数
	t.Uri = t.Uri[:index]    //没有查询参数的uri
	if query != "" {
		querySplit := strings.Split(query, "&")
		for _, v := range querySplit {
			vSplit := strings.Split(v, "=")
			if len(vSplit) == 2 {
				t.Query[strings.ToLower(vSplit[0])] = vSplit[1]
			}
		}
	}

	return nil
}

func (t *Request) parseForm() (map[string]string, error) {
	ct := t.Header["content-type"]
	if ct != "application/x-www-form-urlencoded" {
		return nil, ErrParseFailed
	}
	ret := make(map[string]string)
	bodySplit := strings.Split(t.Body, "&")
	for _, v := range bodySplit {
		vSplit := strings.Split(v, "=")
		if len(vSplit) == 2 {
			ret[vSplit[0]] = vSplit[1]
		}
	}
	return ret, nil
}

func (t *Request) parseJson() (map[string]any, error) {
	ct := t.Header["content-type"]
	if ct != "application/json" {
		return nil, ErrParseFailed
	}
	ret := make(map[string]any)
	err := json.Unmarshal([]byte(t.Body), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (t *Request) KeepAliveOn() {
	t.Header["connection"] = "keep-alive"
}

func (t *Request) KeepAliveOff() {
	t.Header["connection"] = "close"
}

func (t *Request) isKeepAlive() bool {
	if conn, ok := t.Header["connection"]; ok {
		if conn == "keep-alive" {
			return true
		}
	}
	return false
}
