package middleware

import (
	"fmt"
	"net/http"
)

// HTTPHandleFunc 路由对应的处理方法的定义
type HTTPHandleFunc func(p7ctx *HTTPContext)

// HTTPServiceInterface 核心服务的接口定义
type HTTPServiceInterface interface {
	http.Handler
	Start(addr string) error
	MiddlewareInterface
}

// HTTPService 核心服务
type HTTPService struct {
	s5middleware []HTTPMiddleware
}

// 确保 HTTPService 实现了 HTTPServiceInterface 接口
var _ HTTPServiceInterface = &HTTPService{}

func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

func (p7this *HTTPService) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}

	// 倒过来组装，先组装的在里层，里层的后执行
	// 最里层应该是找路由然后执行业务代码
	t4chain := p7this.doServeHTTP
	for i := len(p7this.s5middleware) - 1; i > 0; i-- {
		t4chain = p7this.s5middleware[i](t4chain)
	}
	// 写入响应数据这个中间件应该由框架开发者处理
	// 它是最后一个环节，应该在最外层
	t4m := FlashRespMiddleware()
	t4chain = t4m(t4chain)
	t4chain(p7ctx)
}

func (p7this *HTTPService) doServeHTTP(p7ctx *HTTPContext) {
	p7ctx.RespStatusCode = http.StatusOK
	respData := fmt.Sprintf("doServeHTTP:%s %s\r\n", p7ctx.P7request.Method, p7ctx.P7request.URL.Path)
	respData += fmt.Sprintf("ReqBody:%s\r\n", p7ctx.ReqBody)
	p7ctx.RespData = append(p7ctx.RespData, []byte(respData)...)
}

func (p7this *HTTPService) Start(addr string) error {
	return http.ListenAndServe(addr, p7this)
}

func (p7this *HTTPService) AddMiddleware(s5mw ...HTTPMiddleware) {
	if nil == p7this.s5middleware {
		p7this.s5middleware = s5mw
		return
	}
	p7this.s5middleware = append(p7this.s5middleware, s5mw...)
}
