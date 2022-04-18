package main

import "fmt"

func main() {
  // 因为 u 是结构体，所以方法调用的时候它数据是不会变的
  u := User{
    Name: "a",
    Age:  10,
  }
  fmt.Printf("%+v\r\n", u)
  u.ChangeName("a Change")
  u.ChangeAge(100)
  fmt.Printf("%+v\r\n", u)

  // 因为 up 指针，所以内部的数据是可以被改变的
  up := &User{
    Name: "b",
    Age:  20,
  }
  fmt.Printf("%+v\r\n", up)
  // 自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合中包会包含所有值方法和所有指针方法。
  // 严格来讲，在值传递时只能调用到它的值方法。但是，go 会进行自动地转译，使得在值传递时也能调用到它的指针方法。
  // 语句被自动转译为 (&ac).SetSpecies(s)。因为 ChangeName 的接收器是结构体，所以 up 的数据还是不会变
  up.ChangeName("b Change")
  up.ChangeAge(200)
  fmt.Printf("%+v\r\n", up)
}

type User struct {
  Name string
  Age  int
}

// 结构体接收器，参数是值传递，相当于原数据的一个副本
func (u User) ChangeName(newName string) {
  u.Name = newName
}

// 指针接收器，引用传递，通过指针可以直接修改原数据
func (u *User) ChangeAge(newAge int) {
  u.Age = newAge
}
