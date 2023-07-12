package mq

import (
	"strconv"
	"time"
)

// 消息
type MessageModel struct {
	ID             string    `json:"id" gorm:"column:id" gorm:"primaryKey"`
	Data           string    `json:"data" gorm:"column:data"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime:false"`
	SubscriberName string    `json:"-" gorm:"column:subscriber" gorm:"index"`
	Error          error     `json:"-" gorm:"-"`
}

func (this MessageModel) Check() error {
	if "" == this.ID {
		return ErrMessageIDIsEmpty
	}
	return nil
}

func NewMessageModel(data string) *MessageModel {
	nowTime := time.Now()
	nowTimeMNanoStr := strconv.FormatInt(nowTime.UnixNano(), 10)
	return &MessageModel{
		ID:        nowTimeMNanoStr,
		Data:      data,
		CreatedAt: nowTime,
	}
}

// 订阅者
type SubscriberModel struct {
	Name   string    `json:"name" gorm:"column:subscriber" gorm:"primaryKey"`
	SignAt time.Time `json:"sign_at" gorm:"column:sign_at"`
}
