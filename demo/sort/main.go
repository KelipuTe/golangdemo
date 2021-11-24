package main

import (
  "fmt"
  "sort"
)

//使用sort.Interface接口对结构体切片排序
func main() {
  sli1demo := DemoStructSlice{
    DemoStruct{"10", 10},
    DemoStruct{"20", 20},
    DemoStruct{"50", 50},
    DemoStruct{"40", 40},
    DemoStruct{"30", 30},
  }
  fmt.Println(sli1demo)
  sort.Sort(sli1demo)
  fmt.Println(sli1demo)

  //也可以用这个方法对切片排序
  sort.Slice(sli1demo, func(i, j int) bool { return sli1demo[i].weight > sli1demo[j].weight })
  fmt.Println(sli1demo)
}

//结构体2
type DemoStruct struct {
  name   string
  weight int
}

//给结构体2切片定义一个类型
type DemoStructSlice []DemoStruct

//实现sort.Interface接口的获取元素数量方法
func (sli1 DemoStructSlice) Len() int {
  return len(sli1)
}

//实现sort.Interface接口的比较元素方法
func (sli1 DemoStructSlice) Less(i, j int) bool {
  return sli1[i].weight < sli1[j].weight
}

//实现sort.Interface接口的交换元素方法
func (sli1 DemoStructSlice) Swap(i, j int) {
  sli1[i], sli1[j] = sli1[j], sli1[i]
}
