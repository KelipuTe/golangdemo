package v2

import (
	"fmt"
	"strings"
)

type RouterInterface interface {
	addRoute(method string, path string, f4handler HTTPHandleFunc)
}

// router 路由器
type router struct {
	// m3routingTree 路由树，路由按请求方式分成多个路由树
	// map：Get => Get 的路由树；Post => Post 的路由树。
	m3routingTree map[string]*routingNode
}

func newRouter() router {
	return router{
		m3routingTree: map[string]*routingNode{},
	}
}

func (p7this *router) addRoute(method string, path string, f4h HTTPHandleFunc) {
	if "" == path {
		panic("\"\" == path")
	}
	if '/' != path[0] {
		panic("'/' != path[0]")
	}
	if "/" != path && '/' == path[len(path)-1] {
		panic("\"/\" != path && '/' == path[len(path)-1]")
	}

	p7node, ok := p7this.m3routingTree[method]
	if !ok {
		// 创建路由树根结点
		p7node = &routingNode{
			path: "/",
		}
	}
	if "/" == path {
		if nil != p7node.f4handler {
			panic("\"/\" == path && nil != p7node.f4handler")
		}
		p7node.f4handler = f4h
		return
	}

	s5path := strings.Split(path[1:], "/")
	for _, part := range s5path {
		if "" == part {
			panic(fmt.Sprintf("\"\" == part, path=[%s]", path))
		}
		t4p7child := p7node.findChild(part)
		if nil == t4p7child {
			t4p7child = createChild(part)
		}
		p7node = t4p7child
	}
}
