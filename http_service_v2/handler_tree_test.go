package http_service_v2

import (
	"net/http"
	"testing"
)

func TestHTTPHandlerTreeRegisteRoute(t *testing.T) {
  var p1h HTTPHandler = NewHTTPHandlerTree()
  p1h.RegisteRoute(http.MethodGet, "/user", func(p1c *HTTPContext) {})
  p1h.RegisteRoute(http.MethodGet, "/user/info", func(p1c *HTTPContext) {})
  p1h.RegisteRoute(http.MethodGet, "/user/order", func(p1c *HTTPContext) {})
}
