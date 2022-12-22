package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"testing"
)

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

	if !testIdentifier(testing, expressionStatement.Expression, "foobar") {
		return
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
		t.Fatalf("expression not *ast.Identifier. got=%T", expression)
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

// Function to test assignment expressions
func TestParsingAssignmentExpressions(t *testing.T) {
	assignmentTests := []struct {
		input    string
		leftVal  interface{}
		operator string
		rightVal interface{}
	}{
		{"x = 5;", "x", "=", 5},
		{"y = 10;", "y", "=", 10},
		{"foobar = y;", "foobar", "=", "y"},
	}

	// Test each assignment expression
	for _, assignmentTest := range assignmentTests {
		lex := lexer.New(assignmentTest.input)
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

		testAssignmentExpression(t, statement.Expression, assignmentTest.leftVal, assignmentTest.operator, assignmentTest.rightVal)
	}
}

// Function to test assignment expressions
//
// Tests:
//
// 1. Check if the expression is of type *ast.AssignmentExpression
//
// 2. Check if the left value is correct
//
// 3. Check if the operator is correct
//
// 4. Check if the right value is correct
func testAssignmentExpression(testing *testing.T, expression ast.Expression, left interface{}, operator string, right interface{}) bool {
	assignmentExpression, ok := expression.(*ast.AssignmentExpression)

	// Check if the expression
	if !ok {
		testing.Errorf("expression is not ast.AssignmentExpression. got=%T(%s)", expression, expression)
		return false
	}

	// Check if the left value is correct

	if !testLiteralExpression(testing, assignmentExpression.Name, left) {
		return false
	}

	// Check if the operator is correct
	if assignmentExpression.TokenLiteral() != operator {
		testing.Errorf("expression.Operator is not '%s'. got=%s", operator, assignmentExpression.TokenLiteral())
		return false
	}

	// Check if the right value is correct
	if !testLiteralExpression(testing, assignmentExpression.Value, right) {
		return false
	}

	return true
}
