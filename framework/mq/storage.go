package mq

import "context"

type Storage interface {
	// 向队列尾部插入一条消息
	RightPushToQueue(ctx context.Context, mqName string, msg *MessageModel) error
	RightPushToQueueV2(ctx context.Context, mqName string, subscriberName string, msg *MessageModel) error
	// 从队列头部取出一条消息
	LeftPopFromQueue(ctx context.Context, mqName string) (*MessageModel, error)
	// 从队列头部取出多条消息
	LeftPopCountFromQueue(ctx context.Context, mqName string, count int) ([]*MessageModel, error)
	LeftPopCountFromQueueV2(ctx context.Context, mqName string, subscriberName string, count int) ([]*MessageModel, error)

	// 订阅者订阅
	Subscribe(ctx context.Context, mqName string, subscriber *SubscriberModel) error
	// 订阅者解除订阅
	Unsubscribe(ctx context.Context, mqName string, subscriberName string) error
	// 订阅者登记（类似心跳）
	SubscriberSign(ctx context.Context, mqName string, subscriber *SubscriberModel) error

	// 获取所有的订阅者
	GetAllSubscriber(ctx context.Context, mqName string) ([]*SubscriberModel, error)
	// 移除订阅者的消息队列和订阅信息
	DeleteSubscriber(ctx context.Context, mqName string, subscriberName string) error

	// 从队列头部取出一条消息，并将消息标记为待确认
	ReceivingWithConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error)
	// 移除消息的待确认标记
	DeleteConfirmMessage(ctx context.Context, mqName string, subscriberName string)
	// 获取待确认的消息
	GetConfirmMessage(ctx context.Context, mqName string, subscriberName string) (string, error)

	// 获取所有的待确认消息
	MatchKeysByPrefix(ctx context.Context, prefix string) ([]string, error)
	// 把消息的待确认标记移除，并插入到队列头部
	ResetMessageFromConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error)
}
