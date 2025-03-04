package interfacehkn

import (
	"fmt"
	"testing"
)

// 可以给方法类型添加方法，所以方法类型也可成为接口的实现。
// 如果有一个参数的类型是接口，那么传进去的东西可以是结构体类型也可以是一个方法。

type outerFunc interface {
	outerFunc(string)
}

type innerFunc func(string)

func (t innerFunc) outerFunc(input string) {
	fmt.Println("before innerFunc")
	t(input)
	fmt.Println("after innerFunc")
}

func TestFunc(t *testing.T) {
	var inner innerFunc = func(input string) { fmt.Println(input) }
	// 将 innerFunc 类型的变量赋值 outerFunc 接口类型的变量
	var outer outerFunc = inner
	// 调用接口上的方法
	outer.outerFunc("hello")
}
