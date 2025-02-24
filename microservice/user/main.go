package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Mysql struct {
	Id int `gorm:"column:id;primaryKey"`
}

func (Mysql) TableName() string {
	return "qq"
}

func main() {
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	dsn := "root:root@tcp(golangdemomicroserviceusermysql:23306)/qqq?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	server.GET("/mysql", func(ctx *gin.Context) {
		m := &Mysql{}
		db.First(m)
		ctx.JSON(http.StatusOK, m.Id)
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "golangdemomicroserviceredis:26379",
		Password: "",
		DB:       0,
	})

	server.GET("/redis", func(ctx *gin.Context) {

		ctx2 := context.Background()
		key := "mykey"

		val, _ := rdb.Get(ctx2, key).Result()

		ctx.JSON(http.StatusOK, val)
	})

	server.Run(":8080")
}
