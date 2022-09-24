package main

import (
	"demo-golang/web/middleware"
	"fmt"
)

func main() {
	s := middleware.NewHTTPService()
	s.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)
	err := s.Start("127.0.0.1:9510")
	fmt.Println(err)
}
