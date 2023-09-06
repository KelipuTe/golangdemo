package middleware

import (
	"demo-golang/web"
	"fmt"
)

// [中间件]记录日志
func F8LogMiddleware() web.F8HTTPMiddlewareFunc {
	return func(next web.F8HTTPHandlerFunc) web.F8HTTPHandlerFunc {
		return func(p7ctx *web.S6HTTPContext) {
			fmt.Printf("request path:%s\r\n", p7ctx.P7Request.URL.Path)
			fmt.Println("ReqBody:", string(p7ctx.ReqBody))
			next(p7ctx)
			fmt.Println("RespData:", string(p7ctx.RespData))
		}
	}
}
