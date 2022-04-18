package main

import (
  "fmt"
  "sync"
  "time"
)

var t1mapa map[string]int

var t1syncmapb sync.Map

func main() {
  // mapErrorWhenSync()
  mapSync()
}

// mapErrorWhenSync 触发并发环境读写 map 的错误
// fatal error: concurrent map read and map write
// 并发读写 map 会报错，map 内部对这种并发操作会进行检查
func mapErrorWhenSync() {
  t1mapa = make(map[string]int)
  go func() {
    for {
      t1mapa["a"] = 1
    }
  }()

  go func() {
    for {
      t1int := t1mapa["a"]
      fmt.Printf("%d,", t1int)
    }
  }()

  time.Sleep(5 * time.Second)
}

// mapSync sync.Map 提供了并发安全的 map
func mapSync() {
  go func() {
    for {
      t1syncmapb.Store("a", 1)
    }
  }()

  go func() {
    for {
      if t1int, ok := t1syncmapb.Load("a"); ok {
        fmt.Printf("%d,", t1int)
      }
    }
  }()

  time.Sleep(5 * time.Second)
}
