package middleware

import (
	"bytes"
	"demo-golang/web"
	"io"
)

// [中间件]记录请求参数
func F8ReqBodyMiddleware() web.F8HTTPMiddlewareFunc {
	return func(next web.F8HTTPHandlerFunc) web.F8HTTPHandlerFunc {
		return func(p7ctx *web.S6HTTPContext) {
			//先从流里把数据读出来，记录到自定义上下文里
			var err error
			p7ctx.ReqBody, err = io.ReadAll(p7ctx.P7Request.Body)
			if nil != err {
				return
			}
			//构造一个流放回去，防止下游直接读原始的请求流
			p7ctx.P7Request.Body = io.NopCloser(bytes.NewBuffer(p7ctx.ReqBody))

			next(p7ctx)
		}
	}
}
