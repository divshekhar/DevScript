package parser

import (
	"devscript/src/ast"
	"devscript/src/token"
	"fmt"
)

// Function to return error if there is no prefix parse function for the current token
//
//	5;	// parseIntegerLiteral
//	foo;	// parseIdentifier
//	-5;	// parsePrefixExpression
//	false;	// parseBoolean
func (parser *Parser) noPrefixParseFnError(tokenType token.TokenType) {
	fmt.Printf("No prefix parse function for %s found\n", tokenType)
}

// Function to parse the expressions
//
//	5; 		// parseIntegerLiteral
//	10 + 15; 	// parseInfixExpression
//	-5;		// parsePrefixExpression
//	foo;		// parseIdentifier
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	// Get the prefix function for the current token
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.curToken.Type)
		return nil
	}

	// exp is the left expression of the infix expression
	exp := prefix()

	// Check if the next token is not a semicolon and,
	// precedence of the current token is less than the next token's precedence
	for !parser.peekTokenIs(token.SEMICOLON) && precedence < parser.peekPrecedence() {
		// Get the infix function for the next token
		infix := parser.infixParseFns[parser.peekToken.Type]
		if infix == nil {
			return exp
		}

		parser.nextToken()

		// Parse the infix expression
		// calls parseInfixExpression() that takes the left expression as an argument
		exp = infix(exp)
	}

	return exp
}

// Function to parse the prefix expressions
//
//	-5;	// parsePrefixExpression
//	!true;	// parsePrefixExpression
func (parser *Parser) parsePrefixExpression() ast.Expression {
	// Create a new PrefixExpression struct instance, set the token to the current token
	prefixExpression := &ast.PrefixExpression{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
	}

	// Advance to the next token
	parser.nextToken()

	// Initialize the Right field of the PrefixExpression struct instance, with precedence of PREFIX
	prefixExpression.Right = parser.parseExpression(PREFIX)

	return prefixExpression
}

// Function to parse Infix expressions
//
//	5 + 5;		// parseInfixExpression
func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// Create a new InfixExpression struct instance, set the token to the current token
	infixExpression := &ast.InfixExpression{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
		Left:     left,
	}

	// Get the precedence of the current token
	precedence := parser.curPrecedence()

	// Advance to the next token
	parser.nextToken()

	// Initialize the Right field of the InfixExpression struct instance, with precedence of the current token
	infixExpression.Right = parser.parseExpression(precedence)

	return infixExpression
}

// Function to parse the identifiers.
//
//	foobar; x; y; z;
func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
}

// Function to parse assignment expressions
//
//	foo = 5;	// parseAssignmentExpression
func (parser *Parser) parseAssignmentExpression(left ast.Expression) ast.Expression {
	// Create a new AssignmentExpression struct instance, set the token to the current token
	assignmentExpression := &ast.AssignmentExpression{
		Token: parser.curToken,
		Name:  left.(*ast.Identifier),
	}

	// Advance to the next token
	parser.nextToken()

	// Initialize the Value field of the AssignmentExpression struct instance, with precedence of ASSIGNMENT
	assignmentExpression.Value = parser.parseExpression(ASSIGN)

	return assignmentExpression
}

// Function to parse the string literals
//
//	"foobar"; "foo bar";
func (parser *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: parser.curToken, Value: parser.curToken.Literal}
}
