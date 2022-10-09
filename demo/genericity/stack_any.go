package genericity

import (
	"fmt"
	"reflect"
)

// StackAny 泛型栈
type StackAny[T any] struct {
	// s5value，栈的存储空间。这里的 []T 就是泛型，声明的时候可以传类型进来。
	// 如果是 var stack StackAny[int] 这么声明的，s5value 就是 []int
	// 如果是 var stack StackAny[string] 这么声明的，s5value 就是 []string
	// 如果是 var stack StackAny[any] 这么声明的，s5value 就是 []any
	s5value []T
}

// Init 初始化
func (p7this *StackAny[T]) Init() {
	if nil == p7this.s5value {
		// 泛型切片初始化的时候，make 的第一个参数是 []T
		p7this.s5value = make([]T, 0, 2)
	}
}

// Push 入栈
func (p7this *StackAny[T]) Push(value T) {
	if nil == p7this.s5value {
		return
	}

	p7this.s5value = append(p7this.s5value, value)
}

// Pop 出栈
func (p7this *StackAny[T]) Pop() (T, bool) {
	// 准备一个类型 T 对应的 0 值
	var zero T

	if nil == p7this.s5value {
		return zero, false
	}

	s5len := len(p7this.s5value)
	if 0 == s5len {
		return zero, false
	}

	top := p7this.s5value[s5len-1]
	p7this.s5value = p7this.s5value[:s5len-1]
	return top, true
}

// Contains 是否存在元素
func (p7this *StackAny[T]) Contains(value T) bool {
	for _, v := range p7this.s5value {
		// 这里会报错，这两个泛型变量不能直接进行比较。
		// 如果想直接比较，那么 StackAny 就不能声明为 StackAny[T any]，而是 StackAny[T comparable]。
		//if v == value {
		//	return true
		//}

		// 如果想进行比较，就需要先通过类型断言，然后在进行相应的处理。
		// 但是直接对泛型变量使用类型断言也是不可以的，下面两种方法都会报错。

		// 方法 1，直接断言
		//v2, ok := v.(int)

		// 方法 2，switch type 断言
		//switch v.(type) {
		//case int:
		//	fmt.Println("int")
		//}

		// 在对泛型变量使用类型断言时，需要先使用空接口 interface{} 做泛型约束。
		// 拿到类型之后就可以进行后续的比较操作了，但是可以预见，会很繁琐。
		var vi interface{} = v
		switch vi.(type) {
		case int:
			fmt.Println("type int")
		case string:
			fmt.Println("type string")
		}

		// 也可以使用反射的方式，获取泛型变量的类型信息。
		// 拿到类型之后就可以进行后续的比较操作了，但是可以预见，会很繁琐。
		t4type := reflect.TypeOf(v)
		t4kind := t4type.Kind()
		switch t4kind {
		case reflect.Int:
			fmt.Println("reflect int")
		case reflect.String:
			fmt.Println("reflect string")
		}
	}
	return false
}
