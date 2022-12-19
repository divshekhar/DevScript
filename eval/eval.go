package eval

import (
	"devscript/ast"
	"devscript/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	ZERO  = &object.Integer{Value: 0}
)

// Eval evaluates an AST
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		// Zero is a special case
		if node.Value == 0 {
			return ZERO
		}
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}

	return NULL
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return NULL
	default:
		return NULL
	}
}

// evaluates an integer infix expression
//
//	5 + 5;		// 10
//	5 - 5;		// 0
//	5 * 5;		// 25
//	5 / 5;		// 1
//	5 < 5;		// false
//	5 > 5;		// false
//	5 == 5;		// true
//	5 != 5;		// false
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return evalEqualToExpression(left, right)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

// convert integer 0 to Boolean false value,
// and all other integers Boolean true value
func convertIntegerToBoolean(obj object.Object) object.Object {
	if obj == ZERO {
		return FALSE
	}

	return TRUE
}

// evaluates an equal to expression
func evalEqualToExpression(left, right object.Object) object.Object {
	switch {
	// If left and right are null
	case left == NULL && right == NULL:
		return TRUE
	// If left is null and right is not
	case left == NULL:
		return FALSE
	default:
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.Integer).Value
		return nativeBoolToBooleanObject(leftVal == rightVal)
	}
}

// evalPrefixExpression evaluates a prefix expression
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

// evaluates a bang operator expression
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	case ZERO:
		return TRUE
	default:
		return FALSE
	}
}

// evaluates a minus prefix operator expression
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// nativeBoolToBooleanObject converts a native bool to a Boolean object
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// evalStatements evaluates a list of statements
func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}
