package queue

import (
	"context"
	"sync"
	"sync/atomic"
	"unsafe"
)

// 类似 sync.Cond 的东西。实现了可以控制超时的 Wait()。
// 可以实现 Broadcast()，但是实现不了 Signal()。
type s6WaitCond struct {
	// 锁，和 Cond 一样，需要传一个锁进来
	i9Locker sync.Locker
	// 指向 channel 的指针，发信号用的。用于实现可以控制超时的 Wait()
	p7c7Notify unsafe.Pointer
}

func F8NewS6WaitCond(i9l sync.Locker) *s6WaitCond {
	p7s6wc := &s6WaitCond{i9Locker: i9l}
	p7c7n := make(chan struct{})
	p7s6wc.p7c7Notify = unsafe.Pointer(&p7c7n)
	return p7s6wc
}

// 获取 s6WaitCond.p7c7Notify。一定是加锁之后调用。
func (p7this *s6WaitCond) f8GetNotify() <-chan struct{} {
	// 在调用的时候，外面一定是加锁的，这里的原子操作不知道是干啥的。
	p7c7n := atomic.LoadPointer(&p7this.p7c7Notify)
	return *(*chan struct{})(p7c7n)
}

// 这个就和 Wait() 差不多。一定是加锁之后调用。
func (p7this *s6WaitCond) f8Wait() {
	// 获取一个用于等通知的结构
	p7c7n := p7this.f8GetNotify()
	// 把拿着的锁放掉
	p7this.i9Locker.Unlock()
	// 阻塞，等通知
	<-p7c7n
	// 等到通知了，把锁加回来。
	p7this.i9Locker.Lock()
}

// 这个是可以控制超时的 Wait()。一定是加锁之后调用。
func (p7this *s6WaitCond) f8WaitWithTimeout(i9ctx context.Context) error {
	p7c7n := p7this.f8GetNotify()
	p7this.i9Locker.Unlock()
	select {
	case <-p7c7n:
		// 这里被唤醒说明有 gorouting 调用了 Broadcast() 关闭了 channel。
		// 这个地方并没有完全解决超时的问题，因为这里加锁的逻辑还是有可能被阻塞的。
		p7this.i9Locker.Lock()
	case <-i9ctx.Done():
		// 超时，这里就不用把锁加回来了。外层应该拿到这里的异常，然后执行异常处理逻辑。
		return i9ctx.Err()
	}
	return nil
}

// 这个就和 Broadcast() 差不多。一定是加锁之后调用。
func (p7this *s6WaitCond) f8Broadcast() {
	// 创建一个新的 channel
	p7c7NNew := make(chan struct{})
	// 用新的 channel 替换旧的 channel，注意要用原子操作。
	p7c7NOld := atomic.SwapPointer(&p7this.p7c7Notify, unsafe.Pointer(&p7c7NNew))
	// 关闭旧的 channel，相当于做了一次广播，所有阻塞着的 gorouting 都会读到零值
	close(*(*chan struct{})(p7c7NOld))
}
