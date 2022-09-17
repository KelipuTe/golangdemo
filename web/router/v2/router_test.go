package v2

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_router_addRoute(t *testing.T) {
	s5route := []struct {
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
			path:   "/order",
		},
		{
			method: http.MethodGet,
			path:   "/order/detail",
		},
		{
			method: http.MethodPost,
			path:   "/order/create",
		},
		{
			method: http.MethodPost,
			path:   "/order/delete",
		},
	}

	t4router := newRouter()
	f4handler := func(p7ctx *HTTPContext) { fmt.Println(p7ctx) }
	for _, i4r := range s5route {
		t4router.addRoute(i4r.method, i4r.path, f4handler)
	}

	fmt.Println(t4router)

	//wantRouter := &router{
	//	trees: map[string]*node{
	//		http.MethodGet: {
	//			path: "/",
	//			children: map[string]*node{
	//				"user": {
	//					path: "user",
	//					children: map[string]*node{
	//						"home": {path: "home", handler: mockHandler, typ: nodeTypeStatic},
	//					},
	//					handler: mockHandler,
	//					typ:     nodeTypeStatic,
	//				},
	//				"order": {
	//					path: "order",
	//					children: map[string]*node{
	//						"detail": {path: "detail", handler: mockHandler, typ: nodeTypeStatic},
	//					},
	//					starChild: &node{path: "*", handler: mockHandler, typ: nodeTypeAny},
	//					typ:       nodeTypeStatic,
	//				},
	//				"param": {
	//					path: "param",
	//					paramChild: &node{
	//						path:      ":id",
	//						paramName: "id",
	//						starChild: &node{
	//							path:    "*",
	//							handler: mockHandler,
	//							typ:     nodeTypeAny,
	//						},
	//						children: map[string]*node{"detail": {path: "detail", handler: mockHandler, typ: nodeTypeStatic}},
	//						handler:  mockHandler,
	//						typ:      nodeTypeParam,
	//					},
	//				},
	//			},
	//			starChild: &node{
	//				path: "*",
	//				children: map[string]*node{
	//					"abc": {
	//						path:      "abc",
	//						starChild: &node{path: "*", handler: mockHandler, typ: nodeTypeAny},
	//						handler:   mockHandler,
	//						typ:       nodeTypeStatic,
	//					},
	//				},
	//				starChild: &node{path: "*", handler: mockHandler, typ: nodeTypeAny},
	//				handler:   mockHandler,
	//				typ:       nodeTypeAny,
	//			},
	//			handler: mockHandler,
	//			typ:     nodeTypeStatic,
	//		},
	//		http.MethodPost: {
	//			path: "/",
	//			children: map[string]*node{
	//				"order": {path: "order", children: map[string]*node{
	//					"create": {path: "create", handler: mockHandler, typ: nodeTypeStatic},
	//				}},
	//				"login": {path: "login", handler: mockHandler, typ: nodeTypeStatic},
	//			},
	//			typ: nodeTypeStatic,
	//		},
	//		http.MethodDelete: {
	//			path: "/",
	//			children: map[string]*node{
	//				"reg": {
	//					path: "reg",
	//					typ:  nodeTypeStatic,
	//					regChild: &node{
	//						path:      ":id(.*)",
	//						paramName: "id",
	//						typ:       nodeTypeReg,
	//						handler:   mockHandler,
	//					},
	//				},
	//			},
	//			regChild: &node{
	//				path:      ":name(^.+$)",
	//				paramName: "name",
	//				typ:       nodeTypeReg,
	//				children: map[string]*node{
	//					"abc": {
	//						path:    "abc",
	//						handler: mockHandler,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
}
