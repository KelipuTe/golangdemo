package middleware

func RecoveryMiddleware() HTTPMiddleware {
	return func(next HTTPHandleFunc) HTTPHandleFunc {
		return func(p7ctx *HTTPContext) {
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
