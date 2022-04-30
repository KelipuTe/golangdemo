package main

import (
  "fmt"
  "reflect"
)

type User struct {
  Id    int
  Name  string
  User2 User2
}

type User2 struct {
  Id   int
  Name string
}

func main() {
  u11 := newUser1()
  fmt.Printf("u11,%p\r\n", u11)
  fmt.Println(reflect.TypeOf(u11))
  // u11 = u11.(*User)
  fmt.Printf("u11,%p\r\n", u11)
  fmt.Printf("u11,%p\r\n", &u11.User2)
  fmt.Println(reflect.TypeOf(u11))

  u22 := newUser2()
  fmt.Printf("u22,%p\r\n", &u22)
  fmt.Println(reflect.TypeOf(u22))
  // u22 = u22.(User)
  fmt.Printf("u22,%p\r\n", &u22)
  fmt.Printf("u22,%p\r\n", &u22.User2)
  fmt.Println(reflect.TypeOf(u22))
}

func newUser1() *User {
  u1 := &User{}
  u1.User2 = User2{}
  fmt.Printf("u1,%p\r\n", u1)
  fmt.Printf("u11,%p\r\n", &u1.User2)
  return u1
}

func newUser2() User {
  u2 := User{}
  u2.User2 = User2{}
  fmt.Printf("u2,%p\r\n", &u2)
  fmt.Printf("u22,%p\r\n", &u2.User2)
  return u2
}
