package eval

import (
	"devscript/ast"
	"devscript/object"
)

func evalIfExpression(ifExpression *ast.IfExpression) object.Object {
	condition := Eval(ifExpression.Condition)
	if isTruthy(condition) {
		return Eval(ifExpression.Consequence)
	} else if ifExpression.Alternative != nil {
		return Eval(ifExpression.Alternative)
	} else {
		return NULL
	}
}

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
