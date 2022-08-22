package reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateStructFunc(t *testing.T) {
	testCases := []struct {
		name    string
		input   any
		wantRes map[string]*StructFuncInfo
		wantErr error
	}{
		{
			// 非法输入，nil
			name:    "nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustStructOrStructPointer,
		},
		{
			// 普通结构体
			name:  "normal struct",
			input: User{Name: "aaa", Sex: 1, age: 18},
			wantRes: map[string]*StructFuncInfo{
				"GetName": {
					Name:      "GetName",
					SliInput:  []reflect.Type{reflect.TypeOf(User{})},
					SliOutput: []reflect.Type{reflect.TypeOf("aaa")},
					SliRes:    []any{"aaa"},
				},
			},
			wantErr: nil,
		},
		{
			// 一级结构体指针
			name:  "struct pointer",
			input: &User{Name: "aaa", Sex: 1, age: 18},
			wantRes: map[string]*StructFuncInfo{
				"GetName": {
					Name:      "GetName",
					SliInput:  []reflect.Type{reflect.TypeOf(&User{})},
					SliOutput: []reflect.Type{reflect.TypeOf("aaa")},
					SliRes:    []any{"aaa"},
				},
				"SetSex": {
					Name:      "SetSex",
					SliInput:  []reflect.Type{reflect.TypeOf(&User{}), reflect.TypeOf(1)},
					SliOutput: []reflect.Type{reflect.TypeOf(0)},
					SliRes:    []any{0},
				},
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mapres, err := IterateStructFunc(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, mapres)
		})
	}
}
