package router

import (
	"fmt"
	"net/http"
)

// HTTPHandleFunc 路由对应的处理方法的定义
type HTTPHandleFunc func(p7ctx *HTTPContext)

// HTTPHandlerInterface 核心处理逻辑的接口定义
type HTTPHandlerInterface interface {
	http.Handler
	RouterInterface
}

// HTTPHandler 核心处理逻辑
type HTTPHandler struct {
	router
}

// 确保 HTTPHandler 实现了 HTTPHandlerInterface 接口
var _ HTTPHandlerInterface = &HTTPHandler{}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		router: newRouter(),
	}
}

func (p7this *HTTPHandler) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}
	p7this.doServeHTTP(p7ctx)
}

func (p7this *HTTPHandler) doServeHTTP(p7ctx *HTTPContext) {
	p7ri := p7this.findRoute(p7ctx.P7request.Method, p7ctx.P7request.URL.Path)
	// 如果找不到对应的路由结点或者路由结点上没有处理方法就返回 404
	if nil == p7ri || nil == p7ri.p7node || nil == p7ri.p7node.f4handler {
		p7ctx.I9writer.WriteHeader(404)
		_, _ = p7ctx.I9writer.Write([]byte(fmt.Sprintf("Not Found:%s %s\r\n", p7ctx.P7request.Method, p7ctx.P7request.URL.Path)))
		return
	}
	// 这里可以把匹配结果存下来
	p7ctx.M3pathParam = p7ri.m3pathParam
	p7ctx.p7routingNode = p7ri.p7node
	// 执行业务代码
	p7ri.p7node.f4handler(p7ctx)
}

// Get 包装 addRoute
func (p7this *HTTPHandler) Get(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodGet, path, f4h)
}

// Post 包装 addRoute
func (p7this *HTTPHandler) Post(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodPost, path, f4h)
}
