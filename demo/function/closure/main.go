package main

import "fmt"

func main() {
  editOutside()
  remember()
}

// editOutside 在闭包内部修改引用的变量
func editOutside() {
  fmt.Println("editOutside")
  t1int := 1

  f := func() {
    t1int = 2
  }

  fmt.Println(t1int)
  f()
  fmt.Println(t1int)
}

// remember 闭包的记忆效应
// 被捕获到闭包中的变量，会跟随闭包生命周期一直存在
func remember() {
  fmt.Println("remember")
  t1inta := 1

  // 这里生成闭包方法时，f1 和 f2 里的 t1int 都是 1
  f1 := add1(t1inta)
  f2 := add1(t1inta)

  // 这里对 t1inta 进行修改时，不会影响 f1 和 f2 里的 t1int
  t1inta = 10

  t1intb := f1()
  fmt.Println(t1intb)
  t1intc := f2()
  fmt.Println(t1intc)
}

func add1(t1int int) func() int {
  return func() int {
    t1int++
    return t1int
  }
}
