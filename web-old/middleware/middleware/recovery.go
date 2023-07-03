package middleware

import "demo-golang/web-old/middleware"

func RecoveryMiddleware() middleware.HTTPMiddleware {
	return func(next middleware.HTTPHandleFunc) middleware.HTTPHandleFunc {
		return func(p7ctx *middleware.HTTPContext) {
			defer func() {
				if err := recover(); err != nil {
					p7ctx.RespData = append(p7ctx.RespData, []byte("recover from panic;")...)
				}
			}()
			p7ctx.RespData = append(p7ctx.RespData, []byte("RecoveryMiddleware In;")...)
			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("RecoveryMiddleware Out;")...)
		}
	}
}
