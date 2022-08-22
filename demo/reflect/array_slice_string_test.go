package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateArraySliceString(t *testing.T) {
	slicase := []struct {
		name    string
		input   any
		wantRes []any
		wantErr error
	}{
		{
			// 非法输入 nil
			name:    "nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustArraySliceString,
		},
		{
			// 一维数组
			name:    "array",
			input:   [3]int{1, 2, 3},
			wantRes: []any{1, 2, 3},
			wantErr: nil,
		},
		{
			// 二维数组
			name:    "array multiple",
			input:   [2][2]int{{1, 2}, {3, 4}},
			wantRes: []any{[2]int{1, 2}, [2]int{3, 4}},
			wantErr: nil,
		},
		{
			// 切片
			name:    "slice",
			input:   []int{1, 2, 3},
			wantRes: []any{1, 2, 3},
			wantErr: nil,
		},
		{
			// 字符串
			name:    "string",
			input:   "123",
			wantRes: []any{uint8('1'), uint8('2'), uint8('3')},
			wantErr: nil,
		},
	}

	for _, tc := range slicase {
		t.Run(tc.name, func(t *testing.T) {
			mapres, err := IterateArraySliceString(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, mapres)
		})
	}
}
