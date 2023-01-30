package queue

import "errors"

var (
	ErrQueueIsFull  = errors.New("queue is full")
	ErrQueueIsEmpty = errors.New("queue is empty")
)
