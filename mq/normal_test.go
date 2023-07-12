package mq

import (
	"context"
	"demo-golang/mq/mock"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	TestMQName  string = "normalmq"
	TestMSGData string = "msgdata"
)

// 单元测试，用 mockgen

func Test_RedisMock_NormalMQ_Sending(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	caseSlice := []struct {
		name       string
		clientMock func() redis.Cmdable
		mqName     string
		msg        *MessageModel
		wantResult any
		wantErr    string
	}{
		{
			name: "normalmq_sending_one",
			clientMock: func() redis.Cmdable {
				cmdable := mock.NewMockCmdable(ctrl)
				msg := &MessageModel{Data: TestMSGData}
				data, _ := json.Marshal(msg)
				cmd := redis.NewIntCmd(context.Background(), nil)
				cmd.SetVal(int64(1))
				cmdable.EXPECT().RPush(gomock.Any(), TestMQName, string(data)).Return(cmd)
				return cmdable
			},
			mqName: TestMQName,
			msg:    &MessageModel{Data: TestMSGData},
		},
	}

	for _, value := range caseSlice {
		t.Run(value.name, func(t *testing.T) {
			clientMock := value.clientMock()
			storage := NewRedisStorage(clientMock)
			mq := NewNormalMQ(value.mqName, storage)
			err := mq.Sending(context.Background(), value.msg)

			if "" != value.wantErr {
				assert.EqualError(t, err, value.wantErr)
				return
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_RedisMock_NormalMQ_Receiving(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	caseSlice := []struct {
		name       string
		clientMock func() redis.Cmdable
		mqName     string
		msg        *MessageModel
		wantResult any
		wantErr    string
	}{
		{
			name: "normalmq_receiving_one",
			clientMock: func() redis.Cmdable {
				cmdable := mock.NewMockCmdable(ctrl)
				msg := &MessageModel{Data: TestMSGData}
				data, _ := json.Marshal(msg)
				cmd := redis.NewStringCmd(context.Background(), nil)
				cmd.SetVal(string(data))
				cmdable.EXPECT().LPop(gomock.Any(), TestMQName).Return(cmd)
				return cmdable
			},
			mqName:     TestMQName,
			msg:        &MessageModel{Data: TestMSGData},
			wantResult: &MessageModel{Data: TestMSGData},
		},
	}

	for _, value := range caseSlice {
		t.Run(value.name, func(t *testing.T) {
			clientMock := value.clientMock()
			storage := NewRedisStorage(clientMock)
			mq := NewNormalMQ(value.mqName, storage)
			msg, err := mq.Receiving(context.Background())

			if "" != value.wantErr {
				assert.EqualError(t, err, value.wantErr)
				return
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, msg, value.wantResult)
		})
	}
}

// 集成测试，需要 redis

func Test_Redis_NormalMQ_SendingOne(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	msgsend := &MessageModel{Data: TestMSGData}
	err := mq.Sending(context.Background(), msgsend)
	fmt.Println(err)
}

func Test_Redis_NormalMQ_SendingMany(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	for i := 0; i < 10; i++ {
		msgsend := &MessageModel{Data: TestMSGData}
		err := mq.Sending(context.Background(), msgsend)
		fmt.Println(err)
	}
}

func Test_Redis_NormalMQ_ReceivingOne(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	msgreceive, err := mq.Receiving(context.Background())
	fmt.Println(msgreceive, err)
}

func Test_Redis_NormalMQ_ReceivingMany(t *testing.T) {
	config := NewRedisConfig(".env", "./")
	client := NewRedisClient(config)
	storage := NewRedisStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	for i := 0; i < 10; i++ {
		msgreceive, err := mq.Receiving(context.Background())
		fmt.Println(msgreceive, err)
	}
}

// 集成测试，需要 MySQL

func Test_MySQL_NormalMQ_SendingOne(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	msgsend := NewMessageModel(TestMSGData)
	err := mq.Sending(context.Background(), msgsend)
	fmt.Println(err)
}

func Test_MySQL_NormalMQ_ReceivingOne(t *testing.T) {
	config := NewMySQLConfig(".env", "./")
	client := NewMySQLClient(config)
	storage := NewMySQLStorage(client)
	mq := NewNormalMQ(TestMQName, storage)
	msgreceive, err := mq.Receiving(context.Background())
	fmt.Println(msgreceive, err)
}
