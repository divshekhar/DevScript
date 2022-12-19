package eval

import (
	"devscript/ast"
	"devscript/object"
)

// Evaluate block statements
//
//	{
//	    statement1
//	    statement2
//	    ...
//	}
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			// Get the type of the result
			resultType := result.Type()

			// If the result is a return value or an error, return it
			// This is because we don't want to evaluate the rest of the statements
			// after a return statement or an error
			if resultType == object.RETURN_VALUE_OBJ || resultType == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}
