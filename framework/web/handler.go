package web

import (
	"fmt"
	"net/http"
	"sync"
)

// 路由对应的处理方法
type F8HTTPHandlerFunc func(*S6HTTPContext)

// http 处理逻辑
type I9HTTPHandler interface {
	http.Handler
	I9Router
	I9Middleware
}

// 确保 S6HTTPHandler 实现了 I9HTTPHandler 接口
var _ I9HTTPHandler = &S6HTTPHandler{}

// http 处理逻辑
type S6HTTPHandler struct {
	s6Router
	s5f8MiddlewareFunc []F8HTTPMiddlewareFunc //全局中间件
	ctxPool            sync.Pool              //内存池，复用 S6HTTPContext
	isRunning          bool                   //是否正在运行
}

func NewS6HTTPHandler() *S6HTTPHandler {
	return &S6HTTPHandler{
		s6Router: newRouter(),
		ctxPool: sync.Pool{
			New: func() interface{} {
				return NewHTTPContext()
			},
		},
		isRunning: true,
	}
}

func (p7this *S6HTTPHandler) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	//从内存池里获取 S6HTTPContext
	p7s6ctx := p7this.ctxPool.Get().(*S6HTTPContext)
	//归还资源到资源池
	defer func() {
		p7s6ctx.F8Reset()
		p7this.ctxPool.Put(p7s6ctx)
	}()
	p7s6ctx.I9ResponseWriter = i9w
	p7s6ctx.P7Request = p7r

	//倒过来组装，先组装的在里层，里层的后执行
	//最里层应该是找路由然后执行业务代码
	t4chain := p7this.f8DoServeHTTP
	for i := len(p7this.s5f8MiddlewareFunc) - 1; i >= 0; i-- {
		t4chain = p7this.s5f8MiddlewareFunc[i](t4chain)
	}
	//写入响应数据这个中间件应该由框架开发者处理
	//它是最后一个环节，应该在最外层
	t4m := F8FlashRespMiddleware()
	t4chain = t4m(t4chain)
	t4chain(p7s6ctx)
}

func (p7this *S6HTTPHandler) f8DoServeHTTP(p7s6ctx *S6HTTPContext) {
	//如果服务已经关闭了就直接返回
	if !p7this.isRunning {
		p7s6ctx.I9ResponseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = p7s6ctx.I9ResponseWriter.Write([]byte("服务已关闭"))
		return
	}
	p7s6ri := p7this.f8FindRoute(p7s6ctx.P7Request.Method, p7s6ctx.P7Request.URL.Path)
	//如果找不到对应的路由结点或者路由结点上没有处理方法就返回 404
	if nil == p7s6ri || nil == p7s6ri.p7node || nil == p7s6ri.p7node.f8HandlerFunc {
		p7s6ctx.I9ResponseWriter.WriteHeader(http.StatusNotFound)
		errStr := fmt.Sprintf("Not Found:%s %s\r\n", p7s6ctx.P7Request.Method, p7s6ctx.P7Request.URL.Path)
		_, _ = p7s6ctx.I9ResponseWriter.Write([]byte(errStr))
		return
	}
	//这里可以把匹配结果存下来
	p7s6ctx.M3PathParam = p7s6ri.m3PathParam
	p7s6ctx.p7RoutingNode = p7s6ri.p7node
	//这里用同样的套路，处理路由上的中间件，最后执行业务代码
	t4chain := p7s6ri.p7node.f8HandlerFunc
	for i := len(p7s6ri.p7node.s5f8MiddlewareCache) - 1; i >= 0; i-- {
		t4chain = p7s6ri.p7node.s5f8MiddlewareCache[i](t4chain)
	}
	t4chain(p7s6ctx)
}

// 包装 f8AddRoute
func (p7this *S6HTTPHandler) F8Get(path string, f8hf F8HTTPHandlerFunc, s5f8mw ...F8HTTPMiddlewareFunc) {
	p7this.s6Router.f8AddRoute(http.MethodGet, path, f8hf, s5f8mw...)
}

// 包装 f8AddRoute
func (p7this *S6HTTPHandler) F8Post(path string, f8hf F8HTTPHandlerFunc, s5f8mw ...F8HTTPMiddlewareFunc) {
	p7this.s6Router.f8AddRoute(http.MethodPost, path, f8hf, s5f8mw...)
}

// 路由数据
type S6RouteData struct {
	Method   string
	Path     string
	F4handle F8HTTPHandlerFunc
}

// 添加路由组，思路很简单，给一组路由加一个同样的前缀，然后都用一样的中间件
func (p7this *S6HTTPHandler) F8Group(pathPrefix string, s5f8mw []F8HTTPMiddlewareFunc, s5s6rd []S6RouteData) {
	for _, rd := range s5s6rd {
		t4path := pathPrefix
		if "/" != rd.Path {
			t4path = pathPrefix + rd.Path
		}
		p7this.f8AddRoute(rd.Method, t4path, rd.F4handle, s5f8mw...)
	}
}

func (p7this *S6HTTPHandler) F8AddMiddleware(s5f8mw ...F8HTTPMiddlewareFunc) {
	if nil == p7this.s5f8MiddlewareFunc {
		p7this.s5f8MiddlewareFunc = make([]F8HTTPMiddlewareFunc, 0, len(s5f8mw))
	}
	p7this.s5f8MiddlewareFunc = append(p7this.s5f8MiddlewareFunc, s5f8mw...)
}
