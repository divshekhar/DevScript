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
		return evalProgram(node)
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
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}
	default:
		return NULL
	}
}

func evalProgram(node *ast.Program) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

// nativeBoolToBooleanObject converts a native bool to a Boolean object
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
