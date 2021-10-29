package main

import (
  "fmt"
  "time"

  "github.com/go-redis/redis"
)

func main() {
  redisConf := redis.Options{
    Addr:     "127.0.0.1:6379", //连接路由
    Password: "root",           // 密码，没有密码就填空字符串
    DB:       0,                // 连接哪个库
  } //redis配置
  pRedisClient := redis.NewClient(&redisConf) //创建客户端
  defer pRedisClient.Close()                  //关闭

  result, err := pRedisClient.Ping().Result() //ping测试
  checkErr(err)
  fmt.Println(result) //成功会返回pong

  // KVTest(pRedisClient)
  ListTest(pRedisClient)
}

//处理error
func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

//key value
func KVTest(pRedisClient *redis.Client) {
  testKey := "test_int"
  testInt := 1
  expiration := time.Duration(0)

  //插入
  err := pRedisClient.Set(testKey, testInt, expiration).Err()
  checkErr(err)
  fmt.Printf("set %s=%v\r\n", testKey, testInt)

  //读取
  result, err := pRedisClient.Get(testKey).Result()
  checkErr(err)
  fmt.Printf("get %s=%v\r\n", testKey, result)
}

//list
func ListTest(pRedisClient *redis.Client) {
  var err error
  testKey := "test_list"
  //插入元素
  err = pRedisClient.RPush(testKey, 1, 2, 3).Err()
  checkErr(err)
  //用逻辑下标设置元素，下标从0开始
  index := int64(1)
  err = pRedisClient.LSet(testKey, index, 4).Err()
  checkErr(err)

  //长度
  listLen, err := pRedisClient.LLen(testKey).Result()
  fmt.Printf("%s len=%d\r\n", testKey, listLen)
  //读取
  arr1Result, err := pRedisClient.LRange(testKey, 0, listLen-1).Result()
  checkErr(err)
  fmt.Printf("%s=%v\r\n", testKey, arr1Result)

  //删除
  result, err := pRedisClient.Del(testKey).Result()
  checkErr(err)
  fmt.Printf("lrem %v\r\n", result)

  listLen, err = pRedisClient.LLen(testKey).Result()
  fmt.Printf("%s len=%d\r\n", testKey, listLen)
  arr1Result, err = pRedisClient.LRange(testKey, 0, listLen-1).Result()
  checkErr(err)
  fmt.Printf("%s=%v\r\n", testKey, arr1Result)

  //阻塞式读取
  result2, err := pRedisClient.BLPop(5*time.Second, "test_list").Result()
  checkErr(err)
  fmt.Printf("blpop=%v\r\n", result2)
}
