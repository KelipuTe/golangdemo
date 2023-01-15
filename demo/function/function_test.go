package interfacehkn

import (
	"fmt"
	"testing"
)

func TestFunction(p7s6t *testing.T) {
	// 将匿名方法赋值给 f8Demo 类型的变量
	var f8Demo f8Demo = func(input interface{}) { fmt.Println(input) }
	// 调用方法
	f8Demo("hello")
	// 调用方法的方法
	f8Demo.f8Call("hello2")
	f8Demo.f8CallV2("hello2")
}
