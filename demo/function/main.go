package main

import (
  "errors"
  "fmt"
)

//高阶函数：1、接受其他函数作为传入参数。2、把其他函数作为返回参数。
//高阶函数可以实现闭包，但是这会影响程序的稳定和安全，需要注意。

func main() {
  var op myOperate

  op = getAddOperate()
  fmt.Println(myCalculate(3, 4, op))

  op = getMultiplyOperate()
  fmt.Println(myCalculate(3, 4, op))
}

type myOperate func(x, y int) int //声明函数类型

//调用传入的函数
func myCalculate(x, y int, op myOperate) (int, error) {
  if nil == op {
    return 0, errors.New("op is nil")
  }
  return op(x, y), nil
}

//返回加法函数
func getAddOperate() myOperate {
  return func(x, y int) int {
    return x + y
  }
}

//返回乘法函数
func getMultiplyOperate() myOperate {
  return func(x, y int) int {
    return x * y
  }
}
