package main

import (
  "fmt"
  "sync"
  "time"
)

//sync.RWMutex，读写互斥锁，可以单独锁写或锁读。

//sync.Cond.Wait()
//阻塞时，会解锁当前条件变量基于的互斥锁，阻塞当前goroutine，等待通知到来。
//通知到来时，会唤醒当前goroutine，重新锁定当前条件变量基于的互斥锁。

//在包裹条件变量的Wait方法的时候，应该使用for语句。
//场景1：并发环境下，可能有多个goroutine在等待同一个共享资源。
//如果被唤醒后，发现共享资源的状态不符合它的要求，那么就应该再次调用Wait方法。
//场景2：共享资源的状态大于2个，但是状态在每次改变后的结果只有一个。
//单一的结果不能满足所有goroutine的条件，未被满足的goroutine需要继续等待和检查。

//sync.Cond.Wait()，会把当前的goroutine添加到通知队列的队尾。
//sync.Cond.Signal()，从通知队列的队首开始，找第一个符合条件的goroutine并唤醒。
//sync.Cond.Broadcast()，唤醒通知队列所有符合条件的goroutine。
//所以上面场景2的情况，需要使用Broadcast()，Signal()唤醒的不一定能满足条件。

func main() {
  var (
    mailTime    int           = 3 //发收信测试几次
    mailbox     int               //信箱
    mailboxLock sync.RWMutex      //信箱上的锁
    sendCond    *sync.Cond        //发信条件变量
    receiveCond *sync.Cond        //收信条件变量
    chanSign    chan struct{}     //信号通道

    funcSendMail    func() //发信
    funcReceiveMail func() //收信
  )
  sendCond = sync.NewCond(&mailboxLock)
  receiveCond = sync.NewCond(mailboxLock.RLocker())
  chanSign = make(chan struct{}, 2) //2个信号，发信结束，收信结束

  funcSendMail = func() {
    defer func() { chanSign <- struct{}{} }()

    for i := 1; i <= mailTime; i++ {
      time.Sleep(time.Millisecond * 500)
      mailboxLock.Lock()
      for mailbox == 1 {
        sendCond.Wait()
      }
      fmt.Printf("send-%d:mailbox empty\r\n", i)
      mailbox = 1
      fmt.Printf("send-%d:send mail\r\n", i)
      mailboxLock.Unlock()
      receiveCond.Signal()
    }
  }

  funcReceiveMail = func() {
    defer func() { chanSign <- struct{}{} }()

    for j := 1; j <= mailTime; j++ {
      time.Sleep(time.Millisecond * 500)
      mailboxLock.RLock()
      for mailbox == 0 {
        receiveCond.Wait()
      }
      fmt.Printf("receive-%d:mailbox full\r\n", j)
      mailbox = 0
      fmt.Printf("receive-%d:receive mail\r\n", j)
      mailboxLock.RUnlock()
      sendCond.Signal()
    }
  }

  go funcSendMail()
  go funcReceiveMail()

  <-chanSign
  <-chanSign
}
