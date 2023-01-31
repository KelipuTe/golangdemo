package queue

import (
	"context"
	"sync"
)

// 类似 sync.Cond 的东西。实现了可以控制超时的 Wait()。
// 可以实现 Broadcast()，但是实现不了 Signal()。
// s6WaitCondV2 和 s6WaitCond 的思路是一样的，就是代码上有点区别
type s6WaitCondV2 struct {
	// 锁，和 Cond 一样，需要传一个锁进来
	i9Locker sync.Locker
	// 直接用 channel 好像没什么问题，用于实现可以控制超时的 Wait()
	// s6WaitCond 这里是指向 channel 的指针
	c7Notify chan struct{}
}

func F8News6WaitCondV2(i9l sync.Locker) *s6WaitCondV2 {
	return &s6WaitCondV2{
		i9Locker: i9l,
		c7Notify: make(chan struct{}),
	}
}

// 获取 s6WaitCondV2.p7c7Notify。一定是加锁之后调用。
func (p7this *s6WaitCondV2) f8GetNotify() <-chan struct{} {
	c7n := p7this.c7Notify
	// s6WaitCond 这里的解锁步骤是由外部控制的。
	// 但是基本上，调用 f8GetNotify 之后紧跟着的就是解锁，所以这里放到里面来
	p7this.i9Locker.Unlock()
	return c7n
}

// 这个就和 Wait() 差不多。一定是加锁之后调用。
func (p7this *s6WaitCondV2) f8Wait() {
	p7c7n := p7this.f8GetNotify()
	<-p7c7n
	p7this.i9Locker.Lock()
}

// 这个是可以控制超时的 Wait()。一定是加锁之后调用。
func (p7this *s6WaitCondV2) f8WaitWithTimeout(i9ctx context.Context) error {
	c7n := p7this.f8GetNotify()
	select {
	case <-c7n:
		p7this.i9Locker.Lock()
	case <-i9ctx.Done():
		return i9ctx.Err()
	}
	return nil
}

// 这个就和 Broadcast() 差不多。一定是加锁之后调用。
func (p7this *s6WaitCondV2) f8Broadcast() {
	c7Old := p7this.c7Notify
	c7New := make(chan struct{})
	p7this.c7Notify = c7New
	// s6WaitCond 这里的解锁步骤是由外部控制的
	// 但是基本上，调用 f8Broadcast 之后紧跟着的就是解锁，所以这里放到里面来
	p7this.i9Locker.Unlock()
	close(c7Old)
}
