package main

import "fmt"

// 接口
type Invoker interface {
  Call(interface{})
}

// 方法
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
  // 调用方法本体
  f(p)
}

// 方法类型实现接口
func main() {
  // 将匿名方法转为 FuncCaller 类型，再赋值给接口
  invoker := FuncCaller(func(v interface{}) {
    fmt.Println("from function", v)
  })
  invoker.Call("hello")
}
