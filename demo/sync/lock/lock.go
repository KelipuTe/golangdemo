package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  a := 0
  var lock sync.Mutex
  for i := 0; i < 100; i++ {
    go func() {
      lock.Lock()
      defer lock.Unlock()
      a += 1
      fmt.Println(a)
    }()
  }
  time.Sleep(5 * time.Second)
}
