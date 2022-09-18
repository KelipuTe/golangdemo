package main

import (
	"demo-golang/web/middleware"
)

func main() {
	s := middleware.NewHTTPService()
	s.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)
	s.Start("127.0.0.1:9510")
}
