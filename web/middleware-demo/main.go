package main

import (
	web "demo-golang/web/middleware"
	"demo-golang/web/middleware/middleware"
	"fmt"
)

func main() {
	p7h := web.NewHTTPHandler()

	p7h.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)

	p7s := web.NewHTTPService("9510", p7h)
	err := p7s.Start("127.0.0.1:9510")
	fmt.Println(err)
}
