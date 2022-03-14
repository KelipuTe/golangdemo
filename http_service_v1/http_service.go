package http_service_v1

import (
  "fmt"
  "net/http"
)

type Service interface {
  Start(addr string, port string) error
  Route(method string, pattern string, hhFunc HTTPHandlerFunc)
}

type HTTPService struct {
  Name    string
  handler HTTPHandler
}

func NewHTTPSrevice(name string) Service {
  return &HTTPService{
    Name:    name,
    handler: NewHTTPHandlerMap(),
  }
}

// ServeHTTP 需要实现 http.ListenAndServe(addr string, handler Handler) 的第二个参数。
// Handler 即源码文件 src/net/http/service.go 里的 Handler 接口。
// http.ListenAndServe() 会调用 net.Listen(network, address string)。
// 这里 network 参数设置的是 "tcp"，因为 HTTP/1.1 和 HTTP/2 是基于 TCP 的。
// net.Listen() 底层在监听到 address 上的 TCP 连接时，
// 最终会调用 Handler 接口的实现 也就是这里的 ServeHTTP 处理。
func (s *HTTPService) ServeHTTP(resW http.ResponseWriter, p1Req *http.Request) {
  fmt.Printf("HTTPService,ServeHTTP\n")
  c := NewContext(resW, p1Req)
  s.handler.HandlerHTTP(c)
}

func (s *HTTPService) Start(addr string, port string) error {
  fmt.Printf("%s start at %s.\n", s.Name, addr+":"+port)
  return http.ListenAndServe(addr+":"+port, s)
}

func (s *HTTPService) Route(method string, pattern string, hhFunc HTTPHandlerFunc) {
  s.handler.Route(method, pattern, hhFunc)
}
