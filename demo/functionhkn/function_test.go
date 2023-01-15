package functionhkn

import (
	"fmt"
	"testing"
)

func TestFunctionValue(p7s6t *testing.T) {
	f8Chain()
}

func TestFunction(p7s6t *testing.T) {
	// 将匿名方法赋值给 f8do 类型的变量
	var f8do f8Demo = func(input interface{}) { fmt.Println(input) }
	// 调用方法
	f8do("hello")
	// 调用方法的方法
	f8do.f8Call("hello2")
	f8do.f8CallV2("hello2")
}
