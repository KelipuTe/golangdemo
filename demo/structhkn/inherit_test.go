package structhkn

import (
	"fmt"
	"testing"
)

// 多重继承
func Test_Inherit(t *testing.T) {
	c := new(Cat)
	fmt.Println("Cat: ")
	c.Walk()

	b := new(Owl)
	fmt.Println("Owl: ")
	b.Fly()
	b.Walk()
}

// 能行走
type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("calk")
}

// 能飞行
type Flyable struct{}

func (f *Flyable) Fly() {
	fmt.Println("fly")
}

// 猫
type Cat struct {
	Walkable // 能行走
}

// 猫头鹰
type Owl struct {
	Walkable // 能行走
	Flyable  // 能飞行
}
