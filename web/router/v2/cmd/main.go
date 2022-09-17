package main

import "demo-golang/web/router/v2"

func main() {
	s := v2.NewHTTPService()
	s.Get("/", func(p7ctx *v2.HTTPContext) {
		p7ctx.I9writer.Write([]byte("hello, world"))
	})
	s.Get("/user", func(p7ctx *v2.HTTPContext) {
		p7ctx.I9writer.Write([]byte("hello, user"))
	})

	s.Start("127.0.0.1:9510")
}
