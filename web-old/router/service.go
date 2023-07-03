package router

import (
	"log"
	"net/http"
)

// HTTPServiceInterface 核心服务的接口定义
type HTTPServiceInterface interface {
	// Start 启动服务
	Start(addr string) error
}

// HTTPService 核心服务
type HTTPService struct {
	// name 服务名
	name string
	// p7server 核心处理逻辑
	p7server *http.Server
}

func NewHTTPService(name string, p7h *HTTPHandler) *HTTPService {
	return &HTTPService{
		name: name,
		p7server: &http.Server{
			Handler: p7h,
		},
	}
}

func (p7this *HTTPService) Start(addr string) error {
	p7this.p7server.Addr = addr
	log.Printf("http 服务 %s 启动，监听 %s 端口。\r\n", p7this.name, addr)
	return p7this.p7server.ListenAndServe()
}
