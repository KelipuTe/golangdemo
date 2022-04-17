package main

import (
	"fmt"
)

func waitData(chan1 chan int) {
	for { // 无限循环等待数据
		data := <-chan1 // 从channel中获取一个数据
		if data == 0 {  // 将0视为数据结束
			break
		}
		fmt.Println(data) // 打印数据
	}
	chan1 <- 0 // 通知main已经结束循环
}

func main() {
	chan1 := make(chan int)
	go waitData(chan1) // 并发执行waitData, 传入channel
	for i := 1; i < 100; i++ {
		if i%20 == 0 {
			chan1 <- i // 将数据通过channel传给
		}
	}
	chan1 <- 0 // 通知并发的waitData结束
	<-chan1    // 等待waitData结束
	fmt.Println("main done")
}
