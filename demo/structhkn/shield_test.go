package structhkn

import (
	"fmt"
	"testing"
)

// 结构体多层嵌套时，嵌入层级越深的字段或方法越可能被屏蔽
func Test_Shield(t *testing.T) {
	a1 := A1Shield{"a1"}

	//A1Shield 不会被屏蔽。
	b1 := B1Shield{a1}
	fmt.Println(b1.Name)
	b1.Println()

	//A1Shield 会被屏蔽。
	b2 := B2Shield{"b2", a1}
	fmt.Println(b2.Name)
	b2.Println()
}

type A1Shield struct {
	Name string
}

func (t *A1Shield) Println() {
	fmt.Println("A1Shield")
}

type B1Shield struct {
	A1Shield
}

type B2Shield struct {
	Name string
	A1Shield
}

func (t *B2Shield) Println() {
	fmt.Println("B2Shield")
}
