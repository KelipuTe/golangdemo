package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  gen := func(ctx context.Context) <-chan int {
    dst := make(chan int)
    n := 1
    go func() {
      for {
        select {
        case <-ctx.Done():
          fmt.Println("done")
          return // return结束该goroutine，防止泄露
        case dst <- n:
          n++
        }
      }
    }()
    return dst
  }
  ctx, cancel := context.WithCancel(context.Background())

  for n := range gen(ctx) {
    fmt.Println(n)
    if n == 5 {
      cancel()
      break
    }
  }

  time.Sleep(1 * time.Second)
}
