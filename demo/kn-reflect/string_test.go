package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateString(p7tt *testing.T) {
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
			errWant: ErrMustString,
		},
		{
			// 字符串
			name:    "string",
			input:   "123",
			resWant: []any{uint8('1'), uint8('2'), uint8('3')},
			errWant: nil,
		},
	}

	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			res, err := IterateString(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.resWant, res)
		})
	}
}
