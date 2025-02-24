package mq

import (
	"context"
	_ "embed"
	"encoding/json"
	"time"
)

//go:embed lua/receiving_with_confirm.lua
var luaReceivingWithConfirm string

//go:embed lua/reset_message_from_confirm.lua
var luaResetMessageFromConfirm string

// 确认模式的 MQ
type ConfirmMQ struct {
	mqName  string
	mqType  int
	storage Storage
	// subscriberName=="" 是发布者；subscriberName!="" 是订阅者；
	// 订阅者需要执行确认动作
	subscriberName string
}

func (this *ConfirmMQ) IsPublisher() bool {
	return "" == this.subscriberName
}

func (this *ConfirmMQ) IsSubscriber() bool {
	return "" != this.subscriberName
}

func (this *ConfirmMQ) Sending(ctx context.Context, msg *MessageModel) error {
	if !this.IsPublisher() {
		return ErrNotPublisher
	}
	return this.storage.RightPushToQueue(ctx, this.mqName, msg)
}

func (this *ConfirmMQ) Receiving(ctx context.Context) (*MessageModel, error) {
	keySlice, err := this.storage.MatchKeysByPrefix(ctx, makeSubscriberConfirmKey(this.mqName, "*"))
	if nil != err {
		return nil, err
	}

	dataAny, err := this.storage.ReceivingWithConfirm(
		ctx,
		luaReceivingWithConfirm,
		[]string{
			this.mqName,
			makeSubscriberConfirmKey(this.mqName, this.subscriberName),
		},
		keySlice,
	)
	if nil != err {
		return nil, err
	}

	dataStr, _ := dataAny.(string)
	msg := &MessageModel{}
	err = json.Unmarshal([]byte(dataStr), msg)
	if nil != err {
		return nil, err
	}
	return msg, nil
}

func (this *ConfirmMQ) CommitConfirm(ctx context.Context, msg *MessageModel) (*MessageModel, error) {
	existMsgStr, err := this.storage.GetConfirmMessage(ctx, this.mqName, this.subscriberName)
	if nil != err {
		return nil, err
	}
	existMsg := &MessageModel{}
	err = json.Unmarshal([]byte(existMsgStr), existMsg)
	if nil != err {
		return nil, err
	}

	this.storage.DeleteConfirmMessage(ctx, this.mqName, this.subscriberName)

	// 理论上这种情况不应该出现
	if msg.ID != existMsg.ID {
		return existMsg, ErrMessageNotMatch
	}

	return nil, nil
}

func (this *ConfirmMQ) CheckConfirm(ctx context.Context) error {
	signSlice, err := this.storage.MatchKeysByPrefix(ctx, makeSubscriberSignKey(this.mqName, "*"))
	if nil != err {
		return err
	}

	signMap := make(map[string]string, len(signSlice))
	for _, value := range signSlice {
		subscriberName := cutSubscriberNameFromSubscriberSignKey(value)
		signMap[subscriberName] = subscriberName
	}

	confirmSlice, err := this.storage.MatchKeysByPrefix(ctx, makeSubscriberConfirmKey(this.mqName, "*"))
	if nil != err {
		return err
	}

	for _, value := range confirmSlice {
		subscriberName := cutSubscriberNameFromSubscriberConfirmKey(value)
		if _, ok := signMap[subscriberName]; !ok {
			// 如果发布者发现有消息被标记了但是订阅者死了，就需要把这个消息拿回来
			_, err2 := this.storage.ResetMessageFromConfirm(
				ctx,
				luaResetMessageFromConfirm,
				[]string{
					this.mqName,
					value,
				},
			)
			if nil != err2 {
				return err2
			}
		}
	}

	return nil
}

// 订阅者登记（类似心跳）
func (this *ConfirmMQ) SubscriberSign(ctx context.Context) error {
	if !this.IsSubscriber() {
		return ErrNotSubscriber
	}
	subscriber := &SubscriberModel{
		Name:   this.subscriberName,
		SignAt: time.Now(),
	}
	return this.storage.SubscriberSign(ctx, this.mqName, subscriber)
}

func NewConfirmMQPublisher(mqName string, storage Storage) *ConfirmMQ {
	return &ConfirmMQ{
		mqName:         mqName,
		mqType:         MQTypeConfirm,
		storage:        storage,
		subscriberName: "",
	}
}

func NewConfirmMQSubscriber(mqName string, storage Storage, confirmedBy string) *ConfirmMQ {
	return &ConfirmMQ{
		mqName:         mqName,
		mqType:         MQTypeConfirm,
		storage:        storage,
		subscriberName: confirmedBy,
	}
}
