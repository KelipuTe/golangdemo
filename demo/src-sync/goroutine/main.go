package main

import (
  "fmt"
  "time"
)

//让goroutine依次输出1-100
func main() {
  var total int = 100
  for i := 1; i < 100; i++ {
    go trigger(i)
  }
  //main函数是主goroutine，一旦main中的代码执行完毕，当前的程序就会结束运行。
  //所以最后这里需要有一个等待的操作，等待上面的goroutine执行完。
  trigger(total)
}

var index int = 1

func trigger(i int) {
  for {
    if index == i {
      fmt.Printf("%d,", i)
      index++
      break
    }
    //调度器在需要的时候只会对正在运行的goroutine发出通知，试图让它停下来。但是，它却不会也不能强行让一个goroutine停下来。
    //所以，如果一条for语句过于简单的话，那么当前的goroutine就可能不会去正常响应（或者说没有机会响应）调度器的停止通知。
    time.Sleep(time.Millisecond)
  }
}
