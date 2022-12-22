package eval

import (
	"devscript/src/ast"
	"devscript/src/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	// Get the value of the identifier from the environment
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	// Check if the identifier is a builtin function
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}
