package ast

import (
	"bytes"
	"devscript/src/token"
	"strings"
)

// Expression is a Node that can be evaluated to a value
type Expression interface {
	Node
	expressionNode()
}

// Identifier is a node that represents
// variable/function name
//
//	foobar;
//	a;
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

// IntegerLiteral is a node that represents an integer value
//
//	5;
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

// StringLiteral is a node that represents a string value
//
//	"foobar";
type StringLiteral struct {
	// token.STRING token
	Token token.Token
	// value of the string
	Value string
}

func (stringLiteral *StringLiteral) expressionNode() {}
func (stringLiteral *StringLiteral) TokenLiteral() string {
	return stringLiteral.Token.Literal
}
func (stringLiteral *StringLiteral) String() string {
	return stringLiteral.Token.Literal
}

// Prefix Expression is a node that represents a prefix expression
//
//	-5, !true
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
//	5 + 5, 5 > 5
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

// Boolean is a node that represents a boolean value
//
//	true, false
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression is a node that represents an if expression
//
//	if (x < y) { x } else { y }
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// FunctionLiteral is a node that represents a function literal
//
//	func(x, y) { x + y; }
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// FunctionExpression
type FunctionExpression struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (functionExpression *FunctionExpression) expressionNode() {}
func (functionExpression *FunctionExpression) TokenLiteral() string {
	return functionExpression.Token.Literal
}
func (functionExpression *FunctionExpression) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range functionExpression.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(functionExpression.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(functionExpression.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(functionExpression.Body.String())

	return out.String()
}

// CallExpression is a node that represents a function call
//
//	add(1, 2 * 3, 4 + 5);
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// AssignmentExpression is a node that represents an assignment expression
//
//	x = 5;
type AssignmentExpression struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ae *AssignmentExpression) expressionNode() {}
func (ae *AssignmentExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(" = ")
	out.WriteString(ae.Value.String())

	return out.String()
}
