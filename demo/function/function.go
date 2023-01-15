package interfacehkn

import "fmt"

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
