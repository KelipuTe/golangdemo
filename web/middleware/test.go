package middleware

import (
	"demo-golang/web"
)

// [中间件]测试用的，看看套的对不对
func F8TestMiddleware(code string) web.F8HTTPMiddlewareFunc {
	code = "[" + code + "]"
	return func(next web.F8HTTPHandlerFunc) web.F8HTTPHandlerFunc {
		return func(p7ctx *web.S6HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte(code)...)

			next(p7ctx)
		}
	}
}
