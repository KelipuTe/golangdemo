package functionhkn

import (
	"errors"
	"fmt"
)

// 在 Golang 里，满足下面任意一个条件的方法称为高阶方法。
// 1、接受其他方法作为传入参数。2、把其他方法作为返回参数。

type f8Operate func(x, y int) int

// 加法
func f8GetAddOperate() f8Operate {
	return func(x, y int) int {
		return x + y
	}
}

// 乘法
func f8GetMultiplyOperate() f8Operate {
	return func(x, y int) int {
		return x * y
	}
}

// 调用传入的方法处理传入的两个输入
func f8DoCalculate(f8oe f8Operate, x, y int) (int, error) {
	if nil == f8oe {
		return 0, errors.New("f8oe is nil")
	}
	return f8oe(x, y), nil
}

// 高阶方法可以实现用于实现闭包。
// 在闭包的内部可以修改从外部引用的变量。
// 闭包存在记忆效应，被捕获到闭包中的变量，会跟随闭包生命周期一直存在。
// 需要注意的是，闭包会影响程序的稳定和安全。

func f8EditOutside() {
	t4int := 1
	f8func := func() { t4int = 2 }

	fmt.Println("before edit", t4int)
	f8func()
	fmt.Println("after edit", t4int)
}

func f8Remember() {
	t4int := 1

	// 这里生成闭包方法时，f8func 里的 input 是 1
	f8func := f8Add1(t4int)
	// 这里对 t4int 进行修改时，不会影响 f8func 里的 input
	t4int = 10
	// 这里生成闭包方法时，f8func 里的 input 是 10
	f8funcV2 := f8Add1(t4int)
	// 这里对 t4int 进行修改时，不会影响 f8func 和 f8funcV2 里的 input
	t4int = 100

	fmt.Println(f8func())
	fmt.Println(f8funcV2())
}

func f8Add1(input int) func() int {
	return func() int { return input + 1 }
}
