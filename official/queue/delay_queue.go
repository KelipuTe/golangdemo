package queue

import (
	"context"
	"sync"
	"time"
)

type I9CanDelay interface {
	f8GetExecTime() time.Time
	f8NeedDelay() time.Duration
}

type s6DelayQueueNode struct {
	execTime time.Time
}

// 并发安全的延时队列，优先队列+锁
type S6DelayQueue struct {
	// 优先队列
	p7s6PriorityQueue *s6PriorityQueue[I9CanDelay]
	// 锁，用于保护 p7s6PriorityQueue
	p7s6mutex *sync.Mutex
	// 条件变量，用于控制入队
	p7s6NotFullWaitCond *s6WaitCondV2
	// 条件变量，用于控制出队
	p7s6NotEmptyWaitCond *s6WaitCondV2
}

func f8GetS6DelayQueueNodeCompare() F8DoCompare[I9CanDelay] {
	return func(a I9CanDelay, b I9CanDelay) bool {
		return a.f8GetExecTime().Before(b.f8GetExecTime())
	}
}

func F8NewS6DelayQueue(maxSize int, f8DoCompare F8DoCompare[I9CanDelay]) *S6DelayQueue {
	p7s6mutex := &sync.Mutex{}
	return &S6DelayQueue{
		p7s6PriorityQueue:    F8NewS6PriorityQueue[I9CanDelay](maxSize, f8DoCompare),
		p7s6mutex:            p7s6mutex,
		p7s6NotFullWaitCond:  f8NewS6WaitCondV2(p7s6mutex),
		p7s6NotEmptyWaitCond: f8NewS6WaitCondV2(p7s6mutex),
	}
}

func (p7this *s6DelayQueueNode) f8GetExecTime() time.Time {
	return p7this.execTime
}

func (p7this *s6DelayQueueNode) f8NeedDelay() time.Duration {
	return p7this.execTime.Sub(time.Now())
}

func (p7this *S6DelayQueue) F8Enqueue(i9ctx context.Context, i9cd I9CanDelay) error {
	for {
		if nil != i9ctx.Err() {
			return i9ctx.Err()
		}
		p7this.p7s6mutex.Lock()
		// 这里和并发安全的同步阻塞队列不一样，直接尝试入队
		// 通过尝试入队的结果决定接下来的动作，原理上是差不多的
		err := p7this.p7s6PriorityQueue.f8Enqueue(i9cd)
		switch err {
		case nil:
			p7this.p7s6NotEmptyWaitCond.f8Broadcast()
			return nil
		case ErrQueueIsFull:
			// 队列已满，阻塞自己，等待信号
			c7n := p7this.p7s6NotFullWaitCond.f8GetNotify()
			select {
			case <-c7n:
				// 如果被队列未满的信号拉起来，就重新尝试入队
			case <-i9ctx.Done():
				return err
			}
		default:
			// 意料之外的错误，直接解锁并退出
			p7this.p7s6mutex.Unlock()
			return err
		}
	}
}

func (p7this *S6DelayQueue) F8Dequeue(i9ctx context.Context) (I9CanDelay, error) {
	var p7s6timer *time.Timer
	for {
		if nil != i9ctx.Err() {
			return nil, i9ctx.Err()
		}
		p7this.p7s6mutex.Lock()
		data, err := p7this.p7s6PriorityQueue.f8GetTop()
		switch err {
		case nil:
			// 如果需要延迟的时间已经小于 0 了，证明时间到了这个元素应该出列了
			delay := data.f8NeedDelay()
			if delay <= 0 {
				_, _ = p7this.p7s6PriorityQueue.f8Dequeue()
				p7this.p7s6NotFullWaitCond.f8Broadcast()
				return data, nil
			}
			// 否则的话应该继续延迟或者等待入队信号
			if nil == p7s6timer {
				p7s6timer = time.NewTimer(delay)
			} else {
				p7s6timer.Reset(delay)
			}
			c7n := p7this.p7s6NotEmptyWaitCond.f8GetNotify()
			select {
			case <-i9ctx.Done():
				return nil, i9ctx.Err()
			case <-p7s6timer.C:
				// 如果被延迟结束的信号拉起来，就重新尝试检查
			case <-c7n:
				// 如果被队列不空的信号拉起来，就重新尝试检查
			}
		case ErrQueueIsEmpty:
			// 队列为空，阻塞自己，等待信号
			c7n := p7this.p7s6NotEmptyWaitCond.f8GetNotify()
			select {
			case <-i9ctx.Done():
				return nil, i9ctx.Err()
			case <-c7n:
				// 如果被队列不空的信号拉起来，就重新尝试检查
			}
		default:
			p7this.p7s6mutex.Unlock()
			return nil, err
		}
	}
}
