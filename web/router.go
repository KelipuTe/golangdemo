package web

import (
	"errors"
	"strings"
)

var (
	ErrPathCannotBeEmpty      = errors.New("路由不能为空")
	ErrPartCannotBeEmpty      = errors.New("路由不能存在连续的 `\\`")
	ErrPathMustStartWithSlash = errors.New("路由必须以 `\\` 开始")
	ErrPathCannotEndWithSlash = errors.New("路由不能以 `\\` 结尾")
	ErrPathExist              = errors.New("路由重复注册")
)

// 路由器
type I9Router interface {
	//添加路由
	f8AddRoute(method string, path string, f8hf F8HTTPHandlerFunc, s5f8mf ...F8HTTPMiddlewareFunc)
	//查找路由
	f8FindRoute(method string, path string) *s6RouteInfo
	//服务器动前，将中间件扫描好，缓存到路由结点
	f8CacheMiddleware()
}

// 路由器
type s6Router struct {
	//路由树，路由按请求方式分成多个路由树
	//map：F8Get => F8Get 的路由树；F8Post => F8Post 的路由树。
	m3RoutingTree map[string]*s6RoutingNode
}

func newRouter() s6Router {
	return s6Router{
		m3RoutingTree: map[string]*s6RoutingNode{},
	}
}

func (p7this *s6Router) f8AddRoute(method string, path string, f8hf F8HTTPHandlerFunc, s5f8mf ...F8HTTPMiddlewareFunc) {
	if "" == path {
		panic(ErrPathCannotBeEmpty)
	}
	if '/' != path[0] {
		panic(ErrPathMustStartWithSlash)
	}
	if '/' == path[len(path)-1] && "/" != path {
		panic(ErrPathCannotEndWithSlash)
	}

	//按 http method 区分路由树
	p7node, ok := p7this.m3RoutingTree[method]
	if !ok {
		//创建路由树根结点
		p7node = &s6RoutingNode{
			nodeType: nodeTypeStatic,
			part:     "/",
			path:     "/",
		}
		p7this.m3RoutingTree[method] = p7node
	}
	//处理根路由
	if "/" == path {
		if nil != p7node.f8HandlerFunc {
			panic(ErrPathExist)
		}
		p7node.f8HandlerFunc = f8hf
		return //处理完直接返回
	}
	//分段处理路由
	s5path := strings.Split(path[1:], "/")
	t4path := ""
	for _, itemPart := range s5path {
		if "" == itemPart {
			panic(ErrPartCannotBeEmpty)
		}
		t4path += "/" + itemPart
		t4p7child := p7node.findChild(itemPart)
		if nil == t4p7child {
			t4p7child = p7node.createChild(itemPart, t4path)
		} else {
			t4p7child.checkChild(itemPart)
		}
		p7node = t4p7child
	}
	//给路由添加处理方法
	if nil != p7node.f8HandlerFunc {
		panic(ErrPathExist)
	}
	p7node.f8HandlerFunc = f8hf
	p7node.s5f8Middleware = s5f8mf //给路由添加中间件
}

func (p7this *s6Router) f8FindRoute(method string, path string) *s6RouteInfo {
	p7node, ok := p7this.m3RoutingTree[method]
	if !ok {
		return nil
	}

	if "/" == path {
		return &s6RouteInfo{
			p7node: p7node,
		}
	}

	p7ri := &s6RouteInfo{}
	s5path := strings.Split(path[1:], "/")
	for _, part := range s5path {
		p7node = p7node.matchChild(part)
		if nil == p7node {
			return nil
		}
		if p7node.nodeType == nodeTypeRegexp {
			s5res := p7node.p7RegExp.FindStringSubmatch(part)
			if nil != s5res {
				p7ri.f8AddPathParam(p7node.paramName, s5res[0])
			} else {
				p7ri.f8AddPathParam(p7node.paramName, "")
			}
		} else if p7node.nodeType == nodeTypeParam {
			p7ri.f8AddPathParam(p7node.paramName, part)
		} else if p7node.nodeType == nodeTypeAny {
			if nil == p7node.m3RoutingTree && nil == p7node.p7RegExpChild && nil == p7node.p7ParamChild && nil == p7node.p7AnyChild {
				break
			}
		}
	}
	p7ri.p7node = p7node
	return p7ri
}

func (p7this *s6Router) f8CacheMiddleware() {
	//遍历每一个路由树
	for _, p7node := range p7this.m3RoutingTree {
		p7node.f8CacheMiddleware(nil)
	}
}

// 路由查询结果
type s6RouteInfo struct {
	p7node      *s6RoutingNode    //命中的路由结点
	m3PathParam map[string]string //提取出来的路径参数
}

// 添加路径参数
func (p7this *s6RouteInfo) f8AddPathParam(k string, v string) {
	if nil == p7this.m3PathParam {
		p7this.m3PathParam = map[string]string{k: v}
	}
	p7this.m3PathParam[k] = v
}
