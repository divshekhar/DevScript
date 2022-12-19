package eval

import (
	"devscript/object"
	"fmt"
)

// return a new error object.
//
// This is a helper function to make it easier to create new error objects.
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// check if an object is an error object.
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
