package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_router_addRoute(t *testing.T) {
	s5testRoute := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/hello",
		},
		{
			method: http.MethodGet,
			path:   "/hello/world",
		},
		{
			method: http.MethodGet,
			path:   "/hello/*",
		},
		{
			method: http.MethodGet,
			path:   "/order",
		},
		{
			method: http.MethodGet,
			path:   "/order/:size/:page",
		},
		{
			method: http.MethodGet,
			path:   "/order/:id/detail",
		},
		{
			method: http.MethodPost,
			path:   "/order/create",
		},
		{
			method: http.MethodPost,
			path:   "/order/:id/delete",
		},
	}

	t4router := newRouter()
	f4handler := func(p7ctx *HTTPContext) {}
	for _, i4r := range s5testRoute {
		t4router.addRoute(i4r.method, i4r.path, f4handler)
	}

	wantRouter := router{
		m3routingTree: map[string]*routingNode{
			http.MethodGet: {
				nodeType: nodeTypeStatic,
				part:     "/",
				path:     "/",
				m3routingTree: map[string]*routingNode{
					"hello": {
						nodeType: nodeTypeStatic,
						part:     "hello",
						path:     "/hello",
						m3routingTree: map[string]*routingNode{
							"world": {
								nodeType: nodeTypeStatic,
								part:     "world",
								path:     "/hello/world",
							},
						},
					},
				},
			},
		},
	}

	ok, msg := equalRouter(&t4router, &wantRouter)
	assert.True(t, ok, msg)
}

func equalRouter(p7a *router, p7b *router) (bool, string) {
	return true, ""
}
