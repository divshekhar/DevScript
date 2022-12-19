package eval

import (
	"devscript/ast"
	"devscript/object"
)

// Evaluate If Expressions
//
//	if (condition) {
//	    consequence
//	} else {
//	    alternative
//	}
func evalIfExpression(ifExpression *ast.IfExpression, env *object.Environment) object.Object {
	// Evaluate the condition
	condition := Eval(ifExpression.Condition, env)

	// If the condition is an error, return it
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		// If the condition is true, evaluate the consequence
		return Eval(ifExpression.Consequence, env)
	} else if ifExpression.Alternative != nil {
		// If the condition is false, evaluate the alternative
		return Eval(ifExpression.Alternative, env)
	} else {
		// If there is no alternative
		return NULL
	}
}

// returns false if the object is
//
//	NULL, FALSE, or ZERO
func isTruthy(condition object.Object) bool {
	switch condition {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	case ZERO:
		return false
	default:
		return true
	}
}
