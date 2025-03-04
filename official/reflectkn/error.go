package reflectkn

import "errors"

var (
	ErrMustInt        = errors.New("input must be int")
	ErrMustIntPointer = errors.New("input must be int pointer")
	ErrMustString     = errors.New("input must be string")

	ErrMustArray = errors.New("input must be array")
	ErrMustSlice = errors.New("input must be slice")
	ErrMustMap   = errors.New("input must be map")

	ErrMustStruct           = errors.New("input must be struct")
	ErrMustStructPointer    = errors.New("input must be struct pointer")
	ErrMustStructOrPointer  = errors.New("input must be struct or pointer")
	ErrStructFieldNotFound  = errors.New("struct field not found")
	ErrStructFieldCannotSet = errors.New("struct field cannot set")
)
