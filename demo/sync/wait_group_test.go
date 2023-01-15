package sync

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupWithLimit(p7s6t *testing.T) {
	// 用 channel 模拟令牌桶
	tokenNum := 3
	c7token := make(chan struct{}, tokenNum)
	for i := 0; i < tokenNum; i++ {
		c7token <- struct{}{}
	}

	// 设置随机数，用于模拟请求耗时
	rand.Seed(time.Now().UnixNano())
	p7s6wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		<-c7token
		log.Printf("index=%2d, get token", i)

		p7s6wg.Add(1)
		go f8DoFunc(p7s6wg, c7token, i)
	}
}

func f8DoFunc(s6wg *sync.WaitGroup, c7token chan struct{}, i int) {
	defer s6wg.Done()
	defer func(i int) {
		c7token <- struct{}{}
		log.Printf("index=%2d, reset token", i)
	}(i)
	log.Printf("index=%2d, running...", i)
	cost := rand.Intn(1000)
	time.Sleep(time.Duration(cost) * time.Millisecond)
	log.Printf("index=%2d, done, cost %3d ms", i, cost)
}
