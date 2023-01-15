package interfacehkn

import (
	"fmt"
	"testing"
)

func TestFunction(p7s6t *testing.T) {
	var f8Demo f8Demo = func(input interface{}) { fmt.Println(input) }
	// 将 f8Demo 类型的变量赋值 i9Call 接口类型的变量
	var i9Call i9Call = f8Demo
	// 调用接口上的方法
	i9Call.f8Call("hello")
}
