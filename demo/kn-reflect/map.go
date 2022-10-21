package kn_reflect

import (
	"errors"
	"reflect"
)

var ErrMustMap = errors.New("input must be map")

// IterateMap 通过反射遍历 map
func IterateMap(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	if reflect.Map != inputType.Kind() {
		return nil, nil, ErrMustMap
	}

	// map 有几个元素
	t4len := inputValue.Len()
	s5resKey := make([]any, 0, t4len)
	s5resValue := make([]any, 0, t4len)
	// 遍历 key
	for _, t4k := range inputValue.MapKeys() {
		s5resKey = append(s5resKey, t4k.Interface())
		// 获取 key 对应的 value
		t4v := inputValue.MapIndex(t4k)
		s5resValue = append(s5resValue, t4v.Interface())
	}

	return s5resKey, s5resValue, nil
}

func IterateMapV2(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	if reflect.Map != inputType.Kind() {
		return nil, nil, ErrMustMap
	}

	// map 有几个元素
	t4len := inputValue.Len()
	s5resKey := make([]any, 0, t4len)
	s5resValue := make([]any, 0, t4len)
	// 遍历 key
	t4MapRange := inputValue.MapRange()
	for t4MapRange.Next() {
		s5resKey = append(s5resKey, t4MapRange.Key().Interface())
		s5resValue = append(s5resValue, t4MapRange.Value().Interface())
	}

	return s5resKey, s5resValue, nil
}
