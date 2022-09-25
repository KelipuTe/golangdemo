package middleware

import (
	"demo-golang/web/middleware"
	"fmt"
)

func LogMiddleware() middleware.HTTPMiddleware {
	return func(next middleware.HTTPHandleFunc) middleware.HTTPHandleFunc {
		return func(p7ctx *middleware.HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte("LogMiddleware In\r\n")...)
			fmt.Printf("request path:%s\r\n", p7ctx.P7request.URL.Path)
			fmt.Println("ReqBody:", string(p7ctx.ReqBody))
			next(p7ctx)
			fmt.Println("RespData:", p7ctx.RespData)
			p7ctx.RespData = append(p7ctx.RespData, []byte("LogMiddleware Out\r\n")...)
		}
	}
}
