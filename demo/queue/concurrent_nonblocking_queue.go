package queue

import (
	"context"
	"sync/atomic"
	"unsafe"
)

type s6QueueNodeV2[T any] struct {
	value    T
	p7s6next unsafe.Pointer
}

// S6ConcurrentNonBlockingQueue 并发安全的没有最大容量的队列，链表+CAS
type S6ConcurrentNonBlockingQueueV2[T any] struct {
	// 队头指针
	p7s6head unsafe.Pointer
	// 队尾指针
	p7s6tail unsafe.Pointer
	// 队列当前长度
	nowSize int32
}

func F8NewS6ConcurrentNonBlockingQueueV2[T any]() *S6ConcurrentNonBlockingQueueV2[T] {
	p7s6head := &s6QueueNodeV2[T]{}
	p7s6head2 := unsafe.Pointer(p7s6head)
	return &S6ConcurrentNonBlockingQueueV2[T]{
		p7s6head: p7s6head2,
		p7s6tail: p7s6head2,
	}
}

func (p7this *S6ConcurrentNonBlockingQueueV2[T]) F8Enqueue(i9ctx context.Context, data T) error {
	// 先把新的结点准备好
	p7s6NewNode := &s6QueueNodeV2[T]{
		value: data,
	}
	p7s6NewNode2 := unsafe.Pointer(p7s6NewNode)
	// 然后通过 CAS 操作挂到链表的尾部
	for {
		if nil != i9ctx.Err() {
			return i9ctx.Err()
		}
		// 通过原子操作把队尾拿出来
		p7s6tail2 := atomic.LoadPointer(&p7this.p7s6tail)
		// CAS 操作，如果当前的队尾指针就是上面取到的指针，那么把队尾换成新的结点
		// 这里如果用 SQL 语句表示，就可以翻译成下面这样
		// UPDATE p7this.p7s6tail=p7s6NewNode2 WHERE p7this.p7s6tail=p7s6tail2
		if atomic.CompareAndSwapPointer(&p7this.p7s6tail, p7s6tail2, p7s6NewNode2) {
			// CAS 返回成功，说明队尾没变，可以直接修改
			p7s6tail := (*s6QueueNodeV2[T])(p7s6tail2)
			atomic.StorePointer(&p7s6tail.p7s6next, p7s6NewNode2)
			atomic.AddInt32(&p7this.nowSize, 1)
			return nil
		}
		// CAS 返回失败，说明队尾变了，其他想要入队的，已经抢先入队而且完成了，那就要重头再来
	}
}

func (p7this *S6ConcurrentNonBlockingQueueV2[T]) F8Dequeue(i9ctx context.Context) (T, error) {
	for {
		if nil != i9ctx.Err() {
			var t4t T
			return t4t, i9ctx.Err()
		}
		// 通过原子操作把队头拿出来
		p7s6head2 := atomic.LoadPointer(&p7this.p7s6head)
		p7s6head := (*s6QueueNodeV2[T])(p7s6head2)
		p7s6tail2 := atomic.LoadPointer(&p7this.p7s6tail)
		p7s6tail := (*s6QueueNodeV2[T])(p7s6tail2)
		// 检查队列是否为空
		if p7s6tail == p7s6head {
			var t4t T
			return t4t, ErrQueueIsEmpty
		}
		// 通过原子操作把队头结点的下一个结点
		p7s6next2 := atomic.LoadPointer(&p7s6head.p7s6next)
		// CAS 操作，如果当前的队头指针就是上面取到的指针，那么把队头换成队头结点的下一个结点
		if atomic.CompareAndSwapPointer(&p7this.p7s6head, p7s6head2, p7s6next2) {
			// CAS 返回成功，说明队头没变，可以直接修改
			// 直接把结点上的数据取出来，然后出队就行了
			p7s6next := (*s6QueueNodeV2[T])(p7s6next2)
			return p7s6next.value, nil
		}
		// CAS 返回失败，说明队头变了，其他想要出队的，已经抢先出队而且完成了，那就要重头再来
	}
}
