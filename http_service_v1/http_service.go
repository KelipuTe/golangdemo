package http_service_v1

import (
  "fmt"
  "net/http"
)

// Service 服务接口
type Service interface {
  // Start 服务启动
  Start(addr string, port string) error
  HTTPRoute
}

// HTTPService HTTP 服务
type HTTPService struct {
  // Name 服务的名字
  Name string
  // HTTPHandler 的实例
  handler HTTPHandler
  // 中间件入口函数
  entrance MiddlewareFunc
}

func NewHTTPSrevice(name string, arr1Builder ...MiddlewareBuilder) Service {
  // 这里实例化一个 HTTPHandler
  var h HTTPHandler = NewHTTPHandlerMap()
  // HTTPHandler 的 HandlerHTTP 函数，就是路由对应的处理函数。
  var hf MiddlewareFunc = h.HandlerHTTP

  // 中间件建造器数组反过来遍历。像洋葱一样，数组最前面的对应最外层。
  // 套娃完成后，HandlerHTTP 函数应该在最里面。表示请求通过层层中间件后进入业务逻辑。
  for i := len(arr1Builder) - 1; i > -1; i-- {
    var mf MiddlewareBuilder = arr1Builder[i]
    hf = mf(hf)
  }

  return &HTTPService{
    Name:     name,
    handler:  h,
    entrance: hf,
  }
}

// #### Handler 接口实现 ####

// ServeHTTP 需要实现 http.ListenAndServe(addr string, handler Handler) 的第二个参数。
// Handler 即源码文件 src/net/http/service.go 里的 Handler 接口。
// http.ListenAndServe() 会调用 net.Listen(network, address string)。
// 这里 network 参数设置的是 "tcp"，因为 HTTP/1.1 和 HTTP/2 是基于 TCP 的。
// net.Listen() 底层在监听到 address 上的 TCP 连接时，
// 最终会调用 Handler 接口的实现，也就是这里的 ServeHTTP 处理。
func (s *HTTPService) ServeHTTP(resW http.ResponseWriter, p1Req *http.Request) {
  c := NewContext(resW, p1Req)

  // 不使用中间件时，直接调用 HTTPHandler 的实例处理请求
  // s.handler.HandlerHTTP(c)

  // 使用中间件后，这里就要改成调用中间件入口
  s.entrance(c)
}

// ## Handler ##

// #### Service 接口实现 ####

func (s *HTTPService) Start(addr string, port string) error {
  fmt.Printf("HTTPService %s start at %s...\n", s.Name, addr+":"+port)
  return http.ListenAndServe(addr+":"+port, s)
}

func (s *HTTPService) RegisteRoute(method string, pattern string, hhFunc HTTPHandlerFunc) {
  s.handler.RegisteRoute(method, pattern, hhFunc)
}

// ## Service ##
