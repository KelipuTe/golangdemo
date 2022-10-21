package kn_reflect

import (
	"errors"
	"reflect"
)

var ErrMustString = errors.New("input must be string")

// IterateString 通过反射遍历切片
func IterateString(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustString
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	if reflect.String != inputType.Kind() {
		return nil, ErrMustString
	}

	// 字符串有几个字符
	a5len := inputValue.Len()
	s5res := make([]any, 0, a5len)
	// 按下标遍历字符
	for i := 0; i < a5len; i++ {
		t4item := inputValue.Index(i)
		s5res = append(s5res, t4item.Interface())
	}

	return s5res, nil
}
