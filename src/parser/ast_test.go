/*
	AST Test is done in parser package to avoid cyclic imports between ast and parser packages
*/

package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"devscript/src/token"
	"testing"
)

func TestAST(t *testing.T) {
	input := `
	var x = y;
	return true;
	return 100 < 200;
	`

	expected := &ast.Program{
		Statements: []ast.Statement{
			// var x = y;
			&ast.VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
			// return true;
			&ast.ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
			// return 100 < 200;
			&ast.ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &ast.InfixExpression{
					Token: token.Token{Type: token.LT, Literal: "<"},
					Left: &ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "100"},
						Value: 100,
					},
					Operator: "<",
					Right: &ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "200"},
						Value: 200,
					},
				},
			},
		},
	}

	lex := lexer.New(input)
	parser := New(lex)
	program := parser.ParseProgram()

	if expected.String() != program.String() {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
