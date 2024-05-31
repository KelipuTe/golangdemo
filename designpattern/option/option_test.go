package option

import (
	"log"
	"testing"
)

// option 设计模式

func TestOption(t *testing.T) {
	o := &optionStruct{}
	applyOption(o, structWithString("a"), structWithInt(1))
	log.Println(o)
}

type structOption func(*optionStruct)

func applyOption(o *optionStruct, list ...structOption) {
	for _, v := range list {
		v(o)
	}
}

type optionStruct struct {
	str string
	num int
}

func structWithString(s string) structOption {
	return func(o *optionStruct) {
		o.str = s
	}
}

func structWithInt(n int) structOption {
	return func(o *optionStruct) {
		o.num = n
	}
}
