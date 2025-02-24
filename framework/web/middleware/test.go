package middleware

import (
	web2 "demo-golang/framework/web"
)

// [中间件]测试用的，看看套的对不对
func F8TestMiddleware(code string) web2.F8HTTPMiddlewareFunc {
	code = "[" + code + "]"
	return func(next web2.F8HTTPHandlerFunc) web2.F8HTTPHandlerFunc {
		return func(p7ctx *web2.S6HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte(code)...)

			next(p7ctx)
		}
	}
}
