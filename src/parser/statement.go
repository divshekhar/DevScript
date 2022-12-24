package parser

import (
	"devscript/src/ast"
	"devscript/src/token"
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

	// Check data type of the variable
	// return nil if the next token is not a colon
	if !parser.expectPeek(token.COLON) {
		return nil
	}

	// Check if the next token is a data type
	if !parser.expectPeek(token.DATATYPE) {
		return nil
	}

	// Set the data type of the variable
	switch parser.curToken.Literal {
	case "int":
		statement.DataType = token.INT
	case "string":
		statement.DataType = token.STRING
	case "bool":
		statement.DataType = token.BOOL
	case "function":
		statement.DataType = token.FUNCTION
	default:
		statement.DataType = token.NULL
	}

	// Check if the next token is an assignment operator
	// return nil if the next token is not an assignment operator
	//
	// curToken: {Type: token.IDENT, Literal: "x"}
	// peekToken: {Type: token.ASSIGN, Literal: "="}
	if !parser.peekTokenIs(token.ASSIGN) {

		// if the next token is a semicolon,
		// then the variable statement is a declaration statement
		//
		// curToken: {Type: token.IDENT, Literal: "x"}
		// peekToken: {Type: token.SEMICOLON, Literal: ";"}
		if parser.peekTokenIs(token.SEMICOLON) {
			parser.nextToken()
		}

		// Default value of the variable is 0 (integer)
		defaultValue := parser.defaultValue(statement.DataType)
		statement.Value = defaultValue

		return statement
	}

	// Advance the current token to the next token
	//
	// curToken: {Type: token.ASSIGN, Literal: "="}
	// peekToken: {Type: token.INT, Literal: "5"}
	parser.nextToken()

	// curToken: {Type: token.INT, Literal: "5"}
	// peekToken: {Type: token.SEMICOLON, Literal: ";"}
	parser.nextToken()

	// Parse the expression
	value := parser.parseExpression(LOWEST)

	// Check if the expression is of the same data type as the variable
	if value.TokenType() != string(statement.DataType) {
		parser.errors = append(parser.errors, "type mismatch")
		return nil
	}

	statement.Value = value

	// Check if the next token is a semicolon
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) defaultValue(tokenType token.TokenType) ast.Expression {
	switch tokenType {
	case token.INT:
		return &ast.IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "0"},
			Value: 0,
		}
	case token.STRING:
		return &ast.StringLiteral{
			Token: token.Token{Type: token.STRING, Literal: ""},
			Value: "",
		}
	case token.BOOL:
		return &ast.Boolean{
			Token: token.Token{Type: token.FALSE, Literal: "false"},
			Value: false,
		}
	case token.FUNCTION:
		return &ast.FunctionLiteral{
			Token:      token.Token{Type: token.FUNCTION, Literal: "function"},
			Parameters: []*ast.Identifier{},
			Body: &ast.BlockStatement{
				Token:      token.Token{Type: token.LBRACE, Literal: "{"},
				Statements: []ast.Statement{},
			},
		}
	default:
		return nil
	}
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
	if parser.peekTokenIs(token.SEMICOLON) {
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
