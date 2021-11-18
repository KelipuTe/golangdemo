package main

import (
  "fmt"
  "time"
)

func waitData(chan1 chan int) {
  for {
    data := <-chan1
    if data == 0 {
      break
    }
    fmt.Println(data)
  }
}

func main() {
  ticker1 := time.NewTicker(1 * time.Second)
  defer ticker1.Stop()

  chan1 := make(chan int)
  go waitData(chan1)

  for range ticker1.C {
    chan1 <- 1
  }
}
