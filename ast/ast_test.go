package ast

import (
	"devscript/token"
	"testing"
)

func TestString(t *testing.T) {
	input := "var x = y;"

	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
		},
	}

	if program.String() != input {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
