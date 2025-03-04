package sarama

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
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
	//c.Producer.Partitioner 默认是NewHashPartitioner（哈希），想要达到业务有序非常简单
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addrs, cfg)
	require.NoError(t, err)

	msg := &sarama.ProducerMessage{
		//topic
		Topic: "test_topic",
		//消息本体
		Value: sarama.StringEncoder("hello world"),
		//在生产者和消费者之间传递数据，类似http的header
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("traceId"),
				Value: []byte("traceId123456"),
			},
		},
	}
	_, _, err = producer.SendMessage(msg)
	require.NoError(t, err)
}

// 异步发消息
func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(addrs, cfg)
	require.NoError(t, err)

	go func() {
		msgChan := producer.Input()
		msg := &sarama.ProducerMessage{
			Topic: "test_topic",
			Value: sarama.StringEncoder("hello world"),
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("traceId"),
					Value: []byte("traceId123456"),
				},
			},
		}
		msgChan <- msg
	}()

	succChan := producer.Successes()
	errChan := producer.Errors()

	select {
	case <-succChan:
		t.Log("发送成功")
	case err := <-errChan:
		t.Log("发送失败", err.Error())
	}
}
