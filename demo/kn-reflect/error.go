package kn_reflect

import "errors"

var ErrMustArraySliceString = errors.New("input must be array, slice, string")
var ErrMustMap = errors.New("input must be map")

var ErrMustStruct = errors.New("input must be struct")
var ErrMustStructPointer = errors.New("input must be struct pointer")
var ErrMustStructOrStructPointer = errors.New("input must be struct or struct pointer")

var ErrFieldNotFound = errors.New("field not found")
var ErrFieldCannotSet = errors.New("field cannot set")
