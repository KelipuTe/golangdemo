package main

import (
  "errors"
  "fmt"
)

// panic 一旦发生，控制权就会讯速地沿着调用栈的反方向传播。
// 在引发 panic 的语句之后的所有语句，都不会有任何执行机会。
// 所以 recover 需要结合 defer 写在 panic 前面。

func main() {
  fmt.Println("main in")
  defer func() {
    fmt.Println("defer in")
    if p := recover(); p != nil {
      fmt.Printf("recover: %s\n", p)
    }
    fmt.Println("defer exit")
  }()
  fmt.Println("painc")
  panic(errors.New("painc"))
  // panic 之后的语句不会执行
  // fmt.Println("main exit")
}
