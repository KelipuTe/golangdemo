package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//通过反射遍历 map。这里示例了两种遍历方式。
//一种是通过遍历 key 然后通过 key 拿到 value。
//一种是通过迭代器，一次拿一个键值对，直到拿完。

func visitMap(in any) ([]any, []any, error) {
	if in == nil {
		return nil, nil, ErrMustMap
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	if irt.Kind() != reflect.Map {
		return nil, nil, ErrMustMap
	}

	// 有几个元素，遍历 key，然后通过 key 拿到 value
	length := irv.Len()
	kList := make([]any, 0, length)
	vList := make([]any, 0, length)
	for _, k := range irv.MapKeys() {
		kList = append(kList, k.Interface())
		v := irv.MapIndex(k)
		vList = append(vList, v.Interface())
	}

	return kList, vList, nil
}

func visitMapV2(in any) ([]any, []any, error) {
	if in == nil {
		return nil, nil, ErrMustMap
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	if irt.Kind() != reflect.Map {
		return nil, nil, ErrMustMap
	}

	// 有几个元素
	length := irv.Len()
	kList := make([]any, 0, length)
	vList := make([]any, 0, length)
	// 通过迭代器遍历键值对
	mapRange := irv.MapRange()
	for mapRange.Next() {
		k := mapRange.Key()
		kList = append(kList, k.Interface())
		v := mapRange.Value()
		vList = append(vList, v.Interface())
	}

	return kList, vList, nil
}

func TestVisitMap(t *testing.T) {
	caseList := []struct {
		name          string
		input         any
		wantKeyList   []any
		wantValueList []any
		wantErr       error
	}{
		{
			name:          "非法nil",
			input:         nil,
			wantKeyList:   nil,
			wantValueList: nil,
			wantErr:       ErrMustMap,
		},
		{
			name:          "空map",
			input:         map[string]string{},
			wantKeyList:   []any{},
			wantValueList: []any{},
			wantErr:       nil,
		},
		{
			name:          "普通map",
			input:         map[string]string{"key1": "val1", "key2": "val2"},
			wantKeyList:   []any{"key1", "key2"},
			wantValueList: []any{"val1", "val2"},
			wantErr:       nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			k2, v2, err := visitMap(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantKeyList, k2)
			assert.Equal(t, v.wantValueList, v2)

			k2, v2, err = visitMapV2(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantKeyList, k2)
			assert.Equal(t, v.wantValueList, v2)
		})
	}
}
