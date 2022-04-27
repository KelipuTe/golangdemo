package main

import "fmt"

func main() {
  // sliceInit()
  // sliceAppend()
  // sliceMake()
  sliceFor()
  // sliceFromArray()
  // sliceFromSlice()
  // sliceCopy()
  // sliceDelete()
  // slice2()
}

func sliceInit() {
  fmt.Printf("sliceInit\r\n")

  // slice 的底层实现是个结构体，/src/runtime/slice.go
  // 只进行了声明的切片，长度和容量都为 0
  var t1arr1a []int
  fmt.Printf("t1arr1a: %+v, len: %d, cap: %d\r\n", t1arr1a, len(t1arr1a), cap(t1arr1a))
}

func sliceAppend() {
  fmt.Printf("sliceAppend\r\n")

  // 使用 append() 向切片追加元素的时候，会发生扩容。
  // 如果没有扩容，切片指向的底层数组不会变。
  // 如果发生扩容，切片指向的底层数组会发生变化。
  var t1arr1a []int
  fmt.Printf("t1arr1a: %+v, t1arr1a: %p, len: %d, cap: %d\r\n", t1arr1a, t1arr1a, len(t1arr1a), cap(t1arr1a))
  t1arr1a = append(t1arr1a, 1)
  fmt.Printf("t1arr1a: %+v, t1arr1a: %p, len: %d, cap: %d\r\n", t1arr1a, t1arr1a, len(t1arr1a), cap(t1arr1a))
  t1arr1a = append(t1arr1a, 2)
  fmt.Printf("t1arr1a: %+v, t1arr1a: %p, len: %d, cap: %d\r\n", t1arr1a, t1arr1a, len(t1arr1a), cap(t1arr1a))
  t1arr1a = append(t1arr1a, 3)
  fmt.Printf("t1arr1a: %+v, t1arr1a: %p, len: %d, cap: %d\r\n", t1arr1a, t1arr1a, len(t1arr1a), cap(t1arr1a))
  t1arr1a = append(t1arr1a, 4)
  fmt.Printf("t1arr1a: %+v, t1arr1a: %p, len: %d, cap: %d\r\n", t1arr1a, t1arr1a, len(t1arr1a), cap(t1arr1a))
}

func sliceMake() {
  fmt.Printf("sliceMake\r\n")

  // make() 传入两个参数时，表示分开设置长度和容量
  t1arr1b := make([]int, 2, 4)
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))

  // make() 传入两个参数时，表示长度和容量相等
  // make([]int, 4) 等价于 make([]int, 4, 4)
  t1arr1c := make([]int, 4)
  fmt.Printf("t1arr1c: %+v, len: %d, cap: %d\r\n", t1arr1c, len(t1arr1c), cap(t1arr1c))
}

func sliceFor() {
  fmt.Printf("sliceFor\r\n")

  t1arr1a := []int{11, 22, 33, 44}

  // 可以用下标遍历
  for i := 0; i < len(t1arr1a); i++ {
    fmt.Printf("t1arr1a, i: %d,v: %d; ", i, t1arr1a[i])
  }
  fmt.Printf("\r\n")

  // 也可以用 range 遍历
  for i, v := range t1arr1a {
    fmt.Printf("t1arr1a, i: %d,v: %d; ", i, v)
  }
  fmt.Printf("\r\n")

  // 用 range 遍历，只接收一个返回值时，收到的是下标
  for i := range t1arr1a {
    fmt.Printf("t1arr1a, i: %d; ", i)
  }
  fmt.Printf("\r\n")
}

// sliceFromArray 从数组生成新的切片
func sliceFromArray() {
  fmt.Printf("sliceFromArray\r\n")

  t1arr1a := [4]int{1, 2, 3, 4}

  // 注意左闭右开
  t1arr1d := t1arr1a[1:2]
  fmt.Printf("t1arr1d: %+v, len: %d, cap: %d\r\n", t1arr1d, len(t1arr1d), cap(t1arr1d))
  // 左边界不写表示从 0 开始
  t1arr1b := t1arr1a[:2]
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))
  // 右边界不写表示一直到末尾
  t1arr1c := t1arr1a[2:]
  fmt.Printf("t1arr1c: %+v, len: %d, cap: %d\r\n", t1arr1c, len(t1arr1c), cap(t1arr1c))
}

// sliceFromSlice 从切片生成新的切片
func sliceFromSlice() {
  fmt.Printf("sliceFromSlice\r\n")

  t1arr1a := [4]int{1, 2, 3, 4}

  t1arr1d := t1arr1a[0:3]
  fmt.Printf("t1arr1d: %+v, len: %d, cap: %d\r\n", t1arr1d, len(t1arr1d), cap(t1arr1d))
  t1arr1b := t1arr1d[1:2]
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))
}

// sliceCopy 切片复制
func sliceCopy() {
  fmt.Printf("sliceCopy\r\n")

  t1arr1a := []int{1, 2, 3, 4}

  // 容量够的时候，全部复制，多出来的位置，保持不变
  t1arr1b := make([]int, 5)
  copy(t1arr1b, t1arr1a)
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))

  // 容量不够的时候，只复制容量容得下的部分
  t1arr1c := make([]int, 2)
  copy(t1arr1c, t1arr1a)
  fmt.Printf("t1arr1c: %+v, len: %d, cap: %d\r\n", t1arr1c, len(t1arr1c), cap(t1arr1c))
}
func sliceDelete() {
  fmt.Printf("sliceDelete\r\n")

  t1arr1a := []int{1, 2, 3, 4}

  // 切片没有用下标删除的操作
  // 删除切片中的一个元素，需要把这个元素的前后元素连接起来
  t1arr1b := append(t1arr1a[:1], t1arr1a[2:]...)
  fmt.Printf("t1arr1b: %+v, len: %d, cap: %d\r\n", t1arr1b, len(t1arr1b), cap(t1arr1b))
}

// 二（多）维切片
func slice2() {
  fmt.Printf("slice2\r\n")

  t1arr2a := [][]int{{10}, {100, 200}}

  // 多维切片也可以按下标读取和修改，但是需要注意每一维的类型
  fmt.Printf("t1arr2a[1]: %d\r\n", t1arr2a[1])
  t1arr2a[1] = []int{10, 20}
  fmt.Printf("t1arr2a[1]: %d\r\n", t1arr2a[1])

  // 按下标遍历切片
  for i := 0; i < len(t1arr2a); i++ {
    fmt.Printf("t1arr1a, i: %d, v: %+v; ", i, t1arr2a[i])
  }
  fmt.Printf("\r\n")
}
