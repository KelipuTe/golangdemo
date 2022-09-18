package middleware

import (
	"bytes"
	"io"
)

// ReqBodyMiddleware 提取请求参数
func ReqBodyMiddleware() HTTPMiddleware {
	return func(next HTTPHandleFunc) HTTPHandleFunc {
		return func(p7ctx *HTTPContext) {
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
