package structhkn

import (
	"fmt"
	"testing"
)

// 重写
func Test_Overwrite(t *testing.T) {
	// B1Overwrite 的 Println() 会输出 [A1]:A1。因为调用的是 A1Overwrite 的 Println()。
	// 而 A1Overwrite 的 Println() 会调用 A1Overwrite 的 Name()，不会调用 B1Overwrite 的 Name()。
	b1 := B1Overwrite{
		A1Overwrite{},
	}
	b1.Println()

	// B2Overwrite 的 Println() 会输出 [B2]:B2。
	// 因为 B2Overwrite 的 Println() 屏蔽了 A1Overwrite 的 Println()。
	b2 := B2Overwrite{
		A1Overwrite{},
	}
	b2.Println()
}

type A1Overwrite struct {
}

func (o A1Overwrite) Name() string {
	return "A1"
}

func (o A1Overwrite) Println() {
	fmt.Println("[A1]:" + o.Name())
}

// B1Overwrite 重写了 Name()，没有重写 Println()
type B1Overwrite struct {
	A1Overwrite
}

func (o B1Overwrite) Name() string {
	return "B1"
}

// B2Overwrite 重写了 Name()，重写了 Println()
type B2Overwrite struct {
	A1Overwrite
}

func (o B2Overwrite) Name() string {
	return "B2"
}

func (o B2Overwrite) Println() {
	fmt.Println("[B2]:" + o.Name())
}
