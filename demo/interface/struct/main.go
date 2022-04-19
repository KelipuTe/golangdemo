package main

import "fmt"

func main() {
  var dog Dog = Dog{"little dog"}

  // 对于任何数据类型，只要它的方法集合中完全包含了一个接口的全部特征，那么它就是这个接口的实现类型。
  // 自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合中包会包含所有值方法和所有指针方法。
  // 所以 dog 是没有 SetName 方法的，不满足接口实现类型的定义。
  _, ok := interface{}(dog).(Pet)
  fmt.Printf("dog is Pet: %v\r\n", ok) // false
  _, ok = interface{}(&dog).(Pet)
  fmt.Printf("*dog is Pet: %v\r\n", ok) // true

  var pet Pet = &dog // 这时 pet 的静态类型是 Pet（接口），动态类型是 *Dog（接口实现）
  fmt.Printf("pet is %s, name is %q\r\n", pet.Category(), pet.Name())

  var pet2 Pet2 = dog    // Pet2 去掉了指针方法，所以这里 pet2 的赋值就变成了 dog 的副本
  dog.SetName("big dog") // 这里修改了 dog 的值，但是副本不受影响，所以 pet2 的值不变
  fmt.Printf("pet is %s, name is %q\r\n", pet2.Category(), pet2.Name())
}

type Pet interface {
  SetName(name string)
  Name() string
  Category() string
}

type Pet2 interface {
  Name() string
  Category() string
}

type Dog struct {
  name string
}

func (dog *Dog) SetName(name string) {
  dog.name = name
}

func (dog Dog) Name() string {
  return dog.name
}

func (dog Dog) Category() string {
  return "dog"
}
