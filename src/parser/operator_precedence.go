package parser

import "devscript/src/token"

// Precedence of operators
const (
	_ int = iota
	LOWEST
	// =
	ASSIGN
	// == or !=
	EQUALS
	// < or >
	LESSGREATER
	// + or -
	SUM
	// * or /
	PRODUCT
	// -X or !X
	PREFIX
	// myFunction(X)
	CALL
)

// Map of precedences
var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGN,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

// Peek the precedence of the next token
//
// If the next token is not in the precedences map, return
//
//	LOWEST
func (parser *Parser) peekPrecedence() int {
	if p, ok := precedences[parser.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

// Get the precedence of the current token
//
// If the current token is not in the precedences map, return
//
//	LOWEST
func (parser *Parser) curPrecedence() int {
	if p, ok := precedences[parser.curToken.Type]; ok {
		return p
	}

	return LOWEST
}
