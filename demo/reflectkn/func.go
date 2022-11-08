package reflectkn

import (
	"errors"
	"fmt"
	"reflect"
)

func TestFuncV1() {
	fmt.Println("TestFuncV1")
}

func TestFuncV2() {
	fmt.Println("TestFuncV2")
}

// CallByString 这种方式只能处理没有输入和输出的方法
func CallByString(name string) {
	mapfunc := map[string]func(){
		"TestFuncV1": TestFuncV1,
		"TestFuncV2": TestFuncV2,
	}
	t1func := mapfunc[name]
	t1func()
}

func TestFuncV3(input int) {
	fmt.Println("TestFuncV3", input)
}

func TestFuncV4(input int) int {
	fmt.Println("TestFuncV4", input)
	return input
}

func TestFuncV5(input1 string, input2 string) {
	fmt.Println("TestFuncV5", input1, input2)
}

// CallByReflect 这种方式可以处理有输入和输出的方法
func CallByReflect(name string, sliparam ...any) ([]reflect.Value, error) {
	mapfunc := map[string]any{
		"TestFuncV3": TestFuncV3,
		"TestFuncV4": TestFuncV4,
		"TestFuncV5": TestFuncV5,
	}

	t1type := reflect.TypeOf(mapfunc[name])
	t1val := reflect.ValueOf(mapfunc[name])

	t1num := t1type.NumIn()
	if t1num != len(sliparam) {
		return nil, errors.New("The number of params is not adapted.")
	}

	t1sliInput := make([]reflect.Value, t1num)
	for index, item := range sliparam {
		t1sliInput[index] = reflect.ValueOf(item)
	}

	t1sliOutput := make([]reflect.Value, 0)
	t1sliOutput = t1val.Call(t1sliInput)

	return t1sliOutput, nil
}
