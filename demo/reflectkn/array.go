package reflectkn

import (
	"reflect"
)

// IterateArray 通过反射遍历数组
func IterateArray(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustArray
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	if inputType.Kind() != reflect.Map {
		return nil, ErrMustArray
	}

	if reflect.Array != inputType.Kind() {
		return nil, ErrMustArray
	}

	// 数组有几个元素
	a5len := inputValue.Len()
	s5res := make([]any, 0, a5len)
	// 按下标遍历元素
	for i := 0; i < a5len; i++ {
		t4item := inputValue.Index(i)
		// 如果是二维数组的话，这里会拿到二维数组的元素，也就是一维数组
		s5res = append(s5res, t4item.Interface())
	}

	return s5res, nil
}
