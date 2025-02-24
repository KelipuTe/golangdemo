package shutdown

import (
	"context"
	"log"
	"time"
)

// 上报统计数据
func F8CountShutdownCallback(ctx context.Context) {
	c7signal := make(chan struct{}, 1)
	go func() {
		log.Println("上报统计数据。。。")
		time.Sleep(1 * time.Second)
		c7signal <- struct{}{}
	}()
	select {
	case <-c7signal:
		log.Println("上报统计数据完成。")
	case <-ctx.Done():
		log.Println("上报统计数据超时。")
	}
}
