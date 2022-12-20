package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"fmt"
	"testing"
)

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
