package main

import (
  "errors"
  "fmt"
)

// 匿名方法
func main() {
  callWhenDefine()
  saveToVal()
  returnByFunc()
  funcHigh()
}

// callWhenDefine 定义时调用
func callWhenDefine() {
  func(t1 int) {
    fmt.Println(t1)
  }(100)
}

// saveToVal 用变量存储
func saveToVal() {
  f := func(t1 int) {
    fmt.Println(t1)
  }
  f(200)
}

// returnByFunc 作为方法返回值
func returnByFunc() {
  f := returnByFunc2()
  f(300)
}

func returnByFunc2() func(int) {
  return func(t1 int) {
    fmt.Println(t1)
  }
}

type myOperate func(x, y int) int // 声明方法类型

// 高阶方法：1、接受其他方法作为传入参数。2、把其他方法作为返回参数。
// 高阶方法可以实现闭包，但是这会影响程序的稳定和安全，需要注意。
func funcHigh() {
  fmt.Println("funcHigh")

  // 方法可以用变量存储
  var op myOperate

  op = getAddOperate()
  fmt.Println(myCalculate(3, 4, op))

  op = getMultiplyOperate()
  fmt.Println(myCalculate(3, 4, op))
}

// myCalculate 调用传入的方法
func myCalculate(x, y int, op myOperate) (int, error) {
  if nil == op {
    return 0, errors.New("op is nil")
  }
  return op(x, y), nil
}

// getAddOperate 返回加法方法
func getAddOperate() myOperate {
  return func(x, y int) int {
    return x + y
  }
}

// getMultiplyOperate 返回乘法方法
func getMultiplyOperate() myOperate {
  return func(x, y int) int {
    return x * y
  }
}
