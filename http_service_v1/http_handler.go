package http_service_v1

// HTTPRoute 路由接口
type HTTPRoute interface {
  // RegisteRoute 注册路由。
  // method HTTP 方法；
  // pattern 路由；
  RegisteRoute(method string, pattern string, hhFunc HTTPHandlerFunc)
}

// HTTPHandler HTTP 处理接口
type HTTPHandler interface {
  // HandlerHTTP
  HandlerHTTP(c *Context)
  HTTPRoute
}

// HTTPHandlerFunc 路由对应的 HTTP 处理函数
type HTTPHandlerFunc func(c *Context)
