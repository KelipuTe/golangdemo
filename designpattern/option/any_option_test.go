package option

import (
	"log"
	"testing"
)

// option 设计模式，使用泛型

func TestAnyOption(t *testing.T) {
	o := &optionStruct{}
	applyOption(o, structWithString("a"), structWithInt(1))
	log.Println(o)

	o2 := &anyOptionStruct{}
	applyAnyOption(o2, anyStructWithString("b"), anyStructWithInt(2))
	log.Println(o2)
}

//##使用泛型可以防止这块代码重复写

type anyStructOption[T any] func(*T)

func applyAnyOption[T any](t *T, list ...anyStructOption[T]) {
	for _, v := range list {
		v(t)
	}
}

//##使用泛型可以防止这块代码重复写##

type anyOptionStruct struct {
	str string
	num int
}

func anyStructWithString(s string) anyStructOption[anyOptionStruct] {
	return func(o *anyOptionStruct) {
		o.str = s
	}
}

func anyStructWithInt(n int) anyStructOption[anyOptionStruct] {
	return func(o *anyOptionStruct) {
		o.num = n
	}
}
