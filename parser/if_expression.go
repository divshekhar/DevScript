package parser

import (
	"devscript/ast"
	"devscript/token"
)

// Function to parse if expressions
//
//	if (x < y) { x } else { y };		// parseIfExpression
//	var z = if (x < y) { x } else { y };	// parseIfExpression
func (parser *Parser) parseIfExpression() ast.Expression {
	// Create a new if expression
	exp := &ast.IfExpression{Token: parser.curToken}

	// Check if the next token is a LPAREN token
	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	// Advance the current token to the next token
	parser.nextToken()

	// Parse the condition
	exp.Condition = parser.parseExpression(LOWEST)

	// Check if the next token is a RPAREN token
	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	// Check if the next token is a LBRACE token
	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	// Parse the consequence
	exp.Consequence = parser.parseBlockStatement()

	// Check if the next token is an ELSE token
	if parser.peekTokenIs(token.ELSE) {
		// Advance the current token to the next token
		parser.nextToken()

		// Check if the next token is a LBRACE token
		if !parser.expectPeek(token.LBRACE) {
			return nil
		}

		// Parse the alternative
		exp.Alternative = parser.parseBlockStatement()
	}

	return exp
}
