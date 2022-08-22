package reflect

import "reflect"

func IterateMapV1(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	t1type := reflect.TypeOf(input)
	t1val := reflect.ValueOf(input)

	if t1type.Kind() != reflect.Map {
		return nil, nil, ErrMustMap
	}

	// 有几个元素
	t1len := t1val.Len()
	slikey := make([]any, 0, t1len)
	slival := make([]any, 0, t1len)
	// 遍历 key
	for _, k := range t1val.MapKeys() {
		slikey = append(slikey, k.Interface())
		// 获取 key 对应的 val
		v := t1val.MapIndex(k)
		slival = append(slival, v.Interface())
	}

	return slikey, slival, nil
}

func IterateMapV2(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	t1type := reflect.TypeOf(input)
	t1val := reflect.ValueOf(input)

	if t1type.Kind() != reflect.Map {
		return nil, nil, ErrMustMap
	}

	// 有几个元素
	t1len := t1val.Len()
	slikey := make([]any, 0, t1len)
	slival := make([]any, 0, t1len)
	// 遍历 map
	t1range := t1val.MapRange()
	for t1range.Next() {
		slikey = append(slikey, t1range.Key().Interface())
		slival = append(slival, t1range.Value().Interface())
	}

	return slikey, slival, nil
}
