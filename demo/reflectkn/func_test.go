package reflectkn

import (
	"errors"
	"log"
	"reflect"
	"testing"
)

//通过反射调用函数

func CaseFunc() {
	log.Println("CaseFunc")
}

func CaseFuncV2() {
	log.Println("CaseFuncV2")
}

// 这种方式只能处理没有输入和输出的方法
func CallByString(name string) {
	list := map[string]func(){
		"CaseFunc":   CaseFunc,
		"CaseFuncV2": CaseFuncV2,
	}
	f := list[name]
	f()
}

func TestCallByString(t *testing.T) {
	CallByString("CaseFunc")
	CallByString("CaseFuncV2")
}

func CaseFuncV3(in int) {
	log.Println("CaseFuncV3", in)
}

func CaseFuncV4(in int) int {
	log.Println("CaseFuncV4", in)
	return in
}

func CaseFuncV5(in1 string, in2 string) {
	log.Println("CaseFuncV5", in1, in2)
}

// 这种方式可以处理有输入和输出的方法
func CallByReflect(name string, paramList ...any) ([]reflect.Value, error) {
	list := map[string]any{
		"CaseFuncV3": CaseFuncV3,
		"CaseFuncV4": CaseFuncV4,
		"CaseFuncV5": CaseFuncV5,
	}

	frt := reflect.TypeOf(list[name])
	frv := reflect.ValueOf(list[name])

	paramNum := frt.NumIn()
	if paramNum != len(paramList) {
		return nil, errors.New("the number of params is not adapted")
	}

	inList := make([]reflect.Value, paramNum)
	for i, v := range paramList {
		inList[i] = reflect.ValueOf(v)
	}
	outList := make([]reflect.Value, 0)

	outList = frv.Call(inList)

	return outList, nil
}

func TestCallByReflect(t *testing.T) {
	_, _ = CallByReflect("CaseFuncV3", 111)

	outList, err := CallByReflect("CaseFuncV4", 222)
	log.Println(outList, err)
	for _, v := range outList {
		log.Println(v.Interface())
	}

	_, _ = CallByReflect("CaseFuncV5", "aaa", "bbb")
}
