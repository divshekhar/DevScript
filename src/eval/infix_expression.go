package eval

import "devscript/src/object"

// Evaluates an infix expression
//
//	5 + 5;		// 10 (object.Integer)
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	// if both objects are integers, evaluate the infix expression
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		{
			return evalIntegerInfixExpression(operator, left, right)
		}

	// if both objects are strings, evaluate the infix expression
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		{
			return evalStringInfixExpression(operator, left, right)
		}

	case operator == "==":
		{
			return nativeBoolToBooleanObject(left == right)
		}
	case operator == "!=":
		{
			return nativeBoolToBooleanObject(left != right)
		}
	case left.Type() != right.Type():
		{
			return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
		}
	default:
		{
			return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
		}
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
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evaluates a string infix expression
//
//	"Hello" + "World";		// "HelloWorld"
//	"Hello" == "World";		// false
//	"Hello" != "World";		// true
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
