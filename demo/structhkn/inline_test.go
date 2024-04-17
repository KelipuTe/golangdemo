package structhkn

import (
	"fmt"
	"testing"
)

// 内嵌结构体（匿名结构体）
// 内嵌结构体默认使用类型名作为字段名
func Test_Inline(t *testing.T) {
	//内嵌结构体的字段名没有冲突时，可以直接通过字段名来访问，也可以通过类型名+字段名来访问。
	b1 := B1Inline{
		A1Inline{"a1"},
	}
	fmt.Println(b1.name, b1.A1Inline.name)

	//内嵌结构体的字段名有冲突时，只能通过类型名+字段名来访问。
	//如果直接通过字段名来访问。编译时会报错。
	//编译器会告知选择器 name 有歧义，因为不知道是 A1Inline 还是 A2Inline 的。
	b2 := B2Inline{
		A1Inline{"a1"},
		A2Inline{"a2"},
	}
	fmt.Println(b2.A1Inline.name, b2.A2Inline.name)
}

type A1Inline struct {
	name string
}

type A2Inline struct {
	name string
}

type B1Inline struct {
	A1Inline
}

type B2Inline struct {
	A1Inline
	A2Inline
}
