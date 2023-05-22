package reflectkn

import "errors"

var ErrMustInt = errors.New("input must be int")
var ErrMustIntPointer = errors.New("input must be int pointer")
var ErrMustString = errors.New("input must be string")

var ErrMustArray = errors.New("input must be array")
var ErrMustSlice = errors.New("input must be slice")
var ErrMustMap = errors.New("input must be map")

var ErrMustStruct = errors.New("input must be struct")
var ErrMustStructPointer = errors.New("input must be struct pointer")
var ErrMustStructOrStructPointer = errors.New("input must be struct or struct pointer")
var ErrStructFieldNotFound = errors.New("struct fieldWant not found")
var ErrStructFieldCannotSet = errors.New("struct fieldWant cannot set")
