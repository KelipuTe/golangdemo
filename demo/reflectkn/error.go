package reflectkn

import "errors"

var ErrMustArray = errors.New("input must be array")

var ErrMustMap = errors.New("input must be map")

var ErrMustSlice = errors.New("input must be slice")

var ErrMustString = errors.New("input must be string")

var ErrMustStructPointer = errors.New("input must be struct pointer")
var ErrMustStructOrStructPointer = errors.New("input must be struct or struct pointer")

var ErrFieldNotFound = errors.New("field not found")
var ErrFieldCannotSet = errors.New("field cannot set")
