package mq

import "errors"

var (
	ErrNotPublisher     = errors.New("not publisher")
	ErrNotSubscriber    = errors.New("not subscriber")
	ErrMessageNotMatch  = errors.New("message not match")
	ErrMessageIDIsEmpty = errors.New("message id is empty")
)
