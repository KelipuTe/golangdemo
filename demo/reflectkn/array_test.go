package reflectkn

import (
  "reflect"
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_f8IterateArray(p7test *testing.T) {
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
      errWant:    ErrMustArray,
    },
    {
      // 一维数组
      testName:   "array",
      input:      [3]int{1, 2, 3},
      resultWant: []any{1, 2, 3},
      errWant:    nil,
    },
    {
      // 二维数组
      testName:   "array_multiple",
      input:      [2][2]int{{1, 2}, {3, 4}},
      resultWant: []any{[2]int{1, 2}, [2]int{3, 4}},
      errWant:    nil,
    },
  }

  for _, t4case := range s5s6case {
    p7test.Run(t4case.testName, func(p7test *testing.T) {
      res, err := f8IterateArray(t4case.input)
      assert.Equal(p7test, t4case.errWant, err)
      if err != nil {
        return
      }
      assert.Equal(p7test, t4case.resultWant, res)
    })
  }
}

// f8IterateArray 通过反射遍历数组
func f8IterateArray(input any) ([]any, error) {
  if nil == input {
    return nil, ErrMustArray
  }

  i9InputType := reflect.TypeOf(input)
  inputValue := reflect.ValueOf(input)

  if reflect.Array != i9InputType.Kind() {
    return nil, ErrMustArray
  }

  // 数组有几个元素
  arrayLen := inputValue.Len()
  s5EachItem := make([]any, 0, arrayLen)
  // 按下标遍历元素
  for i := 0; i < arrayLen; i++ {
    t4item := inputValue.Index(i)
    // 如果是二维数组的话，这里会拿到二维数组的元素，也就是一维数组
    s5EachItem = append(s5EachItem, t4item.Interface())
  }

  return s5EachItem, nil
}
