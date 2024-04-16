package http

import "fmt"

// 状态码
const (
	StatusOk                  uint16 = 200
	StatusBadRequest          uint16 = 400
	StatusNotFound            uint16 = 404
	StatusInternalServerError uint16 = 500
)

// 状态码文案
var statusText = map[uint16]string{
	StatusOk:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
}

// Response 响应
type Response struct {
	statusCode uint16            //状态码
	header     map[string]string //响应头
}

func NewResponse() *Response {
	return &Response{
		header: make(map[string]string, 2),
	}
}

func (r *Response) SetStatusCode(code uint16) {
	r.statusCode = code
}

func (r *Response) SetHeader(k string, v string) {
	r.header[k] = v
}

// MakeMsg 构造响应报文
func (r *Response) MakeMsg(body string) string {
	msg := fmt.Sprintf("HTTP/1.1 %d %v\r\n", r.statusCode, statusText[r.statusCode])

	_, ok := r.header["Content-Type"]
	if !ok {
		r.header["Content-Type"] = "text/html; charset=utf8"
	}

	for k, v := range r.header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += fmt.Sprintf("Content-Length: %v\r\n\r\n%s", len(body), body)

	return msg
}
