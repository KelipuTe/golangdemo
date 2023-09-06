package shutdown

import (
	"context"
	"log"
	"time"
)

// 持久化缓存内容
func F8CacheShutdownCallback(ctx context.Context) {
	c7signal := make(chan struct{}, 1)
	go func() {
		log.Println("持久化缓存内容中。。。")
		time.Sleep(1 * time.Second)
		c7signal <- struct{}{}
	}()
	select {
	case <-c7signal:
		log.Println("持久化缓存内容成功。")
	case <-ctx.Done():
		log.Println("持久化缓存内容超时。")
	}
}
