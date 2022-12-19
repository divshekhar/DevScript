package eval

import (
	"devscript/ast"
	"devscript/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	// Get the value of the identifier from the environment
	val, ok := env.Get(node.Value)

	if !ok {
		return newError("identifier not found: " + node.Value)
	}

	return val
}
