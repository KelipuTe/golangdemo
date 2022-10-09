package genericity

import (
	"fmt"
	"testing"
)

func TestStackAnyInt(t *testing.T) {
	var stack StackAny[int]
	stack.Init()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack)
}

func TestStackAnyString(t *testing.T) {
	var stack StackAny[string]
	stack.Init()
	stack.Push("aa")
	stack.Push("bb")
	stack.Push("cc")
	fmt.Println(stack)
}

func TestStackAnyAny(t *testing.T) {
	var stack StackAny[any]
	stack.Init()
	stack.Push(10)
	stack.Push("aa")
	stack.Push(20)
	stack.Push("bb")
	fmt.Println(stack)
}

func TestStackAnyAnyContains(t *testing.T) {
	var stack StackAny[any]
	stack.Init()
	stack.Push(10)
	stack.Push("aa")
	stack.Contains(10)
}
