package ast

import (
	"bytes"
	"devscript/token"
)

// Statement is a Node that can be executed
//
// Example:
//
//	var x = 5;	// var statement
//	return 5;	// return statement
//	5 + 5;		// expression statement
type Statement interface {
	Node
	statementNode()
}

// Var statement is a statement that declares a variable.
//
//	var x = 5; // var statement
type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}
func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Return statement is a statement that
// returns a expression value from a function
//
//	return 5; // return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Expression statement contains an expression
//
//	5 + 5; // expression statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
