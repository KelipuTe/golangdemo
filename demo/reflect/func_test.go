package reflect

import (
	"fmt"
	"testing"
)

func TestCallByString(t *testing.T) {
	CallByString("TestFuncV1")
	CallByString("TestFuncV2")
}

func TestCallByReflect(t *testing.T) {
	CallByReflect("TestFuncV3", 111)

	slires, err := CallByReflect("TestFuncV4", 222)
	fmt.Println(slires, err)
	for _, item := range slires {
		fmt.Println(item.Interface())
	}

	CallByReflect("TestFuncV5", "aaa", "bbb")
}
