package parser

import (
	"devscript/src/ast"
	"devscript/src/token"
)

// Function to parse grouped expressions
//
//	(5 + 5) * 2;	// parseGroupedExpression
func (parser *Parser) parseGroupedExpression() ast.Expression {
	// Advance the current token to the next token
	parser.nextToken()

	// Parse the expression with precedence of LOWEST
	exp := parser.parseExpression(LOWEST)

	// Check if the next token is a RPAREN token
	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
