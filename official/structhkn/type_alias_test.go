package structhkn

import (
	"fmt"
	"testing"
)

// 新类型和别名的区别
func Test_TypeAlias(t *testing.T) {
	a1type := A1Type{}
	a1type.A1TypePrintln()

	// BType 不能直接调用 AType 的方法
	b1type := B1Type{}
	//b1type.A1TypePrintln()

	// 将 BType 转换为 AType，就可以调用
	b2type := A1Type(b1type)
	b2type.A1TypePrintln()

	// A1Alias 可以直接调用 AType 的方法
	a1alias := A1Alias{}
	a1alias.A1TypePrintln()
}

type A1Type struct {
}

func (t A1Type) A1TypePrintln() {
	fmt.Println("AType")
}

// 定义了一个新类型
type B1Type A1Type

// 给 A1Type 类型起了个别名
type A1Alias = A1Type
