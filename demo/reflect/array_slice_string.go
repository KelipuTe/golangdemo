package reflect

import (
	"reflect"
)

// IterateArraySliceString 通过反射遍历数组，切片，字符串
func IterateArraySliceString(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustArraySliceString
	}

	t1type := reflect.TypeOf(input)
	t1val := reflect.ValueOf(input)
	t1kind := t1type.Kind()

	if t1kind != reflect.Array && t1kind != reflect.Slice && t1kind != reflect.String {
		return nil, ErrMustArraySliceString
	}

	// 有几个元素
	t1len := t1val.Len()
	slires := make([]any, 0, t1len)
	// 遍历元素
	for i := 0; i < t1len; i++ {
		t1item := t1val.Index(i)
		// 如果 input 是二维数组的话，这里会拿到二维数组的元素，也就是一维数组
		slires = append(slires, t1item.Interface())
	}

	return slires, nil
}
