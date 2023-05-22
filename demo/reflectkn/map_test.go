package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_f8IterateMap(p7test *testing.T) {
	s5s6case := []struct {
		name        string
		input       any
		wantS5Key   []any
		wantS5Value []any
		wantErr     error
	}{
		{
			// 非法输入 nil
			name:        "nil",
			input:       nil,
			wantS5Key:   nil,
			wantS5Value: nil,
			wantErr:     ErrMustMap,
		},
		{
			// 空 map
			name:        "empty_map",
			input:       map[string]string{},
			wantS5Key:   []any{},
			wantS5Value: []any{},
			wantErr:     nil,
		},
		{
			// 普通 map
			name:        "normal_map",
			input:       map[string]string{"key1": "val1", "key2": "val2"},
			wantS5Key:   []any{"key1", "key2"},
			wantS5Value: []any{"val1", "val2"},
			wantErr:     nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.name, func(p7test *testing.T) {
			t4key, t4value, err := f8IterateMap(t4case.input)
			assert.Equal(p7test, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.wantS5Key, t4key)
			assert.Equal(p7test, t4case.wantS5Value, t4value)

			t4key, t4value, err = f8IterateMapV2(t4case.input)
			assert.Equal(p7test, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.wantS5Key, t4key)
			assert.Equal(p7test, t4case.wantS5Value, t4value)
		})
	}
}

// f8IterateMap 通过反射遍历 map
func f8IterateMap(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	if reflect.Map != i9InputType.Kind() {
		return nil, nil, ErrMustMap
	}

	// map 有几个元素
	mapSize := s6InputValue.Len()
	s5MapKey := make([]any, 0, mapSize)
	s5MapValue := make([]any, 0, mapSize)
	// 遍历 key
	for _, t4key := range s6InputValue.MapKeys() {
		s5MapKey = append(s5MapKey, t4key.Interface())
		// 获取 key 对应的 valueWant
		t4value := s6InputValue.MapIndex(t4key)
		s5MapValue = append(s5MapValue, t4value.Interface())
	}

	return s5MapKey, s5MapValue, nil
}

func f8IterateMapV2(input any) ([]any, []any, error) {
	if nil == input {
		return nil, nil, ErrMustMap
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	if reflect.Map != i9InputType.Kind() {
		return nil, nil, ErrMustMap
	}

	// map 有几个元素
	mapSize := s6InputValue.Len()
	s5MapKey := make([]any, 0, mapSize)
	s5MapValue := make([]any, 0, mapSize)
	// 遍历 key
	t4MapRange := s6InputValue.MapRange()
	for t4MapRange.Next() {
		s5MapKey = append(s5MapKey, t4MapRange.Key().Interface())
		s5MapValue = append(s5MapValue, t4MapRange.Value().Interface())
	}

	return s5MapKey, s5MapValue, nil
}
