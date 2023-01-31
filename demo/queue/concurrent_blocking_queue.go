package queue

import (
	"context"
	"sync"
)

// S6ConcurrentBlockingQueue 并发安全的同步阻塞队列，环形数组+锁
type S6ConcurrentBlockingQueue[T any] struct {
	// 队列数据
	s5Data []T
	// 零值
	zeroData T
	// 队列当前长度
	nowSize int
	// 队列最大长度
	maxSize int
	// 队头指针
	headIndex int
	// 队尾指针
	tailIndex int
	// 锁，用于保护 s5Data
	p7s6mutex *sync.Mutex
	// 条件变量，用于控制入队
	p7s6NotFullWaitCond *s6WaitCond
	// 条件变量，用于控制出队
	p7s6NotEmptyWaitCond *s6WaitCond
}

func F8NewS6ConcurrentBlockingQueue[T any](maxSize int) *S6ConcurrentBlockingQueue[T] {
	p7s6mutex := &sync.Mutex{}
	return &S6ConcurrentBlockingQueue[T]{
		s5Data:               make([]T, maxSize),
		maxSize:              maxSize,
		p7s6mutex:            p7s6mutex,
		p7s6NotFullWaitCond:  f8NewS6WaitCond(p7s6mutex),
		p7s6NotEmptyWaitCond: f8NewS6WaitCond(p7s6mutex),
	}
}

func (p7this *S6ConcurrentBlockingQueue[T]) f8IsFull() bool {
	return p7this.nowSize == p7this.maxSize
}

func (p7this *S6ConcurrentBlockingQueue[T]) f8IsEmpty() bool {
	return 0 == p7this.nowSize
}

func (p7this *S6ConcurrentBlockingQueue[T]) F8Enqueue(i9ctx context.Context, data T) error {
	if nil != i9ctx.Err() {
		return i9ctx.Err()
	}
	p7this.p7s6mutex.Lock()
	// 这里的判空操作一定是 for 而不是 if。
	// 设计上如果队列已经满了，代码跑到这里应该一直等，直到队列里面有空位。所以判空失败的话，应该是继续循环。
	// 如果是 if，那么无论队列满没满，代码跑完判空逻辑，都会继续往下跑。不能达到"一直等，直到队列里面有空位"的目的。
	for p7this.f8IsFull() {
		// 等待队列未满的信号，或者超时的信号
		err := p7this.p7s6NotFullWaitCond.f8WaitWithTimeout(i9ctx)
		if nil != err {
			return err
		}
	}
	// 把数据放进队列
	p7this.s5Data[p7this.tailIndex] = data
	p7this.nowSize++
	// 因为维护了队列当前长度，所以队尾指针直接 +1 就可以了
	p7this.tailIndex++
	if p7this.tailIndex == p7this.maxSize {
		p7this.tailIndex = 0
	}
	// 这里先广播再解锁，让已经在等着的 gorouting 和新来的 gorouting 有机会一起抢锁。
	// 通知出队的 gorouting，队列非空，可以读取
	p7this.p7s6NotEmptyWaitCond.f8Broadcast()
	p7this.p7s6mutex.Unlock()
	return nil
}

func (p7this *S6ConcurrentBlockingQueue[T]) F8Dequeue(i9ctx context.Context) (T, error) {
	var zeroData T
	if nil != i9ctx.Err() {
		return zeroData, i9ctx.Err()
	}
	p7this.p7s6mutex.Lock()
	for p7this.f8IsEmpty() {
		// 等待队列不空的信号，或者超时的信号
		err := p7this.p7s6NotEmptyWaitCond.f8WaitWithTimeout(i9ctx)
		if nil != err {
			return zeroData, err
		}
	}
	data := p7this.s5Data[p7this.headIndex]
	p7this.s5Data[p7this.headIndex] = p7this.zeroData
	p7this.nowSize--
	p7this.headIndex++
	if p7this.headIndex == p7this.maxSize {
		p7this.headIndex = 0
	}
	// 通知入队的 gorouting，队列非满，可以写入
	p7this.p7s6NotFullWaitCond.f8Broadcast()
	p7this.p7s6mutex.Unlock()
	return data, nil
}
