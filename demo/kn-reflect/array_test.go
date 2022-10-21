package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateArray(p7tt *testing.T) {
	s5case := []struct {
		name    string
		input   any
		resWant []any
		errWant error
	}{
		{
			// 非法输入 nil
			name:    "nil",
			input:   nil,
			resWant: nil,
			errWant: ErrMustArray,
		},
		{
			// 一维数组
			name:    "array",
			input:   [3]int{1, 2, 3},
			resWant: []any{1, 2, 3},
			errWant: nil,
		},
		{
			// 二维数组
			name:    "array_multiple",
			input:   [2][2]int{{1, 2}, {3, 4}},
			resWant: []any{[2]int{1, 2}, [2]int{3, 4}},
			errWant: nil,
		},
	}

	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			res, err := IterateArray(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.resWant, res)
		})
	}
}
