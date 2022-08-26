package unsafe

import "errors"

var ErrMustStruct = errors.New("input must be struct")

var ErrFieldNotFound = errors.New("field not found")
var ErrFieldCannotSet = errors.New("field cannot set")
