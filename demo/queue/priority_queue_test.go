package queue

import (
	"log"
	"testing"
	"time"
)

func TestS6PriorityQueueEnqueue(p7s6t *testing.T) {
	p7queue := F8NewS6PriorityQueue[s6DelayQueueNode](100, f8GetS6DelayQueueNodeCompare())

	testTime := 10
	for j := 0; j < testTime; j++ {
		err := p7queue.F8Enqueue(s6DelayQueueNode{
			execTime: time.Now().Add(time.Duration(10*j) * time.Second),
		})
		log.Printf("enqueue:%d, err=%v\n", j, err)
	}

	for _, data := range p7queue.s5data {
		log.Printf("%v\n", data.execTime)
	}
}

func TestS6PriorityQueueDequeue(p7s6t *testing.T) {
	p7queue := F8NewS6PriorityQueue[s6DelayQueueNode](100, f8GetS6DelayQueueNodeCompare())

	testTime := 10
	for j := 0; j < testTime; j++ {
		execTime := time.Now().Add(time.Duration(j) * time.Second)
		err := p7queue.F8Enqueue(s6DelayQueueNode{
			execTime: execTime,
		})
		log.Printf("enqueue:%v, err=%v\n", execTime, err)
	}

	for _, t4value := range p7queue.s5data {
		log.Printf("%v\n", t4value.execTime)
	}

	for j := 0; j < testTime; j++ {
		data, err := p7queue.F8Dequeue()
		log.Printf("dequeue:%v, err=%v\n", data.execTime, err)
	}

	for _, t4value := range p7queue.s5data {
		log.Printf("%v\n", t4value.execTime)
	}
}
