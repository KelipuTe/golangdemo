package main

import (
  "fmt"
  "sort"
)

func main() {
  fmt.Println("demo run")
  RunDemo1()
  RunDemo2()
  fmt.Println("demo finish")
}

//结构体实现接口
func RunDemo1() {
  var pts1 *DemoStruct1 = &DemoStruct1{}
  fmt.Printf("var pts1:%p\r\n", pts1)
  fmt.Println("pts1.ReadData11(),", pts1.ReadData11())
  fmt.Println("pts1.WriteData11(42),", pts1.WriteData11(42))
  fmt.Println("pts1.ReadData11(),", pts1.ReadData11())

  var ts2 DemoStruct1 = DemoStruct1{}
  fmt.Printf("var ts2:%p\r\n", &ts2)
  fmt.Println("ts2.ReadData12(),", ts2.ReadData12())
  fmt.Println("ts2.WriteData12(42),", ts2.WriteData12(42))
  fmt.Println("ts2.ReadData12(),", ts2.ReadData12())
}

//结构体1
type DemoStruct1 struct {
  Data int
}

//接口11
type DemoInterface11 interface {
  WriteData11(data int) error
  ReadData11() int
}

//实现接口11，指针形式
func (pts *DemoStruct1) WriteData11(data int) error {
  fmt.Printf("WriteData1 var pts:%p\r\n", pts) //传入的就是原参数的地址
  pts.Data = data                              //会修改原参数
  fmt.Printf("WriteData1 pts.Data:%d\r\n", pts.Data)
  return nil
}

func (pts *DemoStruct1) ReadData11() int {
  return pts.Data
}

//接口12
type DemoInterface12 interface {
  WriteData12(data int) error
  ReadData12() int
}

//实现接口12，非指针形式
func (ts DemoStruct1) WriteData12(data int) error {
  fmt.Printf("WriteData2 var ts:%p\r\n", &ts) //传入的原参数被复制了
  ts.Data = data                              //修改的是复制的新参数，不会修改原参数
  fmt.Printf("WriteData2 ts.Data:%d\r\n", ts.Data)
  return nil
}

func (ts DemoStruct1) ReadData12() int {
  return ts.Data
}

//使用sort.Interface接口对结构体2切片排序
func RunDemo2() {
  var sli1demo DemoStruct2Slice = DemoStruct2Slice{
    DemoStruct2{"10", 10}, DemoStruct2{"20", 20}, DemoStruct2{"50", 50}, DemoStruct2{"40", 40}, DemoStruct2{"30", 30},
  }
  fmt.Println(sli1demo)
  sort.Sort(sli1demo)
  fmt.Println(sli1demo)

  //也可以用这个方法对切片排序
  sort.Slice(sli1demo, func(i, j int) bool { return sli1demo[i].weight > sli1demo[j].weight })
  fmt.Println(sli1demo)
}

//结构体2
type DemoStruct2 struct {
  name   string
  weight int
}

//给结构体2切片定义一个类型
type DemoStruct2Slice []DemoStruct2

//实现sort.Interface接口的获取元素数量方法
func (sli1 DemoStruct2Slice) Len() int {
  return len(sli1)
}

//实现sort.Interface接口的比较元素方法
func (sli1 DemoStruct2Slice) Less(i, j int) bool {
  return sli1[i].weight < sli1[j].weight
}

//实现sort.Interface接口的交换元素方法
func (sli1 DemoStruct2Slice) Swap(i, j int) {
  sli1[i], sli1[j] = sli1[j], sli1[i]
}
