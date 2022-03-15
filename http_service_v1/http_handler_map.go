package http_service_v1

import (
  "fmt"
  "net/http"
)

// 这个写法可用于确保 HTTPHandlerMap 实现 HTTPHandler 接口。
// 如果 HTTPHandlerMap 没有实现 HTTPHandler 接口，这里就会报错。
var _ HTTPHandler = &HTTPHandlerMap{}

// HTTPHandlerMap 基于 map 实现路由处理
type HTTPHandlerMap struct {
  mapRoute map[string]HTTPHandlerFunc
}

func NewHTTPHandlerMap() *HTTPHandlerMap {
  return &HTTPHandlerMap{
    mapRoute: make(map[string]HTTPHandlerFunc),
  }
}

// #### HTTPHandler 接口实现 ####

func (h *HTTPHandlerMap) HandlerHTTP(c *Context) {
  p1Req := c.P1Req
  fmt.Printf("HTTPHandlerMap, HandlerHTTP, p1Req.Method: %s, p1Req.URL.Path: %s\n", p1Req.Method, p1Req.URL.Path)

  // 路由查询，找到对应的处理函数
  handler, ok := h.mapRoute[p1Req.Method+"#"+p1Req.URL.Path]
  if !ok {
    c.ResW.WriteHeader(http.StatusNotFound)
    _, _ = c.ResW.Write([]byte("route not found"))
    return
  }

  handler(c)
}

func (h *HTTPHandlerMap) RegisteRoute(method string, pattern string, hhFunc HTTPHandlerFunc) {
  // 这里用 HTTP 方法和路由构造一个唯一键，实现区分 HTTP 方法
  h.mapRoute[method+"#"+pattern] = hhFunc
}

// ## HTTPHandler ##
