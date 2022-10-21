package viper

import (
  "fmt"
  "log"

  "github.com/spf13/viper"
)

func init() {
  log.Println("viper config init...")

  LoadEnvConfig()

  log.Println("viper config init done.")
}

const VERSION string = "1.0.0"
const CONFIG_PATH string = "config/"  // 配置文件目录
const ENV_CONFIG_NAME string = ".env" // 环境和敏感配置
var p1envConfig *viper.Viper = nil

// 生成配置
func NewViperFromEnv(path string) (*viper.Viper, error) {
  p1viper := viper.New()
  p1viper.SetConfigFile(path)
  err := p1viper.ReadInConfig()
  if nil != err {
    return nil, err
  }
  return p1viper, nil
}

func LoadEnvConfig() {
  var err error = nil
  p1envConfig, err = NewViperFromEnv(CONFIG_PATH + ENV_CONFIG_NAME)
  if nil != err {
    panic(fmt.Sprintf("env config load with err: %s", err.Error()))
  }
}

func GetEnvConfig(key string) string {
  return p1envConfig.GetString(key)
}

func MakeHttpAddr() string {
  ip := p1envConfig.GetString("APP_HTTP_ADDR")
  port := p1envConfig.GetString("APP_HTTP_PORT")
  return fmt.Sprintf("%s:%s", ip, port)
}
