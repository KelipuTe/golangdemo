package queue

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

func TestS6DelayQueue(p7s6t *testing.T) {
	p7queue := F8NewS6DelayQueue(10, f8GetS6DelayQueueNodeCompare())

	enqueueNum := 5
	dequeueNum := 5
	testTime := 5

	var wg sync.WaitGroup
	wg.Add(enqueueNum + dequeueNum)

	for i := 1; i <= enqueueNum; i++ {
		go func(ii int) {
			for j := 0; j < testTime; j++ {
				data := &s6DelayQueueNode{
					execTime: time.Now().Add(time.Duration(ii*10+j) * time.Second),
				}
				i9ctx := context.Background()
				err := p7queue.F8Enqueue(i9ctx, data)
				log.Printf("enqueue gorouting:%d, enqueue:%v, err=%v\n", ii, data.f8GetExecTime(), err)
			}
			wg.Done()
		}(i)
	}

	for i := 1; i <= dequeueNum; i++ {
		ii := i
		go func() {
			for j := 0; j < testTime; j++ {
				i9ctx := context.Background()
				data, err := p7queue.F8Dequeue(i9ctx)
				log.Printf("dequeue gorouting:%d, dequeue:%v, err=%v\n", ii, data.f8GetExecTime(), err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
