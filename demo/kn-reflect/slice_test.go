package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateSlice(p7tt *testing.T) {
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
			errWant: ErrMustSlice,
		},
		{
			// 切片
			name:    "slice",
			input:   []int{1, 2, 3},
			resWant: []any{1, 2, 3},
			errWant: nil,
		},
		{
			// 二维切片
			name:    "slice_multiple",
			input:   [][]int{{1}, {11, 22}, {111, 222, 333}},
			resWant: []any{[]int{1}, []int{11, 22}, []int{111, 222, 333}},
			errWant: nil,
		},
	}

	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			res, err := IterateSlice(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.resWant, res)
		})
	}
}
