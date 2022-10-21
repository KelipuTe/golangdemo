package main

import (
  "context"
  "errors"
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"

  "golang.org/x/sync/errgroup"
)

// 第 3 周作业
// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

// 实现了 ctrl+c 信号退出和 error 退出两个场景
// 其中 error 退出，没找到合适的运行时抛出 error 的手段，所以选择在 http ListenAndServe 启动时抛出 tcp 配置错误

var (
  // 假装有几个服务
  service1 *http.Server
  service2 *http.Server
  service3 *http.Server
  // ctrl+c 信号退出
  c1signalquit chan bool = make(chan bool)
  // error 退出
  c1errquit chan bool = make(chan bool, 3)
)

// 假装初始化几个服务
// service3 配置错误的 tcp 地址，用于触发一错全关
func init() {
  // 两个正常的配置
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

  // 错误的配置
  service3 = &http.Server{
    Addr: "111",
  }
}

// 等待 ctrl+c 信号
func waitSignal() {
  c1signal := make(chan os.Signal)
  signal.Notify(c1signal, os.Interrupt)
  <-c1signal
  // 被阻塞的 c1signal 等到 ctrl+c 信号就给 c1signalquit 发送数据
  c1signalquit <- true
}

func main() {
  log.Println("main,start")

  // 信号退出的场景，需要注释掉下面的 sleep 和 service3 的启动代码
  // 这段代码放前面或者后面都会被 time.Sleep 阻塞住
  log.Println("wait signal")
  go waitSignal()

  p1group := &errgroup.Group{}

  log.Println("errgroup,GO")
  p1group.Go(service1Start)
  p1group.Go(service2Start)

  // 这里先启动上面两个正常的服务，service1 和 service2 启动之后在 service3 启动之前，是可以访问的
  // service1 和 service2 启动之后延迟 10 秒启动配置错误的 service3，用于模拟出现错误，一起退出的场景
  time.Sleep(10 * time.Second)
  p1group.Go(service3Start)

  // 打印一下是哪种情况，然后关闭所有的 service
  log.Println("wait select")
  select {
  case <-c1signalquit:
    log.Println("c1signalquit")
    stopAllService()
  case <-c1errquit:
    log.Println("c1errquit")
    stopAllService()
  }

  // 这里会得到两种错误
  // 一种是 service3 故意触发的 tcp 配置错误
  // 另一种是信号触发 service 关闭之后的错误
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

func stopAllService() {
  service1Stop()
  service2Stop()
  service3Stop()
}
