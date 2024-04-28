package http

import (
	"fmt"
	"strconv"
	"strings"
)

// Response 响应
type Response struct {
	HeaderLen  int    //响应头长度
	ContentLen int    //响应体长度
	MsgLen     int    //消息长度
	Msg        []byte //消息体

	Version string            //版本
	Status  int               //状态码
	Header  map[string]string //解析后的响应头
	Body    string            //响应体
}

func NewResponse() *Response {
	return &Response{
		Version: VersionX1Y1,
		Header:  make(map[string]string),
	}
}

// Encode 结构体->http响应报文
func (t *Response) Encode() ([]byte, error) {
	msg := fmt.Sprintf("%s %d %v\r\n", t.Version, t.Status, statusText[t.Status])

	for k, v := range t.Header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += fmt.Sprintf("Content-Length: %v\r\n\r\n%s", len(t.Body), t.Body)

	return []byte(msg), nil
}

// Decode http响应报文->结构体
func (t *Response) Decode(buffer []byte, bufferLen int) error {
	bufferStr := string(buffer)

	//找到 \r\n\r\n 的位置，用这个位置可以分隔请求头和请求体
	rnrnIndex := strings.Index(bufferStr, "\r\n\r\n")
	if rnrnIndex <= 0 {
		return ErrParseRespFailed
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
			return ErrParseRespFailed
		}
		t.ContentLen = cl
	}

	t.MsgLen = t.HeaderLen + t.ContentLen
	if t.MsgLen > bufferLen {
		// 计算出来的报文长度大于接收缓冲区中数据的长度
		return ErrParseRespFailed
	}
	t.Msg = buffer[0:t.MsgLen]

	err := t.parseHeader()
	if err != nil {
		return err
	}

	t.Body = string(t.Msg[t.HeaderLen:])

	return nil
}

func (t *Response) parseHeader() error {
	header := string(t.Msg[:t.HeaderLen])
	headerSplit := strings.Split(header, "\r\n")

	//响应行
	statusSplit := strings.Split(headerSplit[0], " ")
	t.Version = statusSplit[0]
	statusCode, err := strconv.Atoi(statusSplit[1])
	if err != nil {
		return ErrParseRespFailed
	}
	t.Status = statusCode

	//响应头
	for _, v := range headerSplit[1:] {
		vSplit := strings.Split(v, ": ") //用 ": " 切成键值
		if len(vSplit) == 2 {
			t.Header[vSplit[0]] = vSplit[1]
		}
	}

	return nil
}
