package eval

import (
	"devscript/ast"
	"devscript/object"
)

// evalStatements evaluates a list of statements
func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}
