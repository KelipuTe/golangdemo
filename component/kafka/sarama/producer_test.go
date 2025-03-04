package sarama

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

//producer生产者demo

//可以使用工具启动一个生产者，方便调试
//go get github.com/IBM/sarama/tools/kafka-console-producer
//go install github.com/IBM/sarama/tools/kafka-console-producer@latest

var addrs = []string{"localhost:9094"}

// 同步发消息
func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	//c.Producer.Partitioner 可以设置Partition逻辑，默认是NewHashPartitioner（哈希）
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addrs, cfg)
	require.NoError(t, err)

	//Topic，topic
	//Value，消息本体
	//Headers，在生产者和消费者之间传递数据，类似http的header
	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("TestSyncProducer"),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("traceId"),
				Value: []byte("TestSyncProducer"),
			},
		},
	}
	_, _, err = producer.SendMessage(msg)
	require.NoError(t, err)

	log.Println("发送成功")
}

// 异步发消息
func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(addrs, cfg)
	require.NoError(t, err)

	go func() {
		msgCh := producer.Input()
		msg := &sarama.ProducerMessage{
			Topic: "test_topic",
			Value: sarama.StringEncoder("TestAsyncProducer"),
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("traceId"),
					Value: []byte("TestAsyncProducer"),
				},
			},
		}
		//如果Kafka那边出现性能瓶颈了，msgCh<-msg 是有可能阻塞的
		select {
		case msgCh <- msg:
			t.Log("生产者不阻塞")
		default:
			t.Log("生产者阻塞")
		}
	}()

	succCh := producer.Successes()
	errCh := producer.Errors()

	select {
	case <-succCh:
		t.Log("发送成功")
	case err := <-errCh:
		t.Log("发送失败", err.Error())
	}
}
