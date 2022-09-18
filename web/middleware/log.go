package middleware

func LogMiddleware() HTTPMiddleware {
	return func(next HTTPHandleFunc) HTTPHandleFunc {
		return func(p7ctx *HTTPContext) {
			p7ctx.RespData = append(p7ctx.RespData, []byte("LogMiddleware In\r\n")...)
			next(p7ctx)
			p7ctx.RespData = append(p7ctx.RespData, []byte("LogMiddleware Out\r\n")...)
		}
	}
}
