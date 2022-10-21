package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateStructField(p7tt *testing.T) {
	s5case := []struct {
		name    string
		input   any
		resWant map[string]any
		errWant error
	}{
		{
			// 非法输入，nil
			name:    "nil",
			input:   nil,
			resWant: nil,
			errWant: ErrMustStructOrStructPointer,
		},
		{
			// 非法输入，int 指针
			name: "int pointer",
			input: func() *int {
				i := 1
				return &i
			}(),
			resWant: nil,
			errWant: ErrMustStructOrStructPointer,
		},
		{
			// 普通结构体
			name:    "normal struct",
			input:   User{Name: "aaa", Sex: 1, age: 18},
			resWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant: nil,
		},
		{
			// 一级结构体指针
			name:    "struct pointer",
			input:   &User{Name: "aaa", Sex: 1, age: 18},
			resWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant: nil,
		},
		{
			// 二级结构体指针
			name: "struct pointer multiple",
			input: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			resWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant: nil,
		},
	}
	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			res, err := IterateStructField(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.resWant, res)
		})
	}
}

func TestSetStructField(p7tt *testing.T) {
	s5case := []struct {
		name    string
		input   any
		field   string
		value   any
		errWant error
	}{
		{
			// 非法输入，nil
			name:    "nil",
			input:   nil,
			field:   "Name",
			value:   "bbb",
			errWant: ErrMustStructPointer,
		},
		{
			// 非法输入，普通结构体
			name:    "struct",
			input:   User{Name: "aaa", Sex: 1, age: 18},
			field:   "Name",
			value:   "bbb",
			errWant: ErrMustStructPointer,
		},
		{
			// 非法输入，二级结构体指针
			name: "struct_pointer_multiple",
			input: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			field:   "Name",
			value:   "bbb",
			errWant: ErrMustStructPointer,
		},
		{
			// 一级结构体指针
			name:    "struct_pointer",
			input:   &User{Name: "aaa", Sex: 1, age: 18},
			field:   "Name",
			value:   "bbb",
			errWant: nil,
		},
	}

	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			err := SetStructField(t4case.input, t4case.field, t4case.value)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
		})
	}
}

func TestIterateStructFunc(p7tt *testing.T) {
	testCases := []struct {
		name    string
		input   any
		resWant map[string]*S6FuncInfo
		errWant error
	}{
		{
			// 非法输入，nil
			name:    "nil",
			input:   nil,
			resWant: nil,
			errWant: ErrMustStructOrStructPointer,
		},
		{
			// 普通结构体
			name:  "normal struct",
			input: User{Name: "aaa", Sex: 1, age: 18},
			resWant: map[string]*S6FuncInfo{
				"GetName": {
					Name:     "GetName",
					S5Input:  []reflect.Type{reflect.TypeOf(User{})},
					S5Output: []reflect.Type{reflect.TypeOf("aaa")},
					S5Res:    []any{"aaa"},
				},
			},
			errWant: nil,
		},
		{
			// 一级结构体指针
			name:  "struct pointer",
			input: &User{Name: "aaa", Sex: 1, age: 18},
			resWant: map[string]*S6FuncInfo{
				"GetName": {
					Name:     "GetName",
					S5Input:  []reflect.Type{reflect.TypeOf(&User{})},
					S5Output: []reflect.Type{reflect.TypeOf("aaa")},
					S5Res:    []any{"aaa"},
				},
				"SetSex": {
					Name:     "SetSex",
					S5Input:  []reflect.Type{reflect.TypeOf(&User{}), reflect.TypeOf(1)},
					S5Output: []reflect.Type{reflect.TypeOf(0)},
					S5Res:    []any{0},
				},
			},
			errWant: nil,
		},
	}

	for _, t4case := range testCases {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			res, err := IterateStructFunc(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.resWant, res)
		})
	}
}
