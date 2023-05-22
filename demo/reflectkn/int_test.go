package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_f8IterateInt(p7test *testing.T) {
	s5s6case := []struct {
		testName   string
		input      any
		resultWant int
		errWant    error
	}{
		{
			// 非法输入 nil
			testName:   "nil",
			input:      nil,
			resultWant: 0,
			errWant:    ErrMustInt,
		},
		{
			// 非法输入 字符串
			testName:   "string",
			input:      "abc",
			resultWant: 0,
			errWant:    ErrMustInt,
		},
		{
			// int
			testName:   "int",
			input:      2,
			resultWant: 2,
			errWant:    nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateInt(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}

// f8IterateInt 通过反射遍历 int
func f8IterateInt(input any) (int, error) {
	if nil == input {
		return 0, ErrMustInt
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	if reflect.Int != i9InputType.Kind() {
		return 0, ErrMustInt
	}

	anyValue := s6InputValue.Interface()
	intValue, ok := anyValue.(int)
	if !ok {
		return 0, ErrMustInt
	}

	return intValue, nil
}

func Test_f8IterateIntPointer(p7test *testing.T) {
	s5s6case := []struct {
		testName   string
		input      any
		resultWant int
		errWant    error
	}{
		{
			// 非法输入 nil
			testName:   "nil",
			input:      nil,
			resultWant: 0,
			errWant:    ErrMustIntPointer,
		},
		{
			// 非法输入 字符串
			testName:   "string",
			input:      "abc",
			resultWant: 0,
			errWant:    ErrMustIntPointer,
		},
		{
			// int
			testName:   "int",
			input:      2,
			resultWant: 2,
			errWant:    nil,
		},
		{
			// int*
			testName: "*int",
			input: func() *int {
				input := 2
				return &input
			}(),
			resultWant: 2,
			errWant:    nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateIntPointer(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}

func f8IterateIntPointer(input any) (int, error) {
	if nil == input {
		return 0, ErrMustIntPointer
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	// 处理结构体指针（一级或多级指针）
	for reflect.Pointer == i9InputType.Kind() {
		i9InputType = i9InputType.Elem()
		s6InputValue = s6InputValue.Elem()
	}
	if reflect.Int != i9InputType.Kind() {
		return 0, ErrMustIntPointer
	}

	anyValue := s6InputValue.Interface()
	intValue, ok := anyValue.(int)
	if !ok {
		return 0, ErrMustInt
	}

	return intValue, nil
}
