package structhkn

import (
	"fmt"
	"testing"
)

// 结构体接收器和指针接收器的区别
// 结构体只能调用结构体接收器的方法，结构体指针可以调用所有的方法。
func Test_Receiver(t *testing.T) {
	a1 := A1Receiver{Name: "a1"}
	fmt.Printf("%+v\n", a1)
	a1.SetName("aa11")
	fmt.Printf("%+v\n", a1)

	a2p7 := &A2Receiver{Name: "a2"}
	fmt.Printf("%+v\r\n", a2p7)
	a2p7.SetName("aa22")
	fmt.Printf("%+v\r\n", a2p7)

	// 严格来讲，结构体只能调用结构体接收器的方法。但是，go会自动进行转译。
	// 这里的 a2.SetName("aa22") 会被转译为 (&a2).SetName("aa22")。
	a2 := A2Receiver{Name: "a2"}
	fmt.Printf("%+v\r\n", a2)
	a2.SetName("aa22")
	fmt.Printf("%+v\r\n", a2)
}

type A1Receiver struct {
	Name string
}

// 结构体接收器，值传递，相当于原数据的一个副本
func (t A1Receiver) SetName(name string) {
	t.Name = name
}

type A2Receiver struct {
	Name string
}

// 指针接收器，引用传递，通过指针可以直接修改原数据
func (t *A2Receiver) SetName(name string) {
	t.Name = name
}
