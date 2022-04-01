package main

import (
  "context"
  "errors"
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"

  "golang.org/x/sync/errgroup"
)

var (
  // 假装有几个服务
  service1 *http.Server
  service2 *http.Server
  service3 *http.Server
  // 信号退出
  c1signalquit chan bool = make(chan bool)
  // err 退出
  c1errquit chan bool = make(chan bool, 3)
)

// 假装初始化几个服务
// service3 配置错误的 tcp 地址，用于触发一错全关
func init() {
  handler1 := http.NewServeMux()
  handler1.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    log.Println("127.0.0.1:9502")
  })
  service1 = &http.Server{
    Addr:    "127.0.0.1:9502",
    Handler: handler1,
  }

  handler2 := http.NewServeMux()
  handler2.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    log.Println("127.0.0.1:9503")
  })
  service2 = &http.Server{
    Addr:    "127.0.0.1:9503",
    Handler: handler2,
  }

  service3 = &http.Server{
    Addr: "111",
  }
}

// 等待 ctrl+c 信号
func waitSignal() {
  c1signal := make(chan os.Signal)
  signal.Notify(c1signal, os.Interrupt)
  <-c1signal
  // 等到 ctrl+c 就退出
  c1signalquit <- true
}

func main() {
  log.Println("main,start")

  p1group := &errgroup.Group{}

  log.Println("errgroup,GO")
  p1group.Go(service1Start)
  p1group.Go(service2Start)
  // p1group.Go(service3Start) // 用于触发 error

  log.Println("wait signal")
  go waitSignal()

  log.Println("wait select")
  select {
  case <-c1signalquit:
    log.Println("c1signalquit")
    service1Stop()
    service2Stop()
    service3Stop()
  case <-c1errquit:
    log.Println("c1errquit")
    service1Stop()
    service2Stop()
    service3Stop()
  }

  if err := p1group.Wait(); nil != err {
    log.Println("errgroup,Wait,", err.Error())
  }

  log.Println("main,done")
}

func service1Start() error {
  if err := service1.ListenAndServe(); nil != err {
    log.Println("service1,err,", err.Error())
    c1errquit <- true
    return err
  }
  return nil
}

func service1Stop() {
  service1.Shutdown(context.TODO())
  log.Println("service1,stop")
}

func service2Start() error {
  defer func() error {
    if p := recover(); p != nil {
      fmt.Printf("recover: %s\n", p)
      return errors.New("panic")
    }
    return nil
  }()
  if err := service2.ListenAndServe(); nil != err {
    log.Println("service2,err,", err.Error())
    c1errquit <- true
    return err
  }
  return nil
}

func service2Stop() {
  service2.Shutdown(context.TODO())
  log.Println("service2,stop")
}

func service3Start() error {
  if err := service3.ListenAndServe(); nil != err {
    log.Println("service3,err,", err.Error())
    c1errquit <- true
    return err
  }
  return nil
}

func service3Stop() {
  service3.Shutdown(context.TODO())
  log.Println("service3,stop")
}
