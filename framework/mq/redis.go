package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"runtime"
	"strings"
	"sync"
)

type RedisStorage struct {
	client redis.Cmdable
}

func (this *RedisStorage) RightPushToQueue(ctx context.Context, mqName string, msg *MessageModel) error {
	msgStr, err := json.Marshal(msg)
	if nil != err {
		return err
	}
	this.client.RPush(ctx, mqName, msgStr)
	return nil
}

func (this *RedisStorage) RightPushToQueueV2(ctx context.Context, mqName string, subscriberName string, msg *MessageModel) error {
	key := this.makeSubscriberMQListKey(mqName, subscriberName)
	return this.RightPushToQueue(ctx, key, msg)
}

func (this *RedisStorage) LeftPopFromQueue(ctx context.Context, mqName string) (*MessageModel, error) {
	msgStr, err := this.client.LPop(ctx, mqName).Result()
	if nil != err {
		return nil, err
	}
	msg := &MessageModel{}
	err = json.Unmarshal([]byte(msgStr), msg)
	if nil != err {
		return nil, err
	}
	return msg, nil
}

func (this *RedisStorage) LeftPopCountFromQueue(ctx context.Context, mqName string, count int) ([]*MessageModel, error) {
	msgStrSlice, err := this.client.LPopCount(ctx, mqName, count).Result()
	if nil != err {
		return nil, err
	}
	msgSlice := make([]*MessageModel, 0, len(msgStrSlice))
	for _, value := range msgStrSlice {
		msg := &MessageModel{}
		err = json.Unmarshal([]byte(value), msg)
		if nil != err {
			msg.Error = err
		}
		msgSlice = append(msgSlice, msg)
	}
	return msgSlice, nil
}

func (this *RedisStorage) LeftPopCountFromQueueV2(ctx context.Context, mqName string, subscriberName string, count int) ([]*MessageModel, error) {
	return this.LeftPopCountFromQueue(ctx, this.makeSubscriberMQListKey(mqName, subscriberName), count)
}

func (this *RedisStorage) Subscribe(ctx context.Context, mqName string, subscriber *SubscriberModel) error {
	this.client.SAdd(ctx, this.makeChannelSubscriberKey(mqName), subscriber.Name)
	return nil
}

func (this *RedisStorage) Unsubscribe(ctx context.Context, mqName string, subscriberName string) error {
	this.client.SRem(ctx, this.makeChannelSubscriberKey(mqName), subscriberName)
	return nil
}

func (this *RedisStorage) SubscriberSign(ctx context.Context, mqName string, subscriber *SubscriberModel) error {
	data, err := json.Marshal(subscriber)
	if nil != err {
		return err
	}
	// 这里订阅者的登记信息是没有过期时间的 key，由生产者来主动检测和移除
	this.client.Set(ctx, makeSubscriberSignKey(mqName, subscriber.Name), string(data), -1)
	return nil
}

