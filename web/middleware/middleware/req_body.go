package middleware

import (
	"bytes"
	"demo-golang/web/middleware"
	"io"
)

// ReqBodyMiddleware 提取请求参数
func ReqBodyMiddleware() middleware.HTTPMiddleware {
	return func(next middleware.HTTPHandleFunc) middleware.HTTPHandleFunc {
		return func(p7ctx *middleware.HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte("ReqBodyMiddleware In\r\n")...)

			var err error
			p7ctx.ReqBody, err = io.ReadAll(p7ctx.P7request.Body)
			if nil != err {
				return
			}
			p7ctx.P7request.Body = io.NopCloser(bytes.NewBuffer(p7ctx.ReqBody))

			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("ReqBodyMiddleware Out\r\n")...)
		}
	}
}
