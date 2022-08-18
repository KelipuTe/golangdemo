package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// KNServiceOption Option 设计模式
type KNServiceOption func(*KNService)

// ShutdownCallback 关闭服务前，需要执行的回调
type ShutdownCallback func(context.Context)

type KNService struct {
	slip1HTTPServer []*KNHTTPServer
	// 等待 waitTimeout 后，开始执行回调
	waitTimeout time.Duration
	// 回调
	sliCB []ShutdownCallback
	// 回调的执行时间限制，超时需要处理
	cbTimeout time.Duration
	// 优雅关闭的执行时间限制，超时强制关闭
	shutdownTimeout time.Duration
}

func NewKNService(slip1HTTPServer []*KNHTTPServer, sliOpt ...KNServiceOption) *KNService {
	p1s := &KNService{
		slip1HTTPServer: slip1HTTPServer,
		waitTimeout:     5 * time.Second,
		cbTimeout:       3 * time.Second,
		shutdownTimeout: 10 * time.Second,
	}

	// 依次执行 Option
	for _, opt := range sliOpt {
		opt(p1s)
	}

	return p1s
}

func (p1this *KNService) Start() {
	// 启动服务
	log.Println("服务启动中")
	for _, s := range p1this.slip1HTTPServer {
		t1s := s
		go func() {
			if err := t1s.Start(); err != nil {
				if http.ErrServerClosed == err {
					log.Printf("子服务%s已关闭", t1s.name)
				} else {
					log.Printf("子服务%s异常退出", t1s.name)
				}
			}
		}()
	}
	log.Println("服务启动完成")

	// 监听 ctrl+c 信号
	chanSgn := make(chan os.Signal, 1)
	signal.Notify(chanSgn, os.Interrupt)
	select {
	case <-chanSgn:
		log.Println("接收到关闭信号，开始关闭服务")
		go func() {
			select {
			case <-chanSgn:
				log.Println("再次接收到关闭信号，服务直接退出")
				os.Exit(1)
			}
		}()
		time.AfterFunc(p1this.shutdownTimeout, func() {
			log.Println("优雅关闭超时，服务直接退出")
			os.Exit(1)
		})
		p1this.Shutdown()
	}
}

func (p1this *KNService) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), p1this.shutdownTimeout)
	defer cancel()

	log.Println("停止接收新请求")
	for _, p1s := range p1this.slip1HTTPServer {
		p1s.Stop()
	}

	log.Printf("等待正在执行的请求结束，等待%d秒。。。", p1this.waitTimeout/time.Second)
	time.Sleep(p1this.waitTimeout)

	log.Println("开始关闭子服务")
	wg := sync.WaitGroup{}
	for _, p1s := range p1this.slip1HTTPServer {
		t1p1s := p1s
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1p1s.ShutDown(ctx)
		}()
	}
	wg.Wait()

	log.Println("开始执行回调")
	for _, cb := range p1this.sliCB {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1ctx, t1cancel := context.WithTimeout(context.Background(), p1this.cbTimeout)
			defer t1cancel()
			cb(t1ctx)
		}()
	}
	wg.Wait()

	log.Println("服务关闭")
}

func WithShutdownCallback(sliCB ...ShutdownCallback) KNServiceOption {
	return func(p1s *KNService) {
		p1s.sliCB = sliCB
	}
}
