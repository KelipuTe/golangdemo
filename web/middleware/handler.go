package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPHandleFunc 路由对应的处理方法的定义
type HTTPHandleFunc func(p7ctx *HTTPContext)

// HTTPHandlerInterface 核心服务的接口定义
type HTTPHandlerInterface interface {
	http.Handler
	MiddlewareInterface
}

// HTTPHandler 核心服务
type HTTPHandler struct {
	s5f4middleware []HTTPMiddleware
}

// 确保 HTTPHandler 实现了 HTTPHandlerInterface 接口
var _ HTTPHandlerInterface = &HTTPHandler{}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (p7this *HTTPHandler) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}

	// 倒过来组装，先组装的在里层，里层的后执行
	// 最里层应该是找路由然后执行业务代码
	t4chain := p7this.doServeHTTP
	for i := len(p7this.s5f4middleware) - 1; i >= 0; i-- {
		t4chain = p7this.s5f4middleware[i](t4chain)
	}
	// 写入响应数据这个中间件应该由框架开发者处理
	// 它是最后一个环节，应该在最外层
	t4m := FlashRespMiddleware()
	t4chain = t4m(t4chain)
	t4chain(p7ctx)
}

func (p7this *HTTPHandler) doServeHTTP(p7ctx *HTTPContext) {
	p7ctx.RespStatusCode = http.StatusOK
	respData := fmt.Sprintf("doServeHTTP:%s %s;", p7ctx.P7request.Method, p7ctx.P7request.URL.Path)
	respData += fmt.Sprintf("ReqBody:%s;", p7ctx.ReqBody)
	p7ctx.RespData = append(p7ctx.RespData, []byte(respData)...)
}

func (p7this *HTTPHandler) AddMiddleware(s5f4mw ...HTTPMiddleware) {
	if nil == p7this.s5f4middleware {
		p7this.s5f4middleware = make([]HTTPMiddleware, 0, len(s5f4mw))
	}
	p7this.s5f4middleware = append(p7this.s5f4middleware, s5f4mw...)
}

// FlashRespMiddleware 写入响应数据
func FlashRespMiddleware() HTTPMiddleware {
	return func(next HTTPHandleFunc) HTTPHandleFunc {
		return func(p7ctx *HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte("FlashRespMiddleware In;")...)
			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("FlashRespMiddleware Out;")...)
			flashResp(p7ctx)
		}
	}
}

func flashResp(p7ctx *HTTPContext) {
	if p7ctx.RespStatusCode > 0 {
		p7ctx.I9writer.WriteHeader(p7ctx.RespStatusCode)
	}
	_, err := p7ctx.I9writer.Write(p7ctx.RespData)
	if err != nil {
		log.Fatalln("flashResp failed", err)
	}
}
