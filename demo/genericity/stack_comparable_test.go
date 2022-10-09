package genericity

import (
	"fmt"
	"testing"
)

func TestStackComparableIntContains(t *testing.T) {
	var stack StackComparable[int]
	stack.Init()
	stack.Push(10)
	ok := stack.Contains(10)
	ok2 := stack.Contains(20)
	fmt.Println(ok, ok2)
}

func TestStackComparableStringContains(t *testing.T) {
	var stack StackComparable[string]
	stack.Init()
	stack.Push("aa")
	ok := stack.Contains("aa")
	ok2 := stack.Contains("bb")
	fmt.Println(ok, ok2)
}
