package main

import "fmt"

func main() {
  RunDemo1()
}

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

//测试结构体1
type DemoStruct1 struct {
  Data int
}

//测试接口11
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

//测试接口12
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
