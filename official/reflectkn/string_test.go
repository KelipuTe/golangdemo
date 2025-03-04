package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// 通过反射访问 string

func visitString(in any) ([]any, error) {
	if in == nil {
		return nil, ErrMustString
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	if irt.Kind() != reflect.String {
		return nil, ErrMustString
	}

	// 有几个字符，按下标遍历
	length := irv.Len()
	charList := make([]any, 0, length)
	for i := 0; i < length; i++ {
		v := irv.Index(i)
		charList = append(charList, v.Interface())
	}

	return charList, nil
}

func TestVisitString(t *testing.T) {
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
			wantErr: ErrMustString,
		},
		{
			name:    "非法int",
			input:   2,
			wantRes: nil,
			wantErr: ErrMustString,
		},
		{
			name:    "string",
			input:   "abc",
			wantRes: []any{uint8('a'), uint8('b'), uint8('v')},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(test *testing.T) {
			res, err := visitString(v.input)
			assert.Equal(test, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(test, v.wantRes, res)
		})
	}
}
