package unsafe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructIntFieldAccessorGetIntField(t *testing.T) {
	s5case := []struct {
		name    string
		entity  interface{}
		field   string
		wantVal int
		wantErr error
	}{
		{
			name:    "invalid field",
			entity:  &User{Sex: 1},
			field:   "Sex2",
			wantVal: 0,
			wantErr: ErrFieldNotFound,
		},
		{
			name:    "normal case",
			entity:  &User{Sex: 1},
			field:   "Sex",
			wantVal: 1,
			wantErr: nil,
		},
	}

	for _, tc := range s5case {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewStructUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			val, err := accessor.GetIntField(tc.field)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestStructIntFieldAccessorSetIntField(t *testing.T) {
	s5case := []struct {
		name    string
		entity  *User
		field   string
		newVal  int
		wantErr error
	}{
		{
			name:    "normal case",
			entity:  &User{},
			field:   "Sex",
			newVal:  1,
			wantErr: nil,
		},
	}

	for _, tc := range s5case {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewStructUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			err = accessor.SetIntField(tc.field, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.newVal, tc.entity.Sex)

		})
	}
}

func TestStructIntFieldAccessorGetAnyField(t *testing.T) {
	s5case := []struct {
		name    string
		entity  interface{}
		field   string
		wantVal any
		wantErr error
	}{
		{
			name:    "normal case string",
			entity:  &User{Name: "aaa"},
			field:   "Name",
			wantVal: "aaa",
			wantErr: nil,
		},
		{
			name:    "normal case int",
			entity:  &User{Sex: 1},
			field:   "Sex",
			wantVal: 1,
			wantErr: nil,
		},
	}

	for _, tc := range s5case {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewStructUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			val, err := accessor.GetAnyField(tc.field)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestStructIntFieldAccessorSetAnyField(t *testing.T) {
	s5case := []struct {
		name    string
		entity  *User
		field   string
		newVal  any
		wantErr error
	}{
		{
			name:    "normal case string",
			entity:  &User{Name: "aaa"},
			field:   "Name",
			newVal:  "bbb",
			wantErr: nil,
		},
		{
			name:    "normal case int",
			entity:  &User{Sex: 0},
			field:   "Sex",
			newVal:  1,
			wantErr: nil,
		},
	}

	for _, tc := range s5case {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewStructUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			err = accessor.SetAnyField(tc.field, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			if "Name" == tc.field {
				assert.Equal(t, tc.newVal, tc.entity.Name)
			} else if "Sex" == tc.field {
				assert.Equal(t, tc.newVal, tc.entity.Sex)
			}
		})
	}
}
