package web

import (
	"context"
	"log"
	"net/http"
)

// http 服务
type I9HTTPService interface {
	//启动服务
	F8Start() error
	//停止服务
	F8Stop()
	//关闭服务
	F8ShutDown(context.Context) error
	I9ShutdownCallback
}

// http 服务
type S6HTTPService struct {
	name                 string               //服务名
	p7server             *http.Server         //调用 ListenAndServe()
	p7handler            *S6HTTPHandler       //http.Handle 接口的实例
	s5f8ShutdownCallback []F8ShutdownCallback //关闭服务时需要执行的回调方法
}

func NewS6HTTPService(name string, addr string, p7h *S6HTTPHandler) *S6HTTPService {
	return &S6HTTPService{
		name: name,
		p7server: &http.Server{
			Addr:    addr,
			Handler: p7h,
		},
		p7handler: p7h,
	}
}

func (p7this *S6HTTPService) F8Start() error {
	log.Printf("服务 %s 启动，监听 %s 端口。", p7this.name, p7this.p7server.Addr)
	p7this.p7handler.s6Router.f8CacheMiddleware()
	return p7this.p7server.ListenAndServe()
}

func (p7this *S6HTTPService) F8Stop() {
	log.Printf("服务 %s 停止。", p7this.name)
	p7this.p7handler.isRunning = false
}

func (p7this *S6HTTPService) F8ShutDown(ctx context.Context) error {
	log.Printf("服务 %s 关闭。", p7this.name)
	return p7this.p7server.Shutdown(ctx)
}

func (p7this *S6HTTPService) F8AddShutdownCallback(s5f4cb ...F8ShutdownCallback) {
	if nil == p7this.s5f8ShutdownCallback {
		p7this.s5f8ShutdownCallback = make([]F8ShutdownCallback, 0, len(s5f4cb))
	}
	p7this.s5f8ShutdownCallback = append(p7this.s5f8ShutdownCallback, s5f4cb...)
}
