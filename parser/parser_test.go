package parser

import (
	"devscript/ast"
	"devscript/lexer"
	"testing"
)

func TestVarStatement(t *testing.T) {
	input := `
	var x = 5;
	var y = 10;
	var z = 10000;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testVarStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}

}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 100;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func testVarStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "var" {
		t.Errorf("statement.TokenLiteral not 'var'. got=%q", statement.TokenLiteral())
		return false
	}

	varStatement, ok := statement.(*ast.VarStatement)
	if !ok {
		t.Errorf("statement not *ast.VarStatement. got=%T", statement)
		return false
	}

	if varStatement.Name.Value != name {
		t.Errorf("varStatement.Name.Value not %s. got=%s", name, varStatement.Name.Value)
		return false
	}

	if varStatement.Name.TokenLiteral() != name {
		t.Errorf("varStatement.Name.TokenLiteral() not %s. got=%s", name, varStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
