package mq

const (
	MQTypeNormal  = 1 + iota // 普通模式
	MQTypePubSub             // 发布订阅模式
	MQTypeConfirm            // 确认模式
)
