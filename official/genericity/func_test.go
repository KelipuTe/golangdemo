package genericity

//泛型函数示例

import (
	"fmt"
	"testing"
)

func PrintAny[T any](value T) {
	fmt.Println(value)
}

func TestAnyFunc(t *testing.T) {
	PrintAny(1)

	PrintAny("1")
}
