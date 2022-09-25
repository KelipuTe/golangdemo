package main

import (
	web "demo-golang/web/middleware"
	"demo-golang/web/middleware/middleware"
	"fmt"
)

func main() {
	s := web.NewHTTPService()
	s.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)
	err := s.Start("127.0.0.1:9510")
	fmt.Println(err)
}
