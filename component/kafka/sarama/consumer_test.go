package sarama

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"log"
	"testing"
	"time"
)

//consumer消费者demo

//可以使用工具启动一个消费者，方便调试
//go get github.com/IBM/sarama/tools/kafka-console-consumer
//go install github.com/IBM/sarama/tools/kafka-console-consumer@latest
//kafka-console-consumer -topic=test_topic -brokers=localhost:9094

func TestConsumer(t *testing.T) {
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup(addrs, "test_group", cfg)
	require.NoError(t, err)

	//这里会阻塞住，直到消费者退出
	err = consumer.Consume(context.Background(), []string{"test_topic"}, testConsumerGroupHandler{})
	require.NoError(t, err)

	log.Println("消费者退出")
}

type testConsumerGroupHandler struct{}

func (t testConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("Setup")
	//用代码可以重新设置消费的位置，但是一般建议手动去Kafka上操作。
	//因为写在代码里的话，每次服务重新启动，这里都会跑一遍，这是不对的。

	//拿到所有topic以及partition
	topics := session.Claims()
	partitions := topics["test_topic"]
	//重新设置每个partition的消费位置
	for _, partition := range partitions {
		session.ResetOffset("test_topic", partition, sarama.OffsetOldest, "")
	}

	return nil
}

func (t testConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("Cleanup")
	return nil
}

// session就是消费者和Kafka的会话
func (t testConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//t.SyncConsume(session, claim)
	t.AsyncConsume(session, claim)
	return nil
}

// 同步消费
func (t testConsumerGroupHandler) SyncConsume(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) {
	msgs := claim.Messages()
	claim.HighWaterMarkOffset()
	for msg := range msgs {
		log.Println("消费消息", string(msg.Value))
		session.MarkMessage(msg, "")
	}
	//msgs被关了（消费者被关闭）才会走到这里
	//要不然应该一直在上面那个循环里面消费消息
}

// 异步消费需要限制goroutine的数量。
// 一种思路是用协程池的思路，直接限制goroutine的数量。
// 一种思路是批量处理，等够一批，然后开goroutine处理。
// 需要注意，业务低峰，消息数量可能很少，不能无限制的等。
func (t testConsumerGroupHandler) AsyncConsume(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) {
	msgs := claim.Messages()
	claim.HighWaterMarkOffset()

	const batchNum = 10

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		var eg errgroup.Group
		var last *sarama.ConsumerMessage
		for i := 0; i < batchNum; i++ {
			done := false
			select {
			case <-ctx.Done():
				//走到这说明，一秒内没有凑到10条，也需要开始处理了
				done = true
			case msg, ok := <-msgs:
				if !ok {
					//消费者被关闭
					cancel()
					return
				}
				last = msg
				eg.Go(func() error {
					log.Println("消费消息", string(msg.Value))
					//不在这里提交，防止中间有消息处理失败
					//session.MarkMessage(msg, "")
					return nil
				})
			}
			if done {
				break
			}
		}
		cancel()
		err := eg.Wait()
		if err != nil {
			//走到这说明，有消息处理失败了
			continue
		}

		//全部处理完成之后，提交最后一个就行
		session.MarkMessage(last, "")
	}
}
