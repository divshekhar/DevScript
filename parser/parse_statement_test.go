package parser

import (
	"devscript/ast"
	"devscript/lexer"
	"testing"
)

func TestVarStatement(testing *testing.T) {

	// Input string
	input := `
	var x = 5;
	var y = 10;
	var z = 10000;
	`
	// Number of statements in the input string
	var inputStatementNumber int = 3

	// lexer instance
	lex := lexer.New(input)

	// creating parser instance that takes lexer instance as input
	parser := New(lex)

	// parse the program
	program := parser.ParseProgram()

	// check for parsing errors
	checkParserErrors(testing, parser)

	/*
		Given input string contain 'n' statements,

		Check if the length of the list of statements is 'n' if not then
		terminate the test and print the error message
	*/
	if len(program.Statements) != inputStatementNumber {
		testing.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	// Expected test results
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}

	for i, testResult := range tests {
		statement := program.Statements[i]
		if !testVarStatement(testing, statement, testResult.expectedIdentifier) {
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
	// Input string
	input := `
	return 5;
	return 10;
	return 100;
	`

	var inputReturnStatements int = 3

	// lexer instance
	lex := lexer.New(input)
	// creating parser instance that takes lexer instance as input
	parser := New(lex)

	// parse the program
	program := parser.ParseProgram()
	// check for parsing errors
	checkParserErrors(testing, parser)

	// Check if the length of the list of statements is equal to 'inputReturnStatements' or not
	if len(program.Statements) != inputReturnStatements {
		testing.Fatalf("program.Statements does not contain %d statements. got=%d", inputReturnStatements, len(program.Statements))
	}

	for _, statement := range program.Statements {
		// get statemet
		returnStatement, ok := statement.(*ast.ReturnStatement)

		// Check if the statement is of type *ast.ReturnStatement
		if !ok {
			testing.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
		}

		// Check if the token literal is 'return'
		if returnStatement.TokenLiteral() != "return" {
			testing.Errorf("returnStmt.TokenLiteral() not 'return', got %q", returnStatement.TokenLiteral())
		}
	}
}
