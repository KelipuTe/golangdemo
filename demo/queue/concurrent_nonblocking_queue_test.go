package queue

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

// 正常逻辑
func TestS6ConcurrentNonBlockingQueueV2(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentNonBlockingQueueV2[int]()

	enqueueNum := 5
	dequeueNum := 5
	testTime := 5

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)

	for i := 1; i <= enqueueNum; i++ {
		go func(ii int) {
			for j := 0; j < testTime; j++ {
				data := ii*10 + j
				i9ctx := context.Background()
				err := p7queue.F8Enqueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
			}
			wg.Done()
		}(i)
	}

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				t4, err := p7queue.F8Dequeue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// 入队超时
func TestS6ConcurrentNonBlockingQueueV2EnqueueTimeout(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentNonBlockingQueueV2[int]()

	enqueueNum := 1
	dequeueNum := 0
	testTime := 1

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)
	// 这里理论上队列是无限长度的，应该不会超时，所以强行延迟一秒
	for i := 1; i <= enqueueNum; i++ {
		go func(ii int) {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				time.Sleep(1 * time.Second)
				data := ii*10 + j
				err := p7queue.F8Enqueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
				f8cancel()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 出队超时
func TestS6ConcurrentNonBlockingQueueV2DequeueTimeout(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentNonBlockingQueueV2[int]()

	enqueueNum := 0
	dequeueNum := 1
	testTime := 1

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)
	// 这里理论上队列为空，应该不会超时，所以强行延迟一秒
	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				time.Sleep(1 * time.Second)
				t4, err := p7queue.F8Dequeue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
				f8cancel()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 入队少于出队
func TestS6ConcurrentNonBlockingQueueV2EnqueueLessThanDequeue(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentNonBlockingQueueV2[int]()

	enqueueNum := 5
	dequeueNum := 10
	testTime := 5

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)

	for i := 1; i <= enqueueNum; i++ {
		go func(ii int) {
			for j := 0; j < testTime; j++ {
				data := ii*10 + j
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				err := p7queue.F8Enqueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
				f8cancel()
			}
			wg.Done()
		}(i)
	}

	// 这里让入队先跑，然后在放出队跑，要不然出队还是有可能先于入队跑完
	time.Sleep(1 * time.Second)

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				t4, err := p7queue.F8Dequeue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
				f8cancel()
			}
			wg.Done()
		}()
	}

	// 出队 gorouting 比入队 gorouting 的数量多，最后因为没有数据入队会疯狂超时

	wg.Wait()
}
