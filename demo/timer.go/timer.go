package main

import (
  "fmt"
  "time"
)

func checkTimeout(chan1 chan int, chan2 chan bool, timer1 *time.Timer) {
  for {
    select {
    case num := <-chan1:
      fmt.Println(num)
    case <-timer1.C:
      fmt.Println("timeout")
      chan2 <- true
    }
  }
}

func main() {
  timer1 := time.NewTimer(5 * time.Second)

  chan1 := make(chan int)
  chan2 := make(chan bool)

  go checkTimeout(chan1, chan2, timer1)

  for i := 0; i < 5; i++ {
    chan1 <- i
    time.Sleep(time.Second)
  }

  <-chan2

  fmt.Println("done")
}
