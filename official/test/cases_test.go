package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Add1(i int) int {
	return i + 1
}

func TestCases(t *testing.T) {
	testCases := []struct {
		name       string
		before     func(t *testing.T)
		after      func(t *testing.T)
		input      int
		wantOutput int
	}{
		{
			name: "测试01",
			before: func(t *testing.T) {
				t.Log("测试01before")
			},
			after: func(t *testing.T) {
				t.Log("测试01after")
			},
			input:      10,
			wantOutput: 11,
		},
		{
			name: "测试02",
			before: func(t *testing.T) {
				t.Log("测试02before")
			},
			after: func(t *testing.T) {
				t.Log("测试02after")
			},
			input:      20,
			wantOutput: 21,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			defer tc.after(t)

			output := Add1(tc.input)
			assert.Equal(t, tc.wantOutput, output)
		})
	}
}
