package main

import "fmt"

// 定义类型别名
type NewNameInt = int

// 别名不能定义方法，报错
// func (t1 NewNameInt) Test() {
// }

// 定义新的类型
type NewTypeInt int

// 新的类型可以定义方法
func (t1 NewTypeInt) Test() {
}

func main() {
  var t1nni NewNameInt = 100
  var t1nti NewTypeInt = 200

  fmt.Printf("t1nni: %d, t1nni: %T\n", t1nni, t1nni)
  fmt.Printf("t1nti: %d, t1nti: %T\n", t1nti, t1nti)
}
