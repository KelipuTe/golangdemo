package queue

import (
	"context"
	"sync"
)

// 同步阻塞队列
type S6ConcurrentBlockingQueue[T any] struct {
	// 队列数据
	s5Data  []T
	maxSize int
	// 锁，用于保护 s5Data
	p7s6mutex *sync.Mutex
}

func F8NewS6ConcurrentBlockingQueue[T any](maxSize int) *S6ConcurrentBlockingQueue[T] {
	p7s6mutex := &sync.Mutex{}
	return &S6ConcurrentBlockingQueue[T]{
		s5Data:    make([]T, maxSize),
		p7s6mutex: p7s6mutex,
	}
}

func (p7this *S6ConcurrentBlockingQueue[T]) F8EnQueue(i9ctx context.Context, data T) {

}

func (p7this *S6ConcurrentBlockingQueue[T]) F8DeQueue(i9ctx context.Context, data T) {

}
