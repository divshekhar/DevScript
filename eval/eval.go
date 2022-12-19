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
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Program is the root node of the AST
	case *ast.Program:
		return evalProgram(node, env)

	// Evaluate Variable Statements
	case *ast.VarStatement:
		{
			val := Eval(node.Value, env)
			if isError(val) {
				return val
			}
			env.Set(node.Name.Value, val)
			return val
		}

	// Evaluate Expressions
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// Evaluate Identifier
	case *ast.Identifier:
		return evalIdentifier(node, env)

	// Evaluate Integer Literals
	case *ast.IntegerLiteral:
		{
			// Zero is a special case
			if node.Value == 0 {
				return ZERO
			}
			return &object.Integer{Value: node.Value}
		}

	// Evaluate Boolean Literals
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	// Evaluate Prefix Expressions
	case *ast.PrefixExpression:
		{
			right := Eval(node.Right, env)
			if isError(right) {
				return right
			}

			return evalPrefixExpression(node.Operator, right)
		}

	// Evaluate Infix Expressions
	case *ast.InfixExpression:
		{
			left := Eval(node.Left, env)
			if isError(left) {
				return left
			}

			right := Eval(node.Right, env)
			if isError(right) {
				return right
			}
			return evalInfixExpression(node.Operator, left, right)
		}

	// Evaluate Block Statements
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	// Evaluate If Expressions
	case *ast.IfExpression:
		return evalIfExpression(node, env)

	// Evaluate Function Literals
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)

	// Evaluate Call Expressions
	case *ast.CallExpression:
		{
			function := Eval(node.Function, env)
			if isError(function) {
				return function
			}

			return evalCallExpression(node, function, env)
		}

	// Evaluate Return Statements
	case *ast.ReturnStatement:
		{
			val := Eval(node.ReturnValue, env)
			if isError(val) {
				return val
			}
			return &object.ReturnValue{Value: val}
		}
	}
	return NULL
}

func evalProgram(node *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
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
