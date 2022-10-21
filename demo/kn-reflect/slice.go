package kn_reflect

import (
	"errors"
	"reflect"
)

var ErrMustSlice = errors.New("input must be slice")

// IterateSlice 通过反射遍历切片
func IterateSlice(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustSlice
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	if reflect.Slice != inputType.Kind() {
		return nil, ErrMustSlice
	}

	// 切片有几个元素
	a5len := inputValue.Len()
	s5res := make([]any, 0, a5len)
	// 按下标遍历元素
	for i := 0; i < a5len; i++ {
		t4item := inputValue.Index(i)
		// 如果是二维切片的话，这里会拿到二维切片的元素，也就是一维切片
		s5res = append(s5res, t4item.Interface())
	}

	return s5res, nil
}
