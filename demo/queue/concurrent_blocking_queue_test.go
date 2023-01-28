package queue

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

// 正常逻辑
func TestS6ConcurrentBlockingQueue(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentBlockingQueue[int](10)

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
				err := p7queue.F8EnQueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
			}
			wg.Done()
		}(i)
	}

	// 这里理论上队列会满，然后入队 gorouting 应该停止入队，卡在那里
	time.Sleep(2 * time.Second)

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				t4, err := p7queue.F8DeQueue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// 入队超时
func TestS6ConcurrentBlockingQueueEnqueueTimeout(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentBlockingQueue[int](10)

	enqueueNum := 5
	dequeueNum := 0
	testTime := 5

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)

	for i := 1; i <= enqueueNum; i++ {
		go func(ii int) {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				data := ii*10 + j
				err := p7queue.F8EnQueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
				f8cancel()
			}
			wg.Done()
		}(i)
	}

	// 这里理论上队列会满，然后入队 gorouting 应该停止入队，最终会触发超时。
	time.Sleep(5 * time.Second)

	wg.Wait()
}

// 出队超时
func TestS6ConcurrentBlockingQueueDequeueTimeout(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentBlockingQueue[int](10)

	enqueueNum := 0
	dequeueNum := 5
	testTime := 5

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				t4, err := p7queue.F8DeQueue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
				f8cancel()
			}
			wg.Done()
		}()
	}

	// 这里理论上队列为空，然后出队 gorouting 应该卡在那里，不停地触发超时。
	time.Sleep(5 * time.Second)

	wg.Wait()
}

// 入队出队都超时
func TestS6ConcurrentBlockingQueueTimeout(p7s6t *testing.T) {
	p7queue := F8NewS6ConcurrentBlockingQueue[int](10)

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
				err := p7queue.F8EnQueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%d, err=%v\n", ii, data, err)
				f8cancel()
			}
			wg.Done()
		}(i)
	}

	// 让入队那边超时，一开始的时候入队会超时
	time.Sleep(1 * time.Second)

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				i9ctx, f8cancel := context.WithTimeout(i9ctx, 500*time.Millisecond)
				t4, err := p7queue.F8DeQueue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%d, err=%v\n", ii, t4, err)
				f8cancel()
			}
			wg.Done()
		}()
	}

	// 出队 gorouting 比入队 gorouting 的数量多，最后因为没有数据入队会疯狂超时

	wg.Wait()
}
