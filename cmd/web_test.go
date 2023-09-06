package cmd

import (
	"demo-golang/web"
	"demo-golang/web/middleware"
	"demo-golang/web/shutdown"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestStart(p7s6t *testing.T) {
	p7s6os := makeOpenService()
	p7s6as := makeAdminService()
	p7s6sm := web.NewServiceManager(
		[]*web.S6HTTPService{p7s6os, p7s6as},
		web.F8SetShutdownTimeOutOption(20*time.Second),
		web.F8SetShutdownWaitTime(10*time.Second),
		web.F8SetShutdownCallbackTimeOut(5*time.Second),
	)
	p7s6sm.F8Start()
}

func makeOpenService() *web.S6HTTPService {
	p7s6hh := web.NewS6HTTPHandler()

	p7s6hh.F8AddMiddleware(
		middleware.F8RecoveryMiddleware(),
		middleware.F8ReqBodyMiddleware(),
		middleware.F8LogMiddleware(),
	)

	f8handler := func(p7s6ctx *web.S6HTTPContext) {
		routingInfo := p7s6ctx.F8GetRoutingInfo()
		pathParam := "pathParam:"
		for key, value := range p7s6ctx.M3PathParam {
			pathParam += fmt.Sprintf("%s=%s;", key, value)
		}
		p7s6ctx.RespData = append(p7s6ctx.RespData, []byte(routingInfo+pathParam)...)
	}

	p7s6hh.F8Get("/", f8handler)

	p7s6hh.F8Get("/hello", f8handler)
	p7s6hh.F8Get("/hello/world", f8handler, middleware.F8TestMiddleware("/hello"), middleware.F8TestMiddleware("/world"))
	p7s6hh.F8Get("/hello/*", f8handler, middleware.F8TestMiddleware("/hello/*"))

	p7s6hh.F8Get("/order", f8handler)
	p7s6hh.F8Get("/order/list/:size/:page", f8handler)
	p7s6hh.F8Get("/order/:id/detail", f8handler)
	p7s6hh.F8Post("/order/create", f8handler)
	p7s6hh.F8Post("/order/:id/delete", f8handler)

	p7s6hs := web.NewS6HTTPService("9510", "127.0.0.1:9510", p7s6hh)

	p7s6hs.F8AddShutdownCallback(
		shutdown.F8CacheShutdownCallback,
		shutdown.F8CountShutdownCallback,
	)

	return p7s6hs
}

func makeAdminService() *web.S6HTTPService {
	p7s6hh := web.NewS6HTTPHandler()

	p7s6hh.F8AddMiddleware(
		middleware.F8RecoveryMiddleware(),
		middleware.F8ReqBodyMiddleware(),
		middleware.F8LogMiddleware(),
	)

	f8handler := func(p7ctx *web.S6HTTPContext) {
		routingInfo := p7ctx.F8GetRoutingInfo()
		pathParam := "pathParam:"
		for key, val := range p7ctx.M3PathParam {
			pathParam += fmt.Sprintf("%s=%s;", key, val)
		}
		p7ctx.RespData = append(p7ctx.RespData, []byte(routingInfo+pathParam)...)
	}

	p7s6hh.F8Group(
		"/admin",
		[]web.F8HTTPMiddlewareFunc{middleware.F8TestMiddleware("admin")},
		[]web.S6RouteData{
			{http.MethodGet, "/", f8handler},
			{http.MethodGet, "/user/list/:size/:page", f8handler},
			{http.MethodGet, "/user/:id/detail", f8handler},
			{http.MethodPost, "/user/create", f8handler},
			{http.MethodPost, "/user/:id/delete", f8handler},
		},
	)

	p7s6hs := web.NewS6HTTPService("9511", "127.0.0.1:9511", p7s6hh)

	return p7s6hs
}
