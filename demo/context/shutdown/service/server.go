package service

import (
	"context"
	"log"
	"net/http"
)

type KNHTTPServer struct {
	name     string
	p1Server *http.Server
	p1Mux    *knHTTPServeMux
}

func NewKNHTTPServer(name string, addr string) *KNHTTPServer {
	p1Mux := &knHTTPServeMux{
		isStop:   false,
		ServeMux: http.NewServeMux(),
	}
	return &KNHTTPServer{
		name: name,
		p1Server: &http.Server{
			Addr:    addr,
			Handler: p1Mux,
		},
		p1Mux: p1Mux,
	}
}

// Handler 注册路由
func (p1this *KNHTTPServer) Handler(p string, h http.Handler) {
	p1this.p1Mux.Handle(p, h)
}

func (p1this *KNHTTPServer) Start() error {
	log.Printf("子服务%s启动", p1this.name)
	return p1this.p1Server.ListenAndServe()
}

func (p1this *KNHTTPServer) Stop() {
	log.Printf("子服务%s停止服务", p1this.name)
	p1this.p1Mux.isStop = true
}

func (p1this *KNHTTPServer) ShutDown(ctx context.Context) error {
	log.Printf("子服务%s关闭", p1this.name)
	return p1this.p1Server.Shutdown(ctx)
}
