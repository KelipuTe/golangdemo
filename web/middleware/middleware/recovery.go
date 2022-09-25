package middleware

import "demo-golang/web/middleware"

func RecoveryMiddleware() middleware.HTTPMiddleware {
	return func(next middleware.HTTPHandleFunc) middleware.HTTPHandleFunc {
		return func(p7ctx *middleware.HTTPContext) {
			defer func() {
				if err := recover(); err != nil {
					p7ctx.RespData = append(p7ctx.RespData, []byte("recover from panic\r\n")...)
				}
			}()
			p7ctx.RespData = append(p7ctx.RespData, []byte("RecoveryMiddleware In\r\n")...)
			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("RecoveryMiddleware Out\r\n")...)
		}
	}
}
