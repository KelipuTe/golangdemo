package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 可以使用 channel 限制 WaitGroup 开启的 gorouting 的数量

	rand.Seed(time.Now().UnixNano())

	tokenNum := 5
	chanToken := make(chan struct{}, tokenNum)
	for i := 0; i < tokenNum; i++ {
		chanToken <- struct{}{}
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		<-chanToken
		log.Printf("index=%2d, get token", i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func(i int) {
				chanToken <- struct{}{}
				log.Printf("index=%2d, reset token", i)
			}(i)
			log.Printf("index=%2d, running...", i)
			cost := rand.Intn(1000)
			time.Sleep(time.Duration(cost) * time.Millisecond)
			log.Printf("index=%2d, done, cost %3d ms", i, cost)
		}(i)
	}
}
