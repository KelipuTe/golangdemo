package funchkn

//链式调用样例

import (
	"log"
	"testing"
)

func TestChain(t *testing.T) {
	funcList := []chainExample{
		chainAdd10,
		chainMul10,
	}

	n := 1
	for _, v := range funcList {
		n = v(n)
		log.Println(n)
	}
}

type chainExample func(int) int

func chainAdd10(n int) int {
	log.Println("chainAdd10")
	return n + 10
}

func chainMul10(n int) int {
	log.Println("chainMul10")
	return n * 10
}
