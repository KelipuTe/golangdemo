package main

import "fmt"

// 用 const 模拟枚举

const (
  Type1 = iota
  Type2
  Type3
  Type4
)

// 如果想将枚举值转换为字符串
// 就需要自定义类型，然后添加 String() 方法

type TypeToString int

const (
  TypeA TypeToString = iota
  TypeB
  TypeC
  TypeD
)

func main() {
  fmt.Println(Type1, Type2, Type3, Type4)
  fmt.Println(TypeA, TypeB, TypeC, TypeD)
}

func (t1 TypeToString) String() string {
  switch t1 {
  case TypeA:
    return "TypeA"
  case TypeB:
    return "TypeB"
  case TypeC:
    return "TypeC"
  case TypeD:
    return "TypeD"
  }
  return "NONE"
}
