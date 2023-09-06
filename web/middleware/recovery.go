package middleware

import "demo-golang/web"

// [中间件]panic恢复
func F8RecoveryMiddleware() web.F8HTTPMiddlewareFunc {
	return func(next web.F8HTTPHandlerFunc) web.F8HTTPHandlerFunc {
		return func(p7ctx *web.S6HTTPContext) {
			defer func() {
				if err := recover(); err != nil {
					p7ctx.RespData = append(p7ctx.RespData, []byte("recover from panic\r\n")...)
				}
			}()
			next(p7ctx)
		}
	}
}
