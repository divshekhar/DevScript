package ast

import (
	"bytes"
	"devscript/token"
)

/*
AST is represented in form of nodes
*/
type Node interface {
	TokenLiteral() string
	String() string
}

/*
Statement is a Node that can be executed
*/
type Statement interface {
	Node
	statementNode()
}

/*
Expression is a Node that can be evaluated to a value
*/
type Expression interface {
	Node
	expressionNode()
}

/*
Program contains a list of statements
*/
type Program struct {
	Statements []Statement
}

/*
Returns the token literal of the first statement,
empty string if there are no statements
*/
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/*
String reprentation of the program
*/
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*
Var statement is a statement that declares a variable.
Example: var x = 5;
*/
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

/*
Identifier is a node that represents
variable/function name
*/
type Identifier struct {
	// token.IDENT token
	Token token.Token
	// value of the identifier (variable/function name)
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

/*
Return statement is a statement that
returns a expression value from a function
*/
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

/*
Expression statement contains an expression
*/
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

// IntegerLiteral is a node that represents an integer value
type IntegerLiteral struct {
	// token.INT token
	Token token.Token
	// value of the integer
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// returns the token literal of the integer literal
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// string representation of the integer literal
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	// Prefix token
	//
	// Example: -5, !true
	//
	// The token is the - (MINUS) or ! (BANG) token
	Token    token.Token
	Operator string
	Right    Expression
}

func (prefixExp *PrefixExpression) expressionNode() {}
func (prefixExp *PrefixExpression) TokenLiteral() string {
	return prefixExp.Token.Literal
}
func (prefixExp *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExp.Operator)
	out.WriteString(prefixExp.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is a node that represents an infix expression
//
// Example: 5 + 5, 5 > 5
type InfixExpression struct {
	// Infix token
	//
	// Example: 5 + 5, 5 > 5
	//
	// The token is the + (PLUS) or > (GT) token
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (infixExp *InfixExpression) expressionNode() {}
func (infixExp *InfixExpression) TokenLiteral() string {
	return infixExp.Token.Literal
}
func (infixExp *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infixExp.Left.String())
	out.WriteString(" " + infixExp.Operator + " ")
	out.WriteString(infixExp.Right.String())
	out.WriteString(")")

	return out.String()
}
