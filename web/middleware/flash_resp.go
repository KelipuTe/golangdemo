package middleware

import "log"

// FlashRespMiddleware 写入响应数据
func FlashRespMiddleware() HTTPMiddleware {
	return func(next HTTPHandleFunc) HTTPHandleFunc {
		return func(p7ctx *HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte("FlashRespMiddleware In\r\n")...)
			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("FlashRespMiddleware Out\r\n")...)
			flashResp(p7ctx)
		}
	}
}

func flashResp(p7ctx *HTTPContext) {
	if p7ctx.RespStatusCode > 0 {
		p7ctx.I9writer.WriteHeader(p7ctx.RespStatusCode)
	}
	_, err := p7ctx.I9writer.Write(p7ctx.RespData)
	if err != nil {
		log.Fatalln("flashResp failed", err)
	}
}
