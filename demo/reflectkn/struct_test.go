package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_f8IterateStructField(p7test *testing.T) {
	s5s6case := []struct {
		testName   string
		input      any
		resultWant map[string]any
		errWant    error
	}{
		{
			// 非法输入 nil
			testName:   "nil",
			input:      nil,
			resultWant: nil,
			errWant:    ErrMustStructOrStructPointer,
		},
		{
			// 非法输入 int 指针
			testName: "int pointer",
			input: func() *int {
				i := 1
				return &i
			}(),
			resultWant: nil,
			errWant:    ErrMustStructOrStructPointer,
		},
		{
			// 普通结构体
			testName:   "normal struct",
			input:      User{Name: "aaa", Sex: 1, age: 18},
			resultWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant:    nil,
		},
		{
			// 一级结构体指针
			testName:   "struct pointer",
			input:      &User{Name: "aaa", Sex: 1, age: 18},
			resultWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant:    nil,
		},
		{
			// 二级结构体指针
			testName: "struct pointer multiple",
			input: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			resultWant: map[string]any{"Name": "aaa", "Sex": 1, "age": 0},
			errWant:    nil,
		},
	}
	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateStructField(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}

func Test_f8SetStructField(p7test *testing.T) {
	s5s6case := []struct {
		testName  string
		input     any
		fieldWant string
		valueWant any
		errWant   error
	}{
		{
			// 非法输入 nil
			testName:  "nil",
			input:     nil,
			fieldWant: "Name",
			valueWant: "bbb",
			errWant:   ErrMustStructPointer,
		},
		{
			// 非法输入 普通结构体
			testName:  "struct",
			input:     User{Name: "aaa", Sex: 1, age: 18},
			fieldWant: "Name",
			valueWant: "bbb",
			errWant:   ErrMustStructPointer,
		},
		{
			// 非法输入 二级结构体指针
			testName: "struct_pointer_multiple",
			input: func() **User {
				p1u := &User{Name: "aaa", Sex: 1, age: 18}
				return &p1u
			}(),
			fieldWant: "Name",
			valueWant: "bbb",
			errWant:   ErrMustStructPointer,
		},
		{
			// 一级结构体指针
			testName:  "struct_pointer",
			input:     &User{Name: "aaa", Sex: 1, age: 18},
			fieldWant: "Name",
			valueWant: "bbb",
			errWant:   nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			err := f8SetStructField(t4case.input, t4case.fieldWant, t4case.valueWant)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
		})
	}
}

func Test_f8IterateStructFunc(p7test *testing.T) {
	s5s6case := []struct {
		testName   string
		input      any
		resultWant map[string]*S6FuncInfo
		errWant    error
	}{
		{
			// 非法输入，nil
			testName:   "nil",
			input:      nil,
			resultWant: nil,
			errWant:    ErrMustStructOrStructPointer,
		},
		{
			// 普通结构体
			testName: "normal struct",
			input:    User{Name: "aaa", Sex: 1, age: 18},
			resultWant: map[string]*S6FuncInfo{
				"GetName": {
					Name:          "GetName",
					S5InputType:   []reflect.Type{reflect.TypeOf(User{})},
					S5OutputType:  []reflect.Type{reflect.TypeOf("aaa")},
					S5OutputValue: []any{"aaa"},
				},
			},
			errWant: nil,
		},
		{
			// 一级结构体指针
			testName: "struct pointer",
			input:    &User{Name: "aaa", Sex: 1, age: 18},
			resultWant: map[string]*S6FuncInfo{
				"GetName": {
					Name:          "GetName",
					S5InputType:   []reflect.Type{reflect.TypeOf(&User{})},
					S5OutputType:  []reflect.Type{reflect.TypeOf("aaa")},
					S5OutputValue: []any{"aaa"},
				},
				"SetSex": {
					Name:          "SetSex",
					S5InputType:   []reflect.Type{reflect.TypeOf(&User{}), reflect.TypeOf(1)},
					S5OutputType:  []reflect.Type{reflect.TypeOf(0)},
					S5OutputValue: []any{0},
				},
			},
			errWant: nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateStructFunc(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}
