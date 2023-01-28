package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestGoPrintNumASC(p7s6t *testing.T) {
	// 让 goroutine 依次输出 1~50
	total := 50
	index := 1
	f8PrintNum := func(i int) {
		for {
			if index == i {
				if 1 < i {
					fmt.Printf(",")
				}
				fmt.Printf("%d", i)
				index++
				break
			}
			// 这里需要 sleep 一下。要不然一个 goroutine 一直占着 cpu 跑，别的 goroutine 没机会上。
			time.Sleep(time.Millisecond)
		}
	}
	for i := 1; i < total; i++ {
		go f8PrintNum(i)
	}
	// 最后一个单独开，防止主协程直接退出了
	f8PrintNum(total)
	fmt.Printf("\n")
}
