package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//通过反射遍历数组

func visitArray(in any) ([]any, error) {
	if in == nil {
		return nil, ErrMustArray
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)
	if irt.Kind() != reflect.Array {
		return nil, ErrMustArray
	}

	// 有几个元素，按下标遍历
	length := irv.Len()
	itemList := make([]any, 0, length)
	for i := 0; i < length; i++ {
		v := irv.Index(i)
		// 如果是二维数组的话，这里会拿到二维数组的元素，也就是一维数组
		itemList = append(itemList, v.Interface())
	}

	return itemList, nil
}

func TestVisitArray(t *testing.T) {
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
			wantErr: ErrMustArray,
		},
		{
			name:    "合法一维数组",
			input:   [3]int{1, 2, 3},
			wantRes: []any{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "合法二维数组",
			input:   [2][2]int{{1, 2}, {3, 4}},
			wantRes: []any{[2]int{1, 2}, [2]int{3, 4}},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitArray(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}
