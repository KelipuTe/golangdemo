package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Option 设计模式
type F8ServiceManagerOption func(*S6ServiceManager)

type S6ServiceManager struct {
	s5p7s6HTTPService []*S6HTTPService // 管理的服务列表
	shutdownTimeOut   time.Duration    // 服务关闭的总超时时间，超时强制关闭
	// 服务关闭时，等待正在处理的请求的时间
	// 等待结束后，开始执行服务关闭时需要执行的回调方法
	shutdownWaitTime        time.Duration
	shutdownCallbackTimeOut time.Duration // 服务关闭时需要执行的回调方法的超时时间
}

func NewServiceManager(s5p7s6hs []*S6HTTPService, s5f8option ...F8ServiceManagerOption) *S6ServiceManager {
	p7sm := &S6ServiceManager{
		s5p7s6HTTPService:       s5p7s6hs,
		shutdownTimeOut:         10 * time.Second,
		shutdownWaitTime:        3 * time.Second,
		shutdownCallbackTimeOut: 3 * time.Second,
	}

	//依次执行 Option
	for _, item := range s5f8option {
		item(p7sm)
	}

	return p7sm
}

func (p7this *S6ServiceManager) F8Start() {
	//启动服务
	log.Println("服务启动中。。。")
	for _, item := range p7this.s5p7s6HTTPService {
		t4p7s6hs := item
		go func() {
			if err := t4p7s6hs.F8Start(); nil != err {
				if http.ErrServerClosed == err {
					log.Printf("子服务 %s 已关闭\n", t4p7s6hs.name)
				} else {
					log.Printf("子服务 %s 异常退出，err:%s\r\n", t4p7s6hs.name, err)
				}
			}
		}()
	}
	log.Println("服务启动完成。")

	//监听 ctrl+c 信号
	c7signal := make(chan os.Signal, 2)
	signal.Notify(c7signal, os.Interrupt)
	select {
	case <-c7signal:
		log.Printf("接收到关闭信号，开始关闭服务，限制 %d 秒内完成。。。\r\n", p7this.shutdownTimeOut/time.Second)
		//再次监听 ctrl+c 信号
		go func() {
			select {
			case <-c7signal:
				log.Println("再次接收到关闭信号，服务直接退出。")
				os.Exit(1)
			}
		}()
		time.AfterFunc(p7this.shutdownTimeOut, func() {
			log.Println("优雅关闭超时，服务直接退出。")
			os.Exit(1)
		})
		p7this.F8Shutdown()
	}
}

func (p7this *S6ServiceManager) F8Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), p7this.shutdownTimeOut)
	defer cancel()

	log.Println("停止接收新请求。")
	for _, item := range p7this.s5p7s6HTTPService {
		item.F8Stop()
	}

	log.Printf("等待正在执行的请求结束，等待 %d 秒。。。", p7this.shutdownWaitTime/time.Second)
	time.Sleep(p7this.shutdownWaitTime)

	log.Println("开始关闭子服务。。。")
	wg := sync.WaitGroup{}
	for _, item := range p7this.s5p7s6HTTPService {
		log.Printf("关闭子服务 %s 。。。", item.name)
		t4p7s := item
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = t4p7s.F8ShutDown(ctx)
		}()
	}
	wg.Wait()

	log.Println("开始执行子服务的关闭回调。。。")
	for _, item := range p7this.s5p7s6HTTPService {
		log.Printf("执行子服务 %s 的关闭回调，限制 %d 秒内完成。。。", item.name, p7this.shutdownCallbackTimeOut/time.Second)
		for _, item2 := range item.s5f8ShutdownCallback {
			t4f8cb := item2
			wg.Add(1)
			go func() {
				defer wg.Done()
				t4ctx, t4cancel := context.WithTimeout(context.Background(), p7this.shutdownCallbackTimeOut)
				defer t4cancel()
				t4f8cb(t4ctx)
			}()
		}
	}
	wg.Wait()

	log.Println("服务关闭完成。")
	os.Exit(0)
}

func F8SetShutdownTimeOutOption(t time.Duration) F8ServiceManagerOption {
	return func(p7sm *S6ServiceManager) {
		p7sm.shutdownTimeOut = t
	}
}

func F8SetShutdownWaitTime(t time.Duration) F8ServiceManagerOption {
	return func(p7sm *S6ServiceManager) {
		p7sm.shutdownWaitTime = t
	}
}

func F8SetShutdownCallbackTimeOut(t time.Duration) F8ServiceManagerOption {
	return func(p7sm *S6ServiceManager) {
		p7sm.shutdownCallbackTimeOut = t
	}
}
