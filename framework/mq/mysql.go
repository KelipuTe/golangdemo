package mq

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
)
import "gorm.io/driver/mysql"

type MySQLStorage struct {
	db *gorm.DB
}

func (this *MySQLStorage) RightPushToQueue(ctx context.Context, mqName string, msg *MessageModel) error {
	db := this.db.Table(mqName).Select("ID", "Data", "CreatedAt").Create(msg)
	fmt.Println(db.Error)
	return nil
}

func (this *MySQLStorage) RightPushToQueueV2(ctx context.Context, mqName string, subscriberName string, msg *MessageModel) error {
	msg.SubscriberName = subscriberName
	db := this.db.Table(this.makeSubscriberMQListKey(mqName)).Create(msg)
	fmt.Println(db.Error)
	return nil
}

func (this *MySQLStorage) LeftPopFromQueue(ctx context.Context, mqName string) (*MessageModel, error) {
	notEmpty := true
	msg := &MessageModel{}
	err := this.db.Transaction(func(tx *gorm.DB) error {
		tx.Table(mqName).Clauses(clause.Locking{Strength: "UPDATE"}).Order("created_at").First(msg)
		if "" != msg.ID {
			notEmpty = false
			tx.Table(mqName).Where("id = ?", msg.ID).Delete(&MessageModel{})
		}
		return nil
	})
	if nil != err {
		return nil, err
	}
	if notEmpty {
		return msg, nil
	}
	return nil, nil
}

func (this *MySQLStorage) LeftPopCountFromQueue(ctx context.Context, mqName string, count int) ([]*MessageModel, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MySQLStorage) LeftPopCountFromQueueV2(ctx context.Context, mqName string, subscriberName string, count int) ([]*MessageModel, error) {
	notEmpty := true
	var msgSlice []*MessageModel
	err := this.db.Transaction(func(tx *gorm.DB) error {
		tx.Table(this.makeSubscriberMQListKey(mqName)).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Order("created_at").Limit(count).Find(&msgSlice)
		if 0 < len(msgSlice) {
			notEmpty = false
			idSlice := make([]string, 0, len(msgSlice))
			for _, value := range msgSlice {
				idSlice = append(idSlice, value.ID)
			}
			tx.Table(this.makeSubscriberMQListKey(mqName)).
				Where("id in ?", idSlice).Delete(&MessageModel{})
		}
		return nil
	})
	if nil != err {
		return nil, err
	}
	if notEmpty {
		return msgSlice, nil
	}
	return nil, nil
}

func (this *MySQLStorage) Subscribe(ctx context.Context, mqName string, subscriber *SubscriberModel) error {
	this.db.Table(this.makeChannelSubscriberKey(mqName)).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "subscriber"}},
		DoUpdates: clause.AssignmentColumns([]string{"sign_at"}),
	}).Create(subscriber)
	return nil
}

func (this *MySQLStorage) Unsubscribe(ctx context.Context, mqName string, subscriberName string) error {
	this.db.Table(this.makeChannelSubscriberKey(mqName)).
		Where("subscriber = ?", subscriberName).Delete(&SubscriberModel{})
	return nil
}

func (this *MySQLStorage) SubscriberSign(ctx context.Context, mqName string, subscriber *SubscriberModel) error {
	return this.Subscribe(ctx, mqName, subscriber)
}

func (this *MySQLStorage) GetAllSubscriber(ctx context.Context, mqName string) ([]*SubscriberModel, error) {
	var subscriberSlice []*SubscriberModel
	this.db.Table(this.makeChannelSubscriberKey(mqName)).Find(&subscriberSlice)
	return subscriberSlice, nil
}

func (this *MySQLStorage) GetConfirmMessage(ctx context.Context, mqName string, subscriberName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MySQLStorage) DeleteConfirmMessage(ctx context.Context, mqName string, subscriberName string) {
	//TODO implement me
	panic("implement me")
}

func (this *MySQLStorage) DeleteSubscriber(ctx context.Context, mqName string, subscriberName string) error {
	err := this.db.Transaction(func(tx *gorm.DB) error {
		this.db.Table(this.makeSubscriberMQListKey(mqName)).
			Where("subscriber = ?", subscriberName).Delete(&MessageModel{})
		this.db.Table(this.makeChannelSubscriberKey(mqName)).
			Where("subscriber = ?", subscriberName).Delete(&SubscriberModel{})
		return nil
	})
	if nil != err {
		return err
	}
	return nil
}

func (this *MySQLStorage) MatchKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MySQLStorage) ReceivingWithConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MySQLStorage) ResetMessageFromConfirm(ctx context.Context, script string, keySlice []string, argvSlice ...any) (any, error) {
	//TODO implement me
	panic("implement me")
}

// 队列名 => 订阅者注册的键
func (this *MySQLStorage) makeChannelSubscriberKey(mqName string) string {
	return mqName + "_subscriber"
}

// 队列名 => 订阅者消息列表的键
func (this *MySQLStorage) makeSubscriberMQListKey(mqName string) string {
	return mqName + "_mq"
}

func NewMySQLClient(config Config) *gorm.DB {
	host := config.GetString("MYSQL_HOST")
	port := config.GetString("MYSQL_PORT")
	user := config.GetString("MYSQL_USERNAME")
	pass := config.GetString("MYSQL_PASSWORD")
	dbname := config.GetString("MYSQL_DATABASE")
	charset := config.GetString("MYSQL_CHARSET")
	parseTime := config.GetString("MYSQL_PARSE_TIME")
	loc := config.GetString("MYSQL_LOC")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, pass, host, port, dbname, charset, parseTime, loc,
	)

	filePath := filepath.Join(config.GetString("GORM_LOG_FILE"))
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		log.Printf("new gorm logger with err: %s\n", err)
		return nil
	}
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{LogLevel: logger.Info},
	)

	log.Printf("new mysql client with: %s\n", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if nil != err {
		log.Printf("new mysql client with err: %s\n", err)
		return nil
	}

	return db
}

func NewMySQLStorage(db *gorm.DB) *MySQLStorage {
	return &MySQLStorage{
		db: db,
	}
}
