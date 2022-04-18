package main

import "fmt"

func main() {
  innerStruct()
}

type A struct {
  ax, ay int
}

type B struct {
  A
  bx, by int
}

// innerStruct 内嵌结构体
func innerStruct() {
  b := B{A{1, 2}, 3, 4}
  // 嵌入结构体的成员，可以通过外部结构体的实例直接访问
  fmt.Println(b.ax, b.ay, b.bx, b.by)
  // 内嵌结构体的字段名是它的类型名，外部结构体可以通过类名访问
  fmt.Println(b.A)
}

type C1 struct {
  cx, cy int
}

type C2 struct {
  cx, cy int
}

type D struct {
  C1
  C2
}

// innerConflict 内嵌结构体字段名冲突
func innerConflict() {
  d := D{}
  // 当内嵌结构体的字段名冲突时，通过外部结构体的实例直接访问会报错。
  // 编译器会告知选择器 cx 引起歧义，因为不知道 cx 是 C1 还是 C2 的。
  // d.cx = 1
  // 这种情况需要通过类名访问
  d.C1.cx = 1
}
