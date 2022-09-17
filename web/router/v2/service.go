package v2

import (
	"net/http"
)

type HTTPHandleFunc func(p7ctx *HTTPContext)

type HTTPServiceInterface interface {
	RouterInterface
	http.Handler
	Start(addr string) error
}

type HTTPService struct {
	router
}

// 确保 HTTPService 实现了 HTTPServiceInterface 接口
var _ HTTPServiceInterface = &HTTPService{}

func NewHTTPService() *HTTPService {
	return &HTTPService{
		router: newRouter(),
	}
}

// Get 包装 addRoute
func (p7this *HTTPService) Get(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodGet, path, f4h)
}

// Post 包装 addRoute
func (p7this *HTTPService) Post(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodPost, path, f4h)
}

func (p7this *HTTPService) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}
	p7this.doServeHTTP(p7ctx)
}

func (p7this *HTTPService) doServeHTTP(p7ctx *HTTPContext) {
	p7ri := p7this.findRoute(p7ctx.P7request.Method, p7ctx.P7request.URL.Path)
	if nil == p7ri || nil == p7ri.p7node || nil == p7ri.p7node.f4handler {
		p7ctx.I9writer.WriteHeader(404)
		p7ctx.I9writer.Write([]byte("Not Found"))
		return
	}
	p7ctx.M3pathParam = p7ri.m3pathParam
	p7ri.p7node.f4handler(p7ctx)
}

func (p7this *HTTPService) Start(addr string) error {
	return http.ListenAndServe(addr, p7this)
}