func (this *RedisStorage) GetAllSubscriber(ctx context.Context, mqName string) ([]*SubscriberModel, error) {
	// 获取注册过的订阅者
	subscriberList, err := this.client.SMembers(ctx, this.makeChannelSubscriberKey(mqName)).Result()
	if nil != err {
		return nil, err
	}
	subscriberMap := make(map[string]*SubscriberModel, len(subscriberList))
	for _, value := range subscriberList {
		subscriberMap[value] = &SubscriberModel{
			Name: value,
		}
	}

	// 获取登记过的订阅者，用前缀来模糊匹配
	keySlice, err := this.MatchKeysByPrefix(ctx, makeSubscriberSignKey(mqName, "*"))
	if nil != err {
		return nil, err
	}
	keySliceLen := len(keySlice)

	subscriberChan := make(chan string, keySliceLen)
	wg := sync.WaitGroup{}
	wg.Add(keySliceLen)
	for i := 0; i < keySliceLen; i++ {
		go func(value2 string) {
			defer wg.Done()
			subscriberStr, err2 := this.client.Get(ctx, value2).Result()
			if nil != err2 {
				return
			}
			subscriberChan <- subscriberStr
		}(keySlice[i])
	}
	wg.Wait()
	close(subscriberChan)

	for {
		value, ok := <-subscriberChan
		if !ok {
			break
		}
		subscriberObj := &SubscriberModel{}
		err2 := json.Unmarshal([]byte(value), subscriberObj)
		if nil != err2 {
			continue
		}
		if value2, ok2 := subscriberMap[subscriberObj.Name]; ok2 {
			// 这里拿出来的是指针，所以，可以这么改
			value2.SignAt = subscriberObj.SignAt
		} else {
			subscriberMap[subscriberObj.Name] = subscriberObj
		}
	}
	subscriberSlice := make([]*SubscriberModel, 0, len(subscriberMap))
	for _, value := range subscriberMap {
		subscriberSlice = append(subscriberSlice, value)
	}

	return subscriberSlice, nil
}

func (this *RedisStorage) GetConfirmMessage(ctx context.Context, mqName string, confirmedBy string) (string, error) {
	key := makeSubscriberConfirmKey(mqName, confirmedBy)
	return this.client.Get(ctx, key).Result()
}

func (this *RedisStorage) DeleteConfirmMessage(ctx context.Context, mqName string, confirmedBy string) {
	key := makeSubscriberConfirmKey(mqName, confirmedBy)
	this.client.Del(ctx, key)

}

func (this *RedisStorage) DeleteSubscriber(ctx context.Context, mqName string, confirmedBy string) error {
	this.client.Del(
		ctx,
		this.makeSubscriberMQListKey(mqName, confirmedBy),
		makeSubscriberSignKey(mqName, confirmedBy),
	)
	return nil
}

func (this *RedisStorage) MatchKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	return this.client.Keys(ctx, prefix).Result()
}

func (this *RedisStorage) ReceivingWithConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error) {
	return this.client.Eval(ctx, script, keySlice, argvSlice...).Result()
}

func (this *RedisStorage) ResetMessageFromConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error) {
	return this.client.Eval(ctx, script, keySlice, argvSlice...).Result()
}

// 队列名 => 订阅者注册的键
func (this *RedisStorage) makeChannelSubscriberKey(mqName string) string {
	return mqName + "_subscriber"
}

// 队列名+订阅者名 => 订阅者消息列表的键
func (this *RedisStorage) makeSubscriberMQListKey(mqName string, subscriberName string) string {
	return mqName + "_mq_" + subscriberName
}

// 队列名+订阅者名 => 订阅者登记的键
func makeSubscriberSignKey(mqName string, subscriberName string) string {
	return mqName + "_sign_" + subscriberName
}

// 订阅者登记的键 => 订阅者名
func cutSubscriberNameFromSubscriberSignKey(key string) string {
	strSlice := strings.Split(key, "_sign_")
	return strSlice[1]
}

// 队列名+订阅者名 => 待确认消息的键
func makeSubscriberConfirmKey(mqName string, subscriberName string) string {
	return mqName + "_confirm_" + subscriberName
}

// 待确认消息的键 => 订阅者名
func cutSubscriberNameFromSubscriberConfirmKey(key string) string {
	strSlice := strings.Split(key, "_confirm_")
	return strSlice[1]
}

func NewRedisClient(config Config) redis.Cmdable {
	host := config.GetString("REDIS_HOST")
	port := config.GetString("REDIS_PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	pass := config.GetString("REDIS_PASSWORD")
	db := config.GetInt("REDIS_DB")

	log.Printf("new redis client with: %s/%d\n", addr, db)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
		PoolSize: runtime.NumCPU() * 8,
	})

	return client
}

func NewRedisStorage(client redis.Cmdable) Storage {
	return &RedisStorage{
		client: client,
	}
}
