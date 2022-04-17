package main

import "fmt"

func main() {
  array1()
  array2()
}

// array1 一维数组
func array1() {
  var t1arr1a [3]int
  // len() 返回数组中元素的个数，cap() 返回数组的容量
  fmt.Printf("t1arr1a: %+v, len: %d, cap: %d\r\n", t1arr1a, len(t1arr1a), cap(t1arr1a))

  // 一维数组可以按下标读取和修改
  fmt.Printf("t1arr1a[1]: %d\r\n", t1arr1a[1])
  t1arr1a[1] = 2
  fmt.Printf("t1arr1a[1]: %d\r\n", t1arr1a[1])

  // 按下标遍历数组
  for i := 0; i < len(t1arr1a); i++ {
    fmt.Printf("t1arr1a[%d]: %d; ", i, t1arr1a[i])
  }
  fmt.Printf("\r\n")

  // 遍历数组输出下标和值
  for i, v := range t1arr1a {
    fmt.Printf("t1arr1a, i: %d, v: %+v; ", i, v)
  }
  fmt.Printf("\r\n")

  // 定义数组时，如果直接赋值，那么大括号里一个也不能少
  t1arr1b := [3]int{1, 2, 3}
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))
}

// array2 二（多）维数组
func array2() {
  var t1arr2a [3][2]int

  // 多维数组也可以按下标读取和修改，但是需要注意每一维的类型
  fmt.Printf("t1arr2a[1]: %d\r\n", t1arr2a[1])
  t1arr2a[1] = [2]int{1, 2}
  fmt.Printf("t1arr2a[1]: %d\r\n", t1arr2a[1])

  // 按下标遍历数组
  for i := 0; i < len(t1arr2a); i++ {
    fmt.Printf("t1arr1a[%d]: %+v; ", i, t1arr2a[i])
  }
  fmt.Printf("\r\n")

  // 遍历数组输出下标和值
  for i, v := range t1arr2a {
    fmt.Printf("t1arr2a, i: %d, v: %+v; ", i, v)
  }
  fmt.Printf("\r\n")
}
