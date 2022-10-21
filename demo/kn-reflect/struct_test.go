package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateStructField(t *testing.T) {
	slicase := []struct {
		name    string
		input   any
		wantRes map[string]any
		wantErr error
	}{
		{
			// 非法输入，nil
			name:    "nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustStruct,
		},
		{
			// 非法输入，数组
			name:    "array",
			input:   [2]int{1, 2},
			wantRes: nil,
			wantErr: ErrMustStruct,
		},
		{
			// 非法输入，切片
			name:    "slice",
			input:   []int{1, 2},
			wantRes: nil,
			wantErr: ErrMustStruct,
		},
		{
			// 非法输入，map
			name:    "map",
			input:   map[string]string{"a": "aa", "b": "bb"},
			wantRes: nil,
			wantErr: ErrMustStruct,
		},
		{
			// 非法输入，int 指针
			name: "int pointer",
			input: func() *int {
				i := 1
				return &i
			}(),
			wantRes: nil,
			wantErr: ErrMustStruct,
		},
		{
			// 普通结构体
			name:    "normal struct",
			input:   User{Name: "aaa", Sex: 1, age: 18},
			wantRes: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			wantErr: nil,
		},
		{
			// 一级结构体指针
			name:    "struct pointer",
			input:   &User{Name: "aaa", Sex: 1, age: 18},
			wantRes: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			wantErr: nil,
		},
		{
			// 二级结构体指针
			name: "struct pointer multiple",
			input: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			wantRes: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			wantErr: nil,
		},
	}
	for _, icase := range slicase {
		t.Run(icase.name, func(t *testing.T) {
			mapres, err := IterateStructField(icase.input)
			assert.Equal(t, icase.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, icase.wantRes, mapres)
		})
	}
}

func TestSetStructField(t *testing.T) {
	slicase := []struct {
		name     string
		instance any
		field    string
		val      any
		wantErr  error
	}{
		{
			// 非法输入，nil
			name:     "nil",
			instance: nil,
			field:    "Name",
			val:      "bbb",
			wantErr:  ErrMustStructPointer,
		},
		{
			// 非法输入，普通结构体
			name:     "struct",
			instance: User{Name: "aaa", Sex: 1, age: 18},
			field:    "Name",
			val:      "bbb",
			wantErr:  ErrMustStructPointer,
		},
		{
			// 非法输入，二级结构体指针
			name: "struct pointer multiple",
			instance: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			field:   "Name",
			val:     "bbb",
			wantErr: ErrMustStructPointer,
		},
		{
			// 一级结构体指针
			name:     "struct pointer",
			instance: &User{Name: "aaa", Sex: 1, age: 18},
			field:    "Name",
			val:      "bbb",
			wantErr:  nil,
		},
	}

	for _, icase := range slicase {
		t.Run(icase.name, func(t *testing.T) {
			err := SetStructField(icase.instance, icase.field, icase.val)
			assert.Equal(t, icase.wantErr, err)
			if err != nil {
				return
			}
		})
	}
}
