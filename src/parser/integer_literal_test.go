package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"fmt"
	"testing"
)

// Function to test the parsing of integer literals
func TestIntegerLiteralExpression(testing *testing.T) {
	// Input string
	input := "5;"

	// Number of statements in the input string
	inputStatementNumber := 1

	// lexer instance
	lex := lexer.New(input)
	// creating parser instance that takes lexer instance as input
	parser := New(lex)

	// parse the program
	program := parser.ParseProgram()
	// check for parsing errors
	checkParserErrors(testing, parser)

	// Check if the length of the list of statements is equal to 1 or not
	if len(program.Statements) != inputStatementNumber {
		testing.Fatalf("program.Statements does not contain %d statements. got=%d", inputStatementNumber, len(program.Statements))
	}

	// get the first statement
	statement := program.Statements[0]

	// Get expression statement from the statement
	expressionStatement, ok := statement.(*ast.ExpressionStatement)
	// Check if the statement is of type *ast.ExpressionStatement
	if !ok {
		testing.Fatalf("statement not *ast.ExpressionStatement. got=%T", statement)
	}

	testIntegerLiteral(testing, expressionStatement.Expression, 5)
}

// Function to test integer literals
//
// Tests:
//
// 1. Check if the expression is of type *ast.IntegerLiteral
//
// 2. Check if the value is correct
//
// 3. Check if the token literal is correct
func testIntegerLiteral(t *testing.T, expression ast.Expression, value int64) bool {
	integer, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("integerLiteral not *ast.IntegerLiteral. got=%T", expression)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
		return false
	}

	return true
}
