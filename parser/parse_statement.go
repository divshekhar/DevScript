package parser

import (
	"devscript/ast"
	"devscript/token"
)

// Program contains a list of statements.
//
//	ParseProgram() calls parseStatement() to parse each statement.
//	parseStatement() calls:
//		parseVarStatement() 	// variable statements
//		parseReturnStatement() 	// return statements
//		parseExpressionStatement() 	// expression statements
func (parser *Parser) parseStatement() ast.Statement {

	// Every statement starts with a keyword. The keyword is the token type.
	// This requires parsing of each statement to be different.
	switch parser.curToken.Type {
	// Parse variable statements
	case token.VAR:
		return parser.parseVarStatement()

	// Parse return statements
	case token.RETURN:
		return parser.parseReturnStatement()

	// Parse expression statements (default)
	default:
		return parser.parseExpressionStatement()
	}
}

// Function parses the variable statements.
//
// Variable statements are statements that start with the keyword "var".
//
//	var x = 5;
func (parser *Parser) parseVarStatement() *ast.VarStatement {

	// Create a new VarStatement struct instance, set the token to the current token
	statement := &ast.VarStatement{Token: parser.curToken}

	// Check if the next token is an identifier
	// return nil if the next token is not an identifier
	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	// Create a new Identifier struct instance,
	// set the token to the current token and the value to the literal value of the current token
	// Update the Name field of the VarStatement struct instance
	statement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// Check if the next token is an assignment operator
	// return nil if the next token is not an assignment operator
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// Advance the current token to the next token
	parser.nextToken()

	// Parse the expression
	statement.Value = parser.parseExpression(LOWEST)

	// Check if the next token is a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

// Function parses the return statements.
//
// Return statements are statements that start with the keyword "return".
//
//	return 5;
//	return x;
func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	/*
		Create a new ReturnStatement struct instance,
		set the token to the current token
	*/
	statement := &ast.ReturnStatement{Token: parser.curToken}

	// Get next token
	parser.nextToken()

	// Parse the expression
	statement.ReturnValue = parser.parseExpression(LOWEST)

	// Check if the next token is a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

// Function to parse the expression statements.
//
// Expression statements are statements that do not start with a keyword.
//
//	5;  		// Integer literal
//	foobar; 	// Identifier
//	-5; 		// Prefix expression
//	x + 5; 		// Infix expression
//	true; 		// Boolean expression
func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// Create a new ExpressionStatement struct instance, set the token to the current token
	statement := &ast.ExpressionStatement{Token: parser.curToken}

	// Parse the expression
	statement.Expression = parser.parseExpression(LOWEST)

	// Check if the next token is a semicolon
	// return nil if the next token is not a semicolon
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

// Function to parse the block statements.
//
// Block statements are statements that start with a LBRACE token and end with a RBRACE token.
//
//	{
//		5;
//		var x = 5;
//	}
func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	// Create a new BlockStatement struct instance, set the token to the current token
	block := &ast.BlockStatement{Token: parser.curToken}

	// Set the statements field to an empty slice
	block.Statements = []ast.Statement{}

	// Advance the current token to the next token
	parser.nextToken()

	// Parse the statements until the next token is a RBRACE token
	for !parser.curTokenIs(token.RBRACE) {
		statement := parser.parseStatement()
		if statement != nil {
			block.Statements = append(block.Statements, statement)
		}
		parser.nextToken()
	}

	return block
}
