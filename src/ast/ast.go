package ast

import (
	"bytes"
)

// AST is represented in form of nodes
type Node interface {
	TokenLiteral() string
	TokenType() string
	String() string
}

// Program contains a list of statements
type Program struct {
	Statements []Statement
}

// Returns the token literal of the first statement,
// empty string if there are no statements
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) TokenType() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenType()
	} else {
		return ""
	}
}

// String reprentation of the program
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
