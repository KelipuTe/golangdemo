package http_service_v1

type HTTPHandler interface {
  HandlerHTTP(c *Context)
  Route(method string, pattern string, hhFunc HTTPHandlerFunc)
}

type HTTPHandlerFunc func(c *Context)
