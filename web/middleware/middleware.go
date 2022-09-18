package middleware

type HTTPMiddleware func(next HTTPHandleFunc) HTTPHandleFunc

type MiddlewareInterface interface {
	AddMiddleware(s5mw ...HTTPMiddleware)
}
