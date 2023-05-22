package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_f8IterateString(p7test *testing.T) {
	s5s6case := []struct {
		testName   string
		input      any
		resultWant []any
		errWant    error
	}{
		{
			// 非法输入 nil
			testName:   "nil",
			input:      nil,
			resultWant: nil,
			errWant:    ErrMustString,
		},
		{
			// 非法输入 int
			testName:   "int",
			input:      2,
			resultWant: nil,
			errWant:    ErrMustString,
		},
		{
			// 字符串
			testName:   "string",
			input:      "abc",
			resultWant: []any{uint8('a'), uint8('b'), uint8('c')},
			errWant:    nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateString(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}

// f8IterateString 通过反射遍历 string
func f8IterateString(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustString
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	if reflect.String != i9InputType.Kind() {
		return nil, ErrMustString
	}

	// 字符串有几个字符
	stringLen := s6InputValue.Len()
	s5EachChar := make([]any, 0, stringLen)
	// 按下标遍历字符
	for i := 0; i < stringLen; i++ {
		t4item := s6InputValue.Index(i)
		s5EachChar = append(s5EachChar, t4item.Interface())
	}

	return s5EachChar, nil
}
