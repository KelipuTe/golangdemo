package interfacehkn

import (
	"fmt"
	"testing"
)

// 在 Golang 里面，可以给方法类型添加方法，所以方法类型也可成为接口的实现。
// 所以如果有一个参数的类型是接口，那么传进去的东西可以是结构体类型也可以是一个方法类型。

type i9Call interface {
	f8Call(interface{})
}

type f8Demo func(interface{})

func (f8this f8Demo) f8Call(input interface{}) {
	fmt.Println("before f8Call f8Demo")
	f8this(input)
	fmt.Println("after f8Call f8Demo")
}

func TestFunction(p7s6t *testing.T) {
	var f8Demo f8Demo = func(input interface{}) { fmt.Println(input) }
	// 将 f8Demo 类型的变量赋值 i9Call 接口类型的变量
	var i9Call i9Call = f8Demo
	// 调用接口上的方法
	i9Call.f8Call("hello")
}
