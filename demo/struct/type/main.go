package main

import "fmt"

func main() {
  fish := Fish{}
  fish.Swim()

  // fake 无法调用原来 Fish 的方法，fake.Swim() 会编译错误
  fake := FakeFish{}
  fake.FakeSwim()

  // 将 FakeFish 转换为 Fish，就可以调用 Fish 的方法
  t1fish := Fish(fake)
  t1fish.Swim()

  // newFish 这种用别名的，可以直接调用原来 Fish 的方法
  newFish := NewFish{}
  newFish.Swim()
}

type Fish struct {
}

func (f Fish) Swim() {
  fmt.Printf("Fish, Swim\r\n")
}

// 定义了一个新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
  fmt.Printf("FakeFish, FakeSwim\r\n")
}

// 为类型起了个别名
type NewFish = Fish
