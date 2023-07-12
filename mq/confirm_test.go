package mq

import (
	"context"
	"fmt"
	"testing"
)

// 集成测试，需要 redis

func Test_Redis_ConfirmMQ_Sending(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewConfirmMQPublisher("confirmmq", storage)
	msgsend := &MessageModel{Data: "msgdata"}
	err := mq.Sending(context.Background(), msgsend)
	fmt.Println(err)
}

func Test_Redis_ConfirmMQ_Receiving(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewConfirmMQSubscriber("confirmmq", storage, "subscriber01")
	msgreceive, err := mq.Receiving(context.Background())
	fmt.Println(msgreceive, err)
}

func Test_Redis_ConfirmMQ_CommitConfirm(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewConfirmMQSubscriber("confirmmq", storage, "subscriber01")
	msgreceive, err := mq.CommitConfirm(context.Background(), &MessageModel{})
	fmt.Println(msgreceive, err)
}

func Test_Redis_ConfirmMQ_SubscriberSign(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewConfirmMQSubscriber("confirmmq", storage, "subscriber01")
	err := mq.SubscriberSign(context.Background())
	fmt.Println(err)
}

func Test_Redis_ConfirmMQ_CheckConfirm(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewConfirmMQPublisher("confirmmq", storage)
	err := mq.CheckConfirm(context.Background())
	fmt.Println(err)
}
