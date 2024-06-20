package funchkn

//回调函数样例

import (
	"log"
	"testing"
)

func TestCallback(t *testing.T) {
	mul10 := func(n int) int {
		log.Println("mul10")
		return n * 10
	}
	cbMul10 := callbackExample(mul10)

	log.Println("cbLogBefore")
	cbLogBefore := cbMul10.logBefore()
	log.Println(cbLogBefore(10))

	log.Println("cbLogAfter")
	cbLogAfter := cbMul10.logAfter()
	log.Println(cbLogAfter(10))
}

type callbackExample func(int) int

func (cb callbackExample) logBefore() callbackExample {
	return func(n int) int {
		log.Println("logBefore")
		result := cb(n)
		return result
	}
}

func (cb callbackExample) logAfter() callbackExample {
	return func(n int) int {
		result := cb(n)
		log.Println("logAfter")
		return result
	}
}
