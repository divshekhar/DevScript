package parser

import (
	"devscript/ast"
	"devscript/lexer"
	"devscript/token"
	"fmt"
	"strconv"
)

// Precedence of operators
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Map of precedences
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

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

	// Initialize the prefixParseFns map
	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	// Register the prefixParseFn for the token type
	parser.registerPrefix(token.IDENT, parser.parseIdentifier)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)
	parser.registerPrefix(token.BANG, parser.parsePrefixExpression)
	parser.registerPrefix(token.MINUS, parser.parsePrefixExpression)
	parser.registerPrefix(token.TRUE, parser.parseBoolean)
	parser.registerPrefix(token.FALSE, parser.parseBoolean)

	// Initialize the infixParseFns map
	parser.infixParseFns = make(map[token.TokenType]infixParseFn)
	// Register the infixParseFn for the token type
	parser.registerInfix(token.PLUS, parser.parseInfixExpression)
	parser.registerInfix(token.MINUS, parser.parseInfixExpression)
	parser.registerInfix(token.SLASH, parser.parseInfixExpression)
	parser.registerInfix(token.ASTERISK, parser.parseInfixExpression)
	parser.registerInfix(token.EQ, parser.parseInfixExpression)
	parser.registerInfix(token.NOT_EQ, parser.parseInfixExpression)
	parser.registerInfix(token.LT, parser.parseInfixExpression)
	parser.registerInfix(token.GT, parser.parseInfixExpression)

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

	// Parse expression statements (default)
	default:
		return parser.parseExpressionStatement()
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
Function to parse the identifiers.
*/
func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
}

/*
Function to parse the expression statements.
*/
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

// Function to return error if there is no prefix parse function for the current token
//
// Example:
//
// for the token "5", the prefix function is "parseIntegerLiteral"
//
// for the token "foo", the prefix function is "parseIdentifier"
//
// for the token "-" (minus) / "!" (bang), the prefix function is "parsePrefixExpression"
func (parser *Parser) noPrefixParseFnError(tokenType token.TokenType) {
	fmt.Printf("No prefix parse function for %s found\n", tokenType)
}

// Function to parse the expressions
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	// Get the prefix function for the current token
	// Example:
	// for the token "5", the prefix function is "parseIntegerLiteral"
	// for the token "foo", the prefix function is "parseIdentifier"
	// for the token "-" (minus) / "!" (bang), the prefix function is "parsePrefixExpression"
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

// Function to parse the integer literals
//
// Eg: 5; 10 + 15;
func (parser *Parser) parseIntegerLiteral() ast.Expression {
	// Create a new IntegerLiteral struct instance, set the token to the current token
	integerLiteral := &ast.IntegerLiteral{Token: parser.curToken}

	// Convert the literal value from string to int64
	value, err := strconv.ParseInt(parser.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", parser.curToken.Literal)
		parser.errors = append(parser.errors, msg)
		return nil
	}

	// Update the Value field of the IntegerLiteral struct instance
	integerLiteral.Value = value

	return integerLiteral
}

func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: parser.curToken,
		Value: parser.curTokenIs(token.TRUE),
	}
}

// Peek the precedence of the next token
func (parser *Parser) peekPrecedence() int {
	if p, ok := precedences[parser.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

// Get the precedence of the current token
func (parser *Parser) curPrecedence() int {
	if p, ok := precedences[parser.curToken.Type]; ok {
		return p
	}

	return LOWEST
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
