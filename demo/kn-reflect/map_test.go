package kn_reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateMap(t *testing.T) {
	slicase := []struct {
		name       string
		input      any
		wantSliKey []any
		wantSliVal []any
		wantErr    error
	}{
		{
			// 非法输入 nil
			name:       "nil",
			input:      nil,
			wantSliKey: nil,
			wantSliVal: nil,
			wantErr:    ErrMustMap,
		},
		{
			// 空 map
			name:       "empty map",
			input:      map[string]string{},
			wantSliKey: []any{},
			wantSliVal: []any{},
			wantErr:    nil,
		},
		{
			// 普通 map
			name:       "normal map",
			input:      map[string]string{"key1": "val1", "key2": "val2"},
			wantSliKey: []any{"key1", "key2"},
			wantSliVal: []any{"val1", "val2"},
			wantErr:    nil,
		},
	}

	for _, tc := range slicase {
		t.Run(tc.name, func(t *testing.T) {
			slikey, slival, err := IterateMapV1(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSliKey, slikey)
			assert.Equal(t, tc.wantSliVal, slival)

			slikey, slival, err = IterateMapV2(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSliKey, slikey)
			assert.Equal(t, tc.wantSliVal, slival)
		})
	}
}
