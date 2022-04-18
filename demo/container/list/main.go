package main

import (
  "container/list"
  "fmt"
)

func main() {
  var t1p1lista *list.List = list.New()

  // 头插尾插都会返回元素句柄
  element := t1p1lista.PushFront(100)
  // 头插
  t1p1lista.PushFront(50)
  // 尾插
  t1p1lista.PushBack(150)

  // 元素句柄可用于向该元素前后插入元素
  t1p1lista.InsertBefore(90, element)
  t1p1lista.InsertAfter(110, element)

  // 从头遍历列表，i 拿到的是每个元素的元素句柄
  for i := t1p1lista.Front(); nil != i; i = i.Next() {
    fmt.Printf("%+v,", i.Value)
  }
  fmt.Printf("\r\n")

  // 删除的时候也需要元素句柄
  t1p1lista.Remove(element)

  for i := t1p1lista.Front(); nil != i; i = i.Next() {
    fmt.Printf("%+v,", i.Value)
  }
  fmt.Printf("\r\n")
}
