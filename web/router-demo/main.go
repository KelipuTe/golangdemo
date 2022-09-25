package main

import (
	web "demo-golang/web/router"
	"fmt"
)

func main() {
	f4handler := func(p7ctx *web.HTTPContext) {
		routingInfo := p7ctx.GetRoutingInfo()
		requestPath := fmt.Sprintf("request path:%s;", p7ctx.P7request.URL.Path)
		pathParam := "pathParam:"
		for key, val := range p7ctx.M3pathParam {
			pathParam += fmt.Sprintf("%s=%s;", key, val)
		}
		_, _ = p7ctx.I9writer.Write([]byte(routingInfo + requestPath + pathParam))
	}

	p7h := web.NewHTTPHandler()
	p7h.Get("/", f4handler)

	p7h.Get("/hello", f4handler)
	p7h.Get("/hello/world", f4handler)
	p7h.Get("/hello/*", f4handler)

	p7h.Get("/order", f4handler)
	p7h.Get("/order/list/:size/:page", f4handler)
	p7h.Get("/order/:id/detail", f4handler)
	p7h.Post("/order/create", f4handler)
	p7h.Post("/order/:id/delete", f4handler)

	p7s := web.NewHTTPService("9510", p7h)
	err := p7s.Start("127.0.0.1:9510")
	fmt.Println(err)
}
