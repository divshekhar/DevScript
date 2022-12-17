package parser

import (
	"devscript/ast"
	"devscript/lexer"
	"devscript/token"
	"fmt"
)

/*
Parser struct represents the parser
*/
type Parser struct {
	// Lexer instance to get token stream
	lexer *lexer.Lexer

	// Pointer to the current token
	curToken token.Token
	// Pointer to the next token
	peekToken token.Token

	// List of errors
	errors []string

	/*
		Map of prefixParseFn functions
		Each function is associated with a token type
		Eg. prefixParseFns = {ADD: parsePrefixFunction, SUB: parsePrefixFunction, ...}
	*/
	prefixParseFns map[token.TokenType]prefixParseFn

	/*
		Map of infixParseFn functions
		Each function is associated with a token type
		Eg. infixParseFns = {ADD: parseInfixFunction, SUB: parseInfixFunction, ...}
	*/
	infixParseFns map[token.TokenType]infixParseFn
}

type (
	/*
		Function to parse a prefix expression.
		Eg: -5, !true, -true, !false, -false
		Returns an ast.Expression
	*/
	prefixParseFn func() ast.Expression

	/*
		Function to parse a infix expression
		Eg: 5 + 5, 5 / 5, 5 > 5, 5 == 5, 5 != 5
		Returns an ast.Expression
	*/
	infixParseFn func(ast.Expression) ast.Expression
)

/*
Register a prefixParseFn function for a token type.
Updates the prefixParseFns map.
*/
func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

/*
Register a infixParseFn function for a token type.
Updates the infixParseFns map.
*/
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

/*
Creates a new Parser instance, takes a Lexer instance as input
*/
func New(lex *lexer.Lexer) *Parser {
	// Create a new Parser struct instance
	parser := &Parser{
		lexer:  lex,
		errors: []string{},
	}

	/*
		After the first call to [nextToken()],
		curToken: 	peekToken = nil,
		peekToken:  NextToken() = <First Token>
	*/
	parser.nextToken()

	/*
		After the second call to [nextToken()],
		curToken: 	peekToken = <First Token>,
		peekToken:  NextToken() = <Second Token>
	*/
	parser.nextToken()

	return parser
}

/*
Returns the list of errors
*/
func (parser *Parser) Errors() []string {
	return parser.errors
}

/*
Function adds a peekError to the list of errors, if the next token is not of the expected type
*/
func (parser *Parser) peekError(nextToken token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", nextToken, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

/*
Function to get the next token from the lexer.
Sets the currentToken to the peekToken and updates the peekToken to the next token from the lexer.
*/
func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

/*
Function to parse the program
*/
func (parser *Parser) ParseProgram() *ast.Program {
	// Create a new Program struct instance
	program := &ast.Program{}
	// Initialize the list of statements
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		// Parse each statement
		statement := parser.parseStatement()

		if statement != nil {
			// Add the statement to the list of statements
			program.Statements = append(program.Statements, statement)
		}

		// Get the next token
		parser.nextToken()
	}
	return program
}

/*
Program contains a list of statements.
[ParseProgram()] calls [parseStatement()] to parse each statement.
*/
func (parser *Parser) parseStatement() ast.Statement {

	/*
		Every statement starts with a keyword. The keyword is the token type.
		This requires parsing of each statement to be different.
	*/
	switch parser.curToken.Type {
	// Parse variable statements
	case token.VAR:
		return parser.parseVarStatement()

	// Parse return statements
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return nil
	}
}

/*
Function parses the variable statements.
*/
func (parser *Parser) parseVarStatement() *ast.VarStatement {

	// Create a new VarStatement struct instance, set the token to the current token
	statement := &ast.VarStatement{Token: parser.curToken}

	// Check if the next token is an identifier
	// return nil if the next token is not an identifier
	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	/*
		Create a new Identifier struct instance,
		set the token to the current token and the value to the literal value of the current token
		Update the Name field of the VarStatement struct instance
	*/
	statement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// Check if the next token is an assignment operator
	// return nil if the next token is not an assignment operator
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until semicolon is encountered
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

/*
Function parses the return statements.
*/
func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	/*
		Create a new ReturnStatement struct instance,
		set the token to the current token
	*/
	statement := &ast.ReturnStatement{Token: parser.curToken}

	// Get next token
	parser.nextToken()

	// TODO: skipping the expressions until semicolon is encountered
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

/*
Function to check if the current token is of the expected type
*/
func (parser *Parser) curTokenIs(token token.TokenType) bool {
	return parser.curToken.Type == token
}

/*
Function to check if the peek token is of the expected type
*/
func (parser *Parser) peekTokenIs(token token.TokenType) bool {
	return parser.peekToken.Type == token
}

/*
Function to update the curToken & peekToken pointer,
if the next token is of the expected type.

Add error to the list of errors,
when the next token is not of the expected type.
*/
func (parser *Parser) expectPeek(token token.TokenType) bool {
	if parser.peekTokenIs(token) {
		parser.nextToken()
		return true
	} else {
		// Add error to the list of errors
		parser.peekError(token)
		return false
	}
}
