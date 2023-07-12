package middleware

import (
	"demo-golang/web"
)

func TestMiddleware(code string) web.HTTPMiddleware {
	code = "[" + code + "]"
	return func(next web.HTTPHandleFunc) web.HTTPHandleFunc {
		return func(p7ctx *web.HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte(code)...)
			next(p7ctx)
		}
	}
}
