package mq

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config interface {
	Init()
	GetString(key string) string
	GetInt(key string) int
}

func newViperAndLoad(filepath string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(filepath)
	err := config.ReadInConfig()
	if nil != err {
		panic(fmt.Sprintf("newViperAndLoad() with err: %s", err))
	}
	return config
}

type RedisConfig struct {
	path     string
	filename string
	config   *viper.Viper
}

func (this *RedisConfig) Init() {
	filePath := this.path + this.filename
	log.Println("load redis config from: " + filePath)
	this.config = newViperAndLoad(filePath)
}

func (this *RedisConfig) GetString(key string) string {
	return this.config.GetString(key)
}

func (this *RedisConfig) GetInt(key string) int {
	return this.config.GetInt(key)
}

func NewRedisConfig(filename string, path string) *RedisConfig {
	config := &RedisConfig{
		filename: filename,
		path:     path,
	}
	config.Init()
	return config
}

type MySQLConfig struct {
	path     string
	filename string
	config   *viper.Viper
}

func (this *MySQLConfig) Init() {
	filePath := this.path + this.filename
	log.Println("load mysql config from: " + filePath)
	this.config = newViperAndLoad(filePath)
}

func (this *MySQLConfig) GetString(key string) string {
	return this.config.GetString(key)
}

func (this *MySQLConfig) GetInt(key string) int {
	return this.config.GetInt(key)
}

func NewMySQLConfig(filename string, path string) *MySQLConfig {
	config := &MySQLConfig{
		filename: filename,
		path:     path,
	}
	config.Init()
	return config
}
