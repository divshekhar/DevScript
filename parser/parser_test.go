package parser

import (
	"devscript/ast"
	"devscript/lexer"
	"fmt"
	"testing"
)

/*
Check if there are any errors in the parser,
if there are, print the errors and fail the test
*/
func checkParserErrors(testing *testing.T, parser *Parser) {
	// Get the errors list from the parser
	errors := parser.Errors()

	// If there are no errors, return
	if len(errors) == 0 {
		return
	}

	// print the number of errors
	testing.Errorf("parser has %d errors", len(errors))

	// print the errors
	for _, msg := range errors {
		testing.Errorf("parser error: %q", msg)
	}

	// Fail the test
	testing.FailNow()
}

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

func TestIdentifierExpression(testing *testing.T) {
	// Input string
	input := "foobar;"

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

	// Check if the statement is of type *ast.ExpressionStatement
	expressionStatement, ok := statement.(*ast.ExpressionStatement)
	if !ok {
		testing.Fatalf("statement not *ast.ExpressionStatement. got=%T", statement)
	}

	testIdentifier(testing, expressionStatement.Expression, "foobar")
}

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

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	// Test each prefix expression
	for _, prefixText := range prefixTests {
		// Create lexer instance
		lex := lexer.New(prefixText.input)
		// Create parser instance
		parser := New(lex)
		// Parse the program
		program := parser.ParseProgram()
		// Check for parsing errors
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("statement.Expression is not ast.PrefixExpression. got=%T", statement.Expression)
		}

		if expression.Operator != prefixText.operator {
			t.Fatalf("expression.Operator is not '%s'. got=%s", prefixText.operator, expression.Operator)
		}

		if !testIntegerLiteral(t, expression.Right, prefixText.integerValue) {
			return
		}
	}
}

// Function to test the parsing of infix expressions
func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		leftVal  interface{}
		operator string
		rightVal interface{}
	}{
		{"10 + 5;", 10, "+", 5},
		{"5 - 10;", 5, "-", 10},
		{"10 * 5;", 10, "*", 5},
		{"5 / 10;", 5, "/", 10},
		{"5 > 5;", 5, ">", 5},
		{"10 < 5;", 10, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},

		// After adding support for boolean literals
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	// Test each infix expression
	for _, infixTest := range infixTests {
		lex := lexer.New(infixTest.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		testInfixExpression(t, statement.Expression, infixTest.leftVal, infixTest.operator, infixTest.rightVal)
	}
}

// Function to test infix expressions
//
// Tests:
//
// 1. Check if the expression is of type *ast.InfixExpression
//
// 2. Check if the left value is correct
//
// 3. Check if the operator is correct
//
// 4. Check if the right value is correct
func testInfixExpression(testing *testing.T, expression ast.Expression, left interface{}, operator string, right interface{}) bool {
	infixExpression, ok := expression.(*ast.InfixExpression)

	// Check if the expression is of type *ast.InfixExpression
	if !ok {
		testing.Errorf("expression is not ast.InfixExpression. got=%T(%s)", expression, expression)
		return false
	}

	// Check if the left value is correct
	if !testLiteralExpression(testing, infixExpression.Left, left) {
		return false
	}

	// Check if the operator is correct
	if infixExpression.Operator != operator {
		testing.Errorf("expression.Operator is not '%s'. got=%s", operator, infixExpression.Operator)
		return false
	}

	// Check if the right value is correct
	if !testLiteralExpression(testing, infixExpression.Right, right) {
		return false
	}

	return true
}

// Function to test literal expressions
//
// Tests IntegerLiteral, Identifier & BooleanLiteral
func testLiteralExpression(testing *testing.T, expression ast.Expression, expected interface{}) bool {
	switch value := expected.(type) {
	case int:
		return testIntegerLiteral(testing, expression, int64(value))
	case int64:
		return testIntegerLiteral(testing, expression, value)
	case string:
		return testIdentifier(testing, expression, value)
	case bool:
		return testBooleanLiteral(testing, expression, value)
	}
	testing.Errorf("[testLiteralExpression] Type of expression not handled, got=%T", expression)
	return false
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

// Function to test identifiers
//
// Tests:
//
// 1. Check if the expression is of type *ast.Identifier
//
// 2. Check if the value is correct
//
// 3. Check if the token literal is correct
func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	identifier, ok := expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression not *ast.Identifier. got=%T", expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got=%s", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", value, identifier.TokenLiteral())
		return false
	}

	return true

}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		testBooleanLiteral(t, stmt.Expression, tt.expected)
	}
}

// Function to test boolean expressions
//
// Tests:
//
// 1. Check if the statement is of type *ast.ExpressionStatement
//
// 2. Check if the expression is of type *ast.Boolean
//
// 3. Check if the value is correct
//
// 4. Check if the token literal is correct
func testBooleanLiteral(t *testing.T, expression ast.Expression, expected bool) bool {
	boolean, ok := expression.(*ast.Boolean)
	if !ok {
		t.Errorf("expression not *ast.Boolean. got=%T", expression)
		return false
	}

	if boolean.Value != expected {
		t.Errorf("boolean.Value not %t. got=%t", expected, boolean.Value)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", expected) {
		t.Errorf("boolean.TokenLiteral not %t. got=%s", expected, boolean.TokenLiteral())
		return false
	}

	return true
}

// Function to test operator precedence parsing
func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		// After adding boolean literals
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
