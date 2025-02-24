package middleware

import (
	web2 "demo-golang/framework/web"
)

// [中间件]panic恢复
func F8RecoveryMiddleware() web2.F8HTTPMiddlewareFunc {
	return func(next web2.F8HTTPHandlerFunc) web2.F8HTTPHandlerFunc {
		return func(p7ctx *web2.S6HTTPContext) {
			defer func() {
				if err := recover(); err != nil {
					p7ctx.RespData = append(p7ctx.RespData, []byte("recover from panic\r\n")...)
				}
			}()
			next(p7ctx)
		}
	}
}
