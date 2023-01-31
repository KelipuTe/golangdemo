package queue

import (
	"sync"
	"time"
)

type I9CanDelay interface {
	f8CanDelay()
}

type s6DelayQueueNode struct {
	execTime time.Time
}

type S6DelayQueue struct {
	// 队列数据
	p7s6PriorityQueue *s6PriorityQueue[I9CanDelay]
	// 锁，用于保护 p7s6PriorityQueue
	p7s6mutex *sync.Mutex
	// 条件变量，用于控制入队
	p7s6NotFullWaitCond *s6WaitCond
	// 条件变量，用于控制出队
	p7s6NotEmptyWaitCond *s6WaitCond
}

func f8GetS6DelayQueueNodeCompare() F8DoCompare[s6DelayQueueNode] {
	return func(a s6DelayQueueNode, b s6DelayQueueNode) bool {
		return a.execTime.Before(b.execTime)
	}
}

func (p7this *s6DelayQueueNode) f8CanDelay() {}
