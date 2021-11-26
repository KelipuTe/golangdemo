package main

import (
  "fmt"
)

//range关键字右边的代码被称为range表达式。
//range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代。
//range表达式的求值结果会被复制，被迭代的对象是range表达式结果值的副本而不是原数据。
//对于不同种类的range表达式结果值，for语句的迭代变量的数量可以有所不同。
func main() {
  arr1Num := [...]int{1, 2, 3, 4, 5, 6} //数组
  arr1NumLen := len(arr1Num)
  for i, val := range arr1Num {
    if i == arr1NumLen-1 {
      arr1Num[0] += val //修改的是副本
    } else {
      arr1Num[i+1] += val
    }
  }
  fmt.Println(arr1Num) //[7 3 5 7 9 11]

  sli1Num := []int{1, 2, 3, 4, 5, 6} //切片
  sli1NumLen := len(sli1Num)
  for i, val := range sli1Num {
    if i == sli1NumLen-1 {
      sli1Num[0] += val //修改的是原数据
    } else {
      sli1Num[i+1] += val
    }
  }
  fmt.Println(sli1Num) //[22 3 6 10 15 21]
}
