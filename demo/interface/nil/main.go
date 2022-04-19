package main

import "fmt"

// 定义一个结构体
type MyImplement struct{}

// 实现 fmt.Stringer 的 String 方法
func (m *MyImplement) String() string {
  return "hi"
}

// 在方法中返回 fmt.Stringer 接口
func GetStringer() fmt.Stringer {
  // 赋 nil
  var s *MyImplement = nil

  if s == nil {
    fmt.Println("s == nil")
  } else {
    fmt.Println("s != nil")
  }

  // 返回变量
  return s
}

func main() {
  // 判断返回值是否为 nil
  if GetStringer() == nil {
    fmt.Println("GetStringer() == nil")
  } else {
    fmt.Println("GetStringer() != nil")
  }
}
