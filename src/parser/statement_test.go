package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"testing"
)

func TestVarStatement(testing *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"var x = 5;", "x", 5},
		{"var y = true;", "y", true},
		{"var foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(testing, parser)

		if len(program.Statements) != 1 {
			testing.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		testVarStatement(testing, statement, tt.expectedIdentifier)

		value := statement.(*ast.VarStatement).Value

		if !testLiteralExpression(testing, value, tt.expectedValue) {
			return
		}
	}
}

// Test the variable statement
//
// Tests:
//
// 1. Check if the statement is of type *ast.VarStatement
//
// 2. Check if the token literal is 'var'
//
// 3. Check if the name of the variable is correct
func testVarStatement(testing *testing.T, statement ast.Statement, name string) bool {
	// Check if the statement is a variable statement
	if statement.TokenLiteral() != "var" {
		testing.Errorf("statement.TokenLiteral not 'var'. got=%q", statement.TokenLiteral())
		return false
	}

	// get the variable statement
	varStatement, ok := statement.(*ast.VarStatement)

	// Check if the variable statement is of type *ast.VarStatement
	if !ok {
		testing.Errorf("statement not *ast.VarStatement. got=%T", statement)
		return false
	}

	// Check if the name of the variable is correct
	if varStatement.Name.Value != name {
		testing.Errorf("varStatement.Name.Value not %s. got=%s", name, varStatement.Name.Value)
		return false
	}

	// Check if the token literal of the variable is correct
	if varStatement.Name.TokenLiteral() != name {
		testing.Errorf("varStatement.Name.TokenLiteral() not %s. got=%s", name, varStatement.Name.TokenLiteral())
		return false
	}

	return true
}

// Test the return statement
//
// Tests:
//
// 1. Check if the statement is of type *ast.ReturnStatement
//
// 2. Check if the token literal is 'return'
func TestReturnStatement(testing *testing.T) {
	tests := []struct {
		input       string
		expectedArg interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(testing, parser)

		if len(program.Statements) != 1 {
			testing.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		testReturnStatement(testing, statement, tt.expectedArg)
	}
}

// Test the return statement
//
// Tests:
//
// 1. Check if the statement is of type *ast.ReturnStatement
//
// 2. Check if the token literal is 'return'
func testReturnStatement(testing *testing.T, statement ast.Statement, expectedArg interface{}) bool {
	// Check if the statement is a return statement
	if statement.TokenLiteral() != "return" {
		testing.Errorf("statement.TokenLiteral not 'return'. got=%q", statement.TokenLiteral())
		return false
	}

	// get the return statement
	returnStatement, ok := statement.(*ast.ReturnStatement)

	// Check if the return statement is of type *ast.ReturnStatement
	if !ok {
		testing.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
		return false
	}

	// Check if the token literal of the return statement is correct
	if returnStatement.TokenLiteral() != "return" {
		testing.Errorf("returnStatement.TokenLiteral() not 'return'. got=%q", returnStatement.TokenLiteral())
		return false
	}

	// Check if the return value is correct
	if !testLiteralExpression(testing, returnStatement.ReturnValue, expectedArg) {
		return false
	}

	return true
}
