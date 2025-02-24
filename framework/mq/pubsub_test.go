package mq

import (
	"context"
	"fmt"
	"testing"
)

// 集成测试，需要 redis

func Test_Redis_PubSubMQ_Sending(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewPubSubMQPublisher("pubsubmq", storage)
	msgsend := &MessageModel{Data: "msgdata"}
	err := mq.Sending(context.Background(), msgsend)
	fmt.Println(err)
}

func Test_Redis_PubSubMQ_Receiving(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber01")
	msgreceive, err := mq.Receiving(context.Background())
	fmt.Println(msgreceive, err)
}

func Test_Redis_PubSubMQ_Subscribe(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber02")
	err := mq.Subscribe(context.Background())
	fmt.Println(err)
	err = mq.SubscriberSign(context.Background())
	fmt.Println(err)
}

func Test_Redis_PubSubMQ_Unsubscribe(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber01")
	err := mq.Unsubscribe(context.Background())
	fmt.Println(err)
}

// 集成测试，需要 MySQL

func Test_MySQL_PubSubMQ_Sending(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewPubSubMQPublisher("pubsubmq", storage)
	msgsend := NewMessageModel(TestMSGData)
	err := mq.Sending(context.Background(), msgsend)
	fmt.Println(err)
}

func Test_MySQL_PubSubMQ_Receiving(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber01")
	msgreceive, err := mq.Receiving(context.Background())
	fmt.Println(msgreceive, err)
}

func Test_MySQL_PubSubMQ_Subscribe(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber01")
	err := mq.Subscribe(context.Background())
	fmt.Println(err)
	err = mq.SubscriberSign(context.Background())
	fmt.Println(err)
}

func Test_MySQL_PubSubMQ_Unsubscribe(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewPubSubMQSubscriber("pubsubmq", storage, "subscriber01")
	err := mq.Unsubscribe(context.Background())
	fmt.Println(err)
}
