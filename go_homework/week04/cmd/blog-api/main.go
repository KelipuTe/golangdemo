package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"

  "github.com/gin-gonic/gin"

  api_v1_http "demo-golang/go_homework/week04/api/v1/http"
  tp_viper "demo-golang/go_homework/week04/internal/third-party/viper"
)

var p1ge *gin.Engine
var c1signalquit chan bool = make(chan bool)

func main() {
  log.Printf("week04 version :%s", tp_viper.VERSION)
  p1ge = gin.Default()

  // p1ge.Use(biz_mid.ReqInMiddleWare)

  api_v1_http.HttpRegister(p1ge)

  // http服务
  httpAddr := tp_viper.MakeHttpAddr()
  p1httpServer := &http.Server{
    Addr:    httpAddr,
    Handler: p1ge,
  }
  go p1httpServer.ListenAndServe()
  log.Printf("http server running in:%s", httpAddr)

  go waitSignal()

  // 服务启动完成
  // 会调到注册中心去，带上本机的地址和服务类型，注册服务

  select {
  case <-c1signalquit:
    log.Println("get signal, stop http server...")
    // 关闭的时候，先行去服务中心注销服务
    // 一般这里会等待几秒，防止服务中心那边处理有延迟
    // 然后关闭服务
    p1httpServer.Shutdown(context.TODO())
    log.Println("http server shutdown by signal")
  }
}

func waitSignal() {
  c1signal := make(chan os.Signal)
  signal.Notify(c1signal, os.Interrupt)
  <-c1signal
  c1signalquit <- true
}
