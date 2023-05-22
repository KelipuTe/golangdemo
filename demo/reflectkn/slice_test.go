package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_f8IterateSlice(p7test *testing.T) {
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
			errWant:    ErrMustSlice,
		},
		{
			// 切片
			testName:   "slice",
			input:      []int{1, 2, 3},
			resultWant: []any{1, 2, 3},
			errWant:    nil,
		},
		{
			// 二维切片
			testName:   "slice_multiple",
			input:      [][]int{{1}, {11, 22}, {111, 222, 333}},
			resultWant: []any{[]int{1}, []int{11, 22}, []int{111, 222, 333}},
			errWant:    nil,
		},
	}

	for _, t4case := range s5s6case {
		p7test.Run(t4case.testName, func(p7test *testing.T) {
			res, err := f8IterateSlice(t4case.input)
			assert.Equal(p7test, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7test, t4case.resultWant, res)
		})
	}
}

// f8IterateSlice 通过反射遍历切片
func f8IterateSlice(input any) ([]any, error) {
	if nil == input {
		return nil, ErrMustSlice
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	if reflect.Slice != i9InputType.Kind() {
		return nil, ErrMustSlice
	}

	// 切片有几个元素
	sliceLen := s6InputValue.Len()
	s5EachItem := make([]any, 0, sliceLen)
	// 按下标遍历元素
	for i := 0; i < sliceLen; i++ {
		t4item := s6InputValue.Index(i)
		// 如果是二维切片的话，这里会拿到二维切片的元素，也就是一维切片
		s5EachItem = append(s5EachItem, t4item.Interface())
	}

	return s5EachItem, nil
}
