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

/*
Check if the statement is a variable statement and if the name of the variable is correct
*/
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

	// Check if the name of the variable is correct
	if varStatement.Name.TokenLiteral() != name {
		testing.Errorf("varStatement.Name.TokenLiteral() not %s. got=%s", name, varStatement.Name.TokenLiteral())
		return false
	}

	return true
}

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

	// Check if the expression is of type *ast.Identifier
	identifier, ok := expressionStatement.Expression.(*ast.Identifier)
	if !ok {
		testing.Fatalf("expression not *ast.Identifier. got=%T", expressionStatement.Expression)
	}

	// Check if the name of the identifier is correct
	if identifier.Value != "foobar" {
		testing.Errorf("identifier.Value not %s. got=%s", "foobar", identifier.Value)
	}

	// Check if the token literal of the identifier is correct
	if identifier.TokenLiteral() != "foobar" {
		testing.Errorf("identifier.TokenLiteral not %s. got=%s", "foobar", identifier.TokenLiteral())
	}
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

	// Check if the expression is of type *ast.IntegerLiteral
	integerLiteral, ok := expressionStatement.Expression.(*ast.IntegerLiteral)
	if !ok {
		testing.Fatalf("expression not *ast.IntegerLiteral. got=%T", expressionStatement.Expression)
	}

	// Check if the value of the integer literal is correct
	if integerLiteral.Value != 5 {
		testing.Errorf("integerLiteral.Value not %d. got=%d", 5, integerLiteral.Value)
	}

	// Check if the token literal of the integer literal is correct
	if integerLiteral.TokenLiteral() != "5" {
		testing.Errorf("integerLiteral.TokenLiteral not %s. got=%s", "5", integerLiteral.TokenLiteral())
	}
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

func testIntegerLiteral(t *testing.T, integerLiteral ast.Expression, value int64) bool {
	integer, ok := integerLiteral.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("integerLiteral not *ast.IntegerLiteral. got=%T", integerLiteral)
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

// Function to test the parsing of infix expressions
func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		leftVal  int
		operator string
		rightVal int
	}{
		{"10 + 5;", 10, "+", 5},
		{"5 - 10;", 5, "-", 10},
		{"10 * 5;", 10, "*", 5},
		{"5 / 10;", 5, "/", 10},
		{"5 > 5;", 5, ">", 5},
		{"10 < 5;", 10, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
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

		infixExp, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("statement.Expression is not ast.InfixExpression. got=%T", statement.Expression)
		}

		// check if the left value is correct
		if !testIntegerLiteral(t, infixExp.Left, int64(infixTest.leftVal)) {
			return
		}

		// check if the operator is correct
		if infixExp.Operator != infixTest.operator {
			t.Fatalf("infixExp.Operator is not '%s'. got=%s", infixTest.operator, infixExp.Operator)
		}

		// check if the right value is correct
		if !testIntegerLiteral(t, infixExp.Right, int64(infixTest.rightVal)) {
			return
		}
	}
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
