package parser

import (
	"devscript/ast"
	"devscript/token"
)

func (parser *Parser) parseFunctionExpression() ast.Expression {
	// Create a new function expression
	functionExpression := &ast.FunctionExpression{Token: parser.curToken}

	// Check if the next token is a LPAREN token
	// For function expressions, peek token should be {token.IDENT, "foo"}
	if !parser.peekTokenIs(token.IDENT) {
		// if the next token is a LPAREN token, then it is a function literal
		if parser.peekTokenIs(token.LPAREN) {
			return parser.parseFunctionLiteral()
		}

		return nil
	}

	// Advance the token
	// from {token.FUNCTION, "func"} to {token.IDENT, "foo"}
	parser.nextToken()

	// Parse the function name
	functionExpression.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// Check if the next token is a LPAREN token
	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse the function parameters
	functionExpression.Parameters = parser.parseFunctionParameters()

	// Check if the next token is a LBRACE token
	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	// Parse the function body
	functionExpression.Body = parser.parseBlockStatement()

	return functionExpression
}

// Function to parse the function literals
//
//	fn(x, y) { x + y };	// parseFunctionLiteral
func (parser *Parser) parseFunctionLiteral() ast.Expression {
	// Create a new function literal
	functionLiteral := &ast.FunctionLiteral{Token: parser.curToken}

	// Check if the next token is a LPAREN token
	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse the function parameters
	functionLiteral.Parameters = parser.parseFunctionParameters()

	// Check if the next token is a LBRACE token
	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	// Parse the function body
	functionLiteral.Body = parser.parseBlockStatement()

	return functionLiteral
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
