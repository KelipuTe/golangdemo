package http_service_v1

import (
  "fmt"
  "net/http"
)

// 可用于确保 HTTPHandlerMap 实现 HTTPHandler 接口。
// 如果 HTTPHandlerMap 没有实现 HTTPHandler 接口，这里就会报错。
var _ HTTPHandler = &HTTPHandlerMap{}

type HTTPHandlerMap struct {
  mapRoute map[string]HTTPHandlerFunc
}

func NewHTTPHandlerMap() *HTTPHandlerMap {
  return &HTTPHandlerMap{
    mapRoute: make(map[string]HTTPHandlerFunc),
  }
}

func (h *HTTPHandlerMap) HandlerHTTP(c *Context) {
  p1Req := c.P1Req
  fmt.Printf("HTTPHandlerMap,HandlerHTTP,Method=%s,URL.Path=%s\n", p1Req.Method, p1Req.URL.Path)

  handler, ok := h.mapRoute[p1Req.Method+"#"+p1Req.URL.Path]
  if !ok {
    c.ResW.WriteHeader(http.StatusNotFound)
    _, _ = c.ResW.Write([]byte("not any router match"))
    return
  }

  handler(c)
}

func (h *HTTPHandlerMap) Route(method string, pattern string, hhFunc HTTPHandlerFunc) {
  h.mapRoute[method+"#"+pattern] = hhFunc
}
