package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//通过反射访问 int 和 int 指针

func visitInt(in any) (int, error) {
	if in == nil {
		return 0, ErrMustInt
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	if irt.Kind() != reflect.Int {
		return 0, ErrMustInt
	}

	anyv := irv.Interface()
	intv, ok := anyv.(int)
	if !ok {
		return 0, ErrMustInt
	}

	return intv, nil
}

func TestVisitInt(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes int
		wantErr error
	}{
		{
			name:    "非法nil",
			input:   nil,
			wantRes: 0,
			wantErr: ErrMustInt,
		},
		{
			name:    "非法string",
			input:   "abc",
			wantRes: 0,
			wantErr: ErrMustInt,
		},
		{
			name:    "合法int",
			input:   2,
			wantRes: 2,
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitInt(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}

func visitIntPointer(in any) (int, error) {
	if in == nil {
		return 0, ErrMustIntPointer
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	// 处理结构体指针（一级或多级指针）
	for irt.Kind() == reflect.Pointer {
		irt = irt.Elem()
		irv = irv.Elem()
	}

	if irt.Kind() != reflect.Int {
		return 0, ErrMustIntPointer
	}

	anyv := irv.Interface()
	intv, ok := anyv.(int)
	if !ok {
		return 0, ErrMustInt
	}

	return intv, nil
}

func TestVisitIntPointer(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes int
		wantErr error
	}{
		{
			name:    "非法nil",
			input:   nil,
			wantRes: 0,
			wantErr: ErrMustIntPointer,
		},
		{
			name:    "非法string",
			input:   "abc",
			wantRes: 0,
			wantErr: ErrMustIntPointer,
		},
		{
			name:    "合法int",
			input:   2,
			wantRes: 2,
			wantErr: nil,
		},
		{
			name: "合法*int",
			input: func() *int {
				input := 2
				return &input
			}(),
			wantRes: 2,
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitIntPointer(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}
