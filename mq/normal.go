package mq

import (
	"context"
)

// 普通模式的 MQ
type NormalMQ struct {
	mqName  string
	mqType  int
	storage Storage
}

func (this *NormalMQ) Sending(ctx context.Context, msg *MessageModel) error {
	return this.storage.RightPushToQueue(ctx, this.mqName, msg)
}

func (this *NormalMQ) Receiving(ctx context.Context) (*MessageModel, error) {
	return this.storage.LeftPopFromQueue(ctx, this.mqName)
}

func NewNormalMQ(mqName string, storage Storage) *NormalMQ {
	return &NormalMQ{
		mqName:  mqName,
		mqType:  MQTypeNormal,
		storage: storage,
	}
}
