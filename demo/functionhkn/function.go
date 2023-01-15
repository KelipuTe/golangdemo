package functionhkn

import "fmt"

// 在 Golang 里面，方法可以保存在变量中。

func f8Chain() {
	t4int := 1
	t4int2 := 2

	// 方法直接保存在变量中
	f8add1 := func(input int) int { return input + 1 }

	// 想要做成链式调用的形式，需要一组输入和输出相同的方法
	f8times2 := func(input int) int { return input * 2 }
	a5f8func := []func(int) int{f8add1, f8times2}

	// 直接调用保存在变量中的方法
	fmt.Println("call f8add1", f8add1(t4int))

	// 链式调用，把前一个方法的输出，作为后一个方法的输入
	for _, t4f8func := range a5f8func {
		t4int2 = t4f8func(t4int2)
	}
	fmt.Println("call chain", t4int2)
}

// 在 Golang 里面，可以给方法类型添加方法。
// 所以一个方法类型的变量，一方面作为方法类型，它本身可以直接调用。
// 另一方面作为添加了方法的方法类型，它可以调用添加到它上面去的方法。

type f8Demo func(interface{})

func (f8this f8Demo) f8Call(input interface{}) {
	fmt.Println("before f8Call f8Demo")
	f8this(input)
	fmt.Println("after f8Call f8Demo")
}

func (f8this f8Demo) f8CallV2(input interface{}) {
	fmt.Println("before f8CallV2 f8Demo")
	f8this(input)
	fmt.Println("after f8CallV2 f8Demo")
}
