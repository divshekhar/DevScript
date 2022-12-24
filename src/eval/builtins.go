package eval

import (
	"devscript/src/object"
)

var builtins = map[string]*object.Builtin{
	"len":     {Function: lenFunction},
	"print":   {Function: printFunction},
	"println": {Function: printlnFunction},
}

// lenFunction returns the length of a string
func lenFunction(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

// Print function prints the value of the object
func printFunction(args ...object.Object) object.Object {
	for _, arg := range args {
		print(arg.Inspect() + " ")
	}
	return nil
}

// Println function prints the value of the object with a newline
func printlnFunction(args ...object.Object) object.Object {
	for _, arg := range args {
		println(arg.Inspect())
	}
	return nil
}
