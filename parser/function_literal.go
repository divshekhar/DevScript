package parser

import (
	"devscript/ast"
	"devscript/token"
)

// Function to parse the function literals
//
//	fn(x, y) { x + y };	// parseFunctionLiteral
func (parser *Parser) parseFunctionLiteral() ast.Expression {
	// Create a new function literal
	lit := &ast.FunctionLiteral{Token: parser.curToken}

	// Check if the next token is a LPAREN token
	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse the function parameters
	lit.Parameters = parser.parseFunctionParameters()

	// Check if the next token is a LBRACE token
	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	// Parse the function body
	lit.Body = parser.parseBlockStatement()

	return lit
}

// Function to parse the function parameters
//
//	(x, y, z)	// parseFunctionParameters
func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	// Create a new slice of identifiers
	identifiers := []*ast.Identifier{}

	// Check if the next token is a RPAREN token
	if parser.peekTokenIs(token.RPAREN) {
		// Advance the current token to the next token
		parser.nextToken()
		return identifiers
	}

	// Advance the current token to the next token
	parser.nextToken()

	// Create a new identifier
	identifier := &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// Append the identifier to the identifiers slice
	identifiers = append(identifiers, identifier)

	// Check if the next token is a COMMA token
	for parser.peekTokenIs(token.COMMA) {
		// Advance the current token to the next token
		parser.nextToken()

		// Advance the current token to the next token
		parser.nextToken()

		// Create a new identifier
		identifier := &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

		// Append the identifier to the identifiers slice
		identifiers = append(identifiers, identifier)
	}

	// Check if the next token is a RPAREN token
	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}
