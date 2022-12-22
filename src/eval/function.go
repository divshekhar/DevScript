package eval

import (
	"devscript/src/ast"
	"devscript/src/object"
)

// evaluates a function expression
//
//	func add(x, y) {
//	  return x + y;
//	}
func evalFunctionExpression(node *ast.FunctionExpression, env *object.Environment) object.Object {
	name := node.Name
	parameters := node.Parameters
	body := node.Body
	obj := &object.Function{Name: name, Parameters: parameters, Body: body}

	// mapping function object to function name
	env.Set(name.Value, obj)

	return obj
}

// evaluates a function literal.
//
//	fn(x, y) { x + y; }; // function literal
func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	params := node.Parameters
	body := node.Body
	return &object.Function{Parameters: params, Env: env, Body: body}
}

// evaluates a call expression.
//
//	add(1, 2); // call expression
func evalCallExpression(node *ast.CallExpression, function object.Object, env *object.Environment) object.Object {
	args := evalExpressions(node.Arguments, env)

	// If the first argument is an error, return it.
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args)
}

// Evaluates a list of expressions.
//
//	1, 2, 3; // list of expressions
//
// Function arguments are also expressions, and list of expressions are evaluated
// from left to right.
func evalExpressions(expressions []ast.Expression, env *object.Environment) []object.Object {
	// A slice to hold the evaluated expressions
	var result []object.Object

	for _, expression := range expressions {
		// Evaluate the expression
		// The result is an object.Object (Integer, Boolean, etc.)
		evaluated := Eval(expression, env)

		// If the evaluated expression is an error, return it.
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

// Applies a function to a list of arguments.
// Takes function and argument list as arguments.
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		// Evaluate the function body in the new environment
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	// If the function is a builtin function, call it.
	case *object.Builtin:
		return fn.Function(args...)
		
	default:
		return newError("not a function: %s", fn.Type())
	}
}

// Creates a new environment for the function.
// Sets the function parameters to the arguments.
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	// Create a new environment with the function's environment as the outer environment
	env := object.NewEnclosedEnvironment(fn.Env)

	// Set the function parameters to the arguments
	// The arguments are evaluated from left to right
	// Example:
	//		var x = 10;
	//		func(x, y) { x + y; };
	//		add(1, 2);
	//
	// Parent Environment:
	//		{
	//			x: 10
	//		}
	//
	// New environment:
	//		{
	//			x: 1,
	//			y: 2
	//		}
	for parameterIndex, parameter := range fn.Parameters {
		env.Set(parameter.Value, args[parameterIndex])
	}

	return env
}

// If the object is a ReturnValue, return the value from the object.
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}
