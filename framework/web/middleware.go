package web

import "log"

// 中间件的处理方法
type F8HTTPMiddlewareFunc func(next F8HTTPHandlerFunc) F8HTTPHandlerFunc

// 中间件
type I9Middleware interface {
	// 添加中间件
	F8AddMiddleware(...F8HTTPMiddlewareFunc)
}

// [中间件]把自定义请求上下文里的响应数据写入 http 请求响应
func F8FlashRespMiddleware() F8HTTPMiddlewareFunc {
	return func(next F8HTTPHandlerFunc) F8HTTPHandlerFunc {
		return func(p7ctx *S6HTTPContext) {
			next(p7ctx)

			if p7ctx.RespStatusCode > 0 {
				p7ctx.I9ResponseWriter.WriteHeader(p7ctx.RespStatusCode)
			}
			_, err := p7ctx.I9ResponseWriter.Write(p7ctx.RespData)
			if err != nil {
				log.Fatalln("flashResp failed", err)
			}
		}
	}
}
