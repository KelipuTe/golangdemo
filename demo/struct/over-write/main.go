package main

import "fmt"

func main() {
	// Son 的 SayHello() 会输出 Parent。因为调用的方法是 Parent 的，
	// 而 Parent 的 SayHello() 会调用 Parent 的 Name()，不会调用 Son 的。
	son := Son{
		Parent{},
	}
	son.SayHello()

	// Son2 的 SayHello() 会输出 Son2。
	// 因为 Sun2 的 SayHello() 屏蔽了 Parent 的 SayHello()。
	son2 := Son2{
		Parent{},
	}
	son2.SayHello()
}

type Parent struct {
}

func (p Parent) Name() string {
	return "Parent"
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

type Son struct {
	Parent
}

// Son 定义了自己的 Name() 方法
func (s Son) Name() string {
	return "Son"
}

type Son2 struct {
	Parent
}

// Son2 定义了自己的 Name() 方法
func (s Son2) Name() string {
	return "Son2"
}

// Son2 定义了自己的 SayHello() 方法
func (s Son2) SayHello() {
	fmt.Println("I am " + s.Name())
}
