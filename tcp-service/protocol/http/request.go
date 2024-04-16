package http

import "fmt"

// 请求方法
const (
	MethodGet  string = "GET"
	MethodPost string = "POST"
)

// Request 请求
type Request struct {
	method string            //请求方法
	url    string            //请求路由
	param  map[string]string //请求参数
	header map[string]string //请求头
}

func NewRequest() *Request {
	return &Request{
		param:  make(map[string]string, 2),
		header: make(map[string]string, 2),
	}
}

func (r *Request) SetMethod(method string) {
	r.method = method
}

func (r *Request) SetUrl(url string) {
	r.url = url
}

func (r *Request) SetParam(k string, v string) {
	r.param[k] = v
}

func (r *Request) SetHeader(k string, v string) {
	r.header[k] = v
}

// MakeMsg 构造请求报文
func (r *Request) MakeMsg(body string) string {
	msg := fmt.Sprintf("%s %s", r.method, r.url)

	if len(r.param) > 0 {
		msg += "?"
		for k, v := range r.header {
			msg += fmt.Sprintf("%s: %s", k, v)
		}
		msg += " HTTP/1.1\r\n"
	}

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
