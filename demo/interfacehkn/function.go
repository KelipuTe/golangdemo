package interfacehkn

import "fmt"

type i9Call interface {
	f8Call(interface{})
}

// 在 Golang 里面，可以给方法类型添加方法，所以方法类型也可成为接口的实现。
// 所以如果有一个参数的类型是接口，那么传进去的东西可以是结构体类型也可以是一个方法类型。
type f8Demo func(interface{})

func (f8this f8Demo) f8Call(input interface{}) {
	fmt.Println("before f8Call f8Demo")
	f8this(input)
	fmt.Println("after f8Call f8Demo")
}
