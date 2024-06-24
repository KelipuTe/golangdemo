package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//通过反射遍历切片

func visitSlice(in any) ([]any, error) {
	if in == nil {
		return nil, ErrMustSlice
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	if irt.Kind() != reflect.Slice {
		return nil, ErrMustSlice
	}

	// 有几个元素，按下标遍历
	length := irv.Len()
	itemList := make([]any, 0, length)
	for i := 0; i < length; i++ {
		v := irv.Index(i)
		// 如果是二维切片的话，这里会拿到二维切片的元素，也就是一维切片
		itemList = append(itemList, v.Interface())
	}

	return itemList, nil
}

func TestVisitSlice(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes []any
		wantErr error
	}{
		{
			name:    "非法nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustSlice,
		},
		{
			name:    "切片",
			input:   []int{1, 2, 3},
			wantRes: []any{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "二维切片",
			input:   [][]int{{1}, {11, 22}, {111, 222, 333}},
			wantRes: []any{[]int{1}, []int{11, 22}, []int{111, 222, 333}},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitSlice(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}
