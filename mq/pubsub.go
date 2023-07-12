package mq

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

const (
	SignTimeMaxInterval time.Duration = 30 * time.Minute
)

// 发布订阅模式的 MQ
type PubSubMQ struct {
	mqName  string
	mqType  int
	storage Storage
	// subscriberName=="" 是发布者；subscriberName!="" 是订阅者；
	subscriberName string
}

func (this *PubSubMQ) IsPublisher() bool {
	return "" == this.subscriberName
}

func (this *PubSubMQ) IsSubscriber() bool {
	return "" != this.subscriberName
}

func (this *PubSubMQ) Sending(ctx context.Context, msg *MessageModel) error {
	if !this.IsPublisher() {
		return ErrNotPublisher
	}
	// 尝试格式化
	_, err := json.Marshal(msg)
	if nil != err {
		return err
	}
	subscriberSlice, err := this.storage.GetAllSubscriber(ctx, this.mqName)
	if nil != err {
		return err
	}

	subscriberSliceLen := len(subscriberSlice)
	wg := sync.WaitGroup{}
	wg.Add(subscriberSliceLen)
	for i := 0; i < subscriberSliceLen; i++ {
		go func(value2 *SubscriberModel) {
			defer wg.Done()
			if value2.SignAt.Add(SignTimeMaxInterval).After(time.Now()) {
				// 这方法里面就一个格式化会报错，前面已经尝试过了，这里直接忽略错误
				_ = this.storage.RightPushToQueueV2(ctx, this.mqName, value2.Name, msg)
			} else {
				// 如果 30 分钟没有登记，生产者就主动解除订阅者的注册
				_ = this.storage.Unsubscribe(ctx, this.mqName, value2.Name)
				_ = this.storage.DeleteSubscriber(ctx, this.mqName, value2.Name)
			}
		}(subscriberSlice[i])
	}
	wg.Wait()

	return nil
}

func (this *PubSubMQ) Receiving(ctx context.Context) ([]*MessageModel, error) {
	if !this.IsSubscriber() {
		return nil, ErrNotSubscriber
	}
	return this.storage.LeftPopCountFromQueueV2(ctx, this.mqName, this.subscriberName, 10)
}

// 订阅者注册
func (this *PubSubMQ) Subscribe(ctx context.Context) error {
	if !this.IsSubscriber() {
		return ErrNotSubscriber
	}
	subscriber := &SubscriberModel{
		Name:   this.subscriberName,
		SignAt: time.Now(),
	}
	return this.storage.Subscribe(ctx, this.mqName, subscriber)
}

// 订阅者解除注册
func (this *PubSubMQ) Unsubscribe(ctx context.Context) error {
	if !this.IsSubscriber() {
		return ErrNotSubscriber
	}
	return this.storage.Unsubscribe(ctx, this.mqName, this.subscriberName)
}

// 订阅者登记（类似心跳）
func (this *PubSubMQ) SubscriberSign(ctx context.Context) error {
	if !this.IsSubscriber() {
		return ErrNotSubscriber
	}
	subscriber := &SubscriberModel{
		Name:   this.subscriberName,
		SignAt: time.Now(),
	}
	return this.storage.SubscriberSign(ctx, this.mqName, subscriber)
}

func NewPubSubMQPublisher(mqName string, storage Storage) *PubSubMQ {
	return &PubSubMQ{
		mqName:  mqName,
		mqType:  MQTypePubSub,
		storage: storage,
	}
}

func NewPubSubMQSubscriber(mqName string, storage Storage, subscriberName string) *PubSubMQ {
	return &PubSubMQ{
		mqName:         mqName,
		mqType:         MQTypePubSub,
		storage:        storage,
		subscriberName: subscriberName,
	}
}
