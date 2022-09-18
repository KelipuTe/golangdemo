package middleware

import (
	"fmt"
	"net/http"
)

type HTTPHandleFunc func(p7ctx *HTTPContext)

type HTTPServiceInterface interface {
	http.Handler
	Start(addr string) error
	MiddlewareInterface
}

type HTTPService struct {
	s5middleware []HTTPMiddleware
}

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
	// 最里层应该是找路由然后执行代码的逻辑
	t4chain := p7this.doServeHTTP
	for i := len(p7this.s5middleware) - 1; i > 0; i-- {
		t4chain = p7this.s5middleware[i](t4chain)
	}
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
