package main

import (
  "flag"
  "fmt"
)

// flag.String(name string, value string, usage string)
// name，名称；value，默认值；usage，使用信息；
// 运行 go run main.go，p1mode 的值就是 mode1
// 运行 go run main.go --mode=mode2，p1mode 的值就是 mode2。
var p1mode = flag.String("mode", "mode1", "flag mode")

func main() {
  t1int := 100
  fmt.Printf("t1int: %d, t1int: %T\r\n", t1int, t1int)

  t1p1a := &t1int
  fmt.Printf("*t1p1a: %d, t1p1a: %T, t1p1a: %p\r\n", *t1p1a, t1p1a, t1p1a)

  // new() 为对应类型分配内存然后返回指针
  t1p1b := new(int)
  *t1p1b = 200
  fmt.Printf("*t1p1b: %d, t1p1b: %T, t1p1b: %p\r\n", *t1p1b, t1p1b, t1p1b)

  // 指针可以用来获取命令行的输入信息
  // flag.Parse() 从 arguments 中解析注册的 flag。必须在所有 flag 都注册好而未访问其值时执行。
  flag.Parse()
  fmt.Printf("p1mode: %s\r\n", *p1mode)
}
