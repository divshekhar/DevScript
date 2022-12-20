package parser

import (
	"devscript/src/ast"
	"devscript/src/lexer"
	"devscript/src/token"
)

// Parser struct represents the parser
type Parser struct {
	// Lexer instance to get token stream
	lexer *lexer.Lexer

	// Pointer to the current token
	curToken token.Token
	// Pointer to the next token
	peekToken token.Token

	// List of errors
	errors []string

	// Map of prefixParseFn functions
	// Each function is associated with a token type
	// Eg. prefixParseFns = {ADD: parsePrefixFunction, SUB: parsePrefixFunction, ...}
	prefixParseFns map[token.TokenType]prefixParseFn

	// Map of infixParseFn functions
	// Each function is associated with a token type
	// Eg. infixParseFns = {ADD: parseInfixFunction, SUB: parseInfixFunction, ...}
	infixParseFns map[token.TokenType]infixParseFn
}

type (

	// Function to parse a prefix expression.
	// Eg: -5, !true, -true, !false, -false
	// Returns an ast.Expression
	prefixParseFn func() ast.Expression

	// Function to parse a infix expression
	// Eg: 5 + 5, 5 / 5, 5 > 5, 5 == 5, 5 != 5
	// eturns an ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

// Creates a new Parser instance, takes a Lexer instance as input
func New(lex *lexer.Lexer) *Parser {
	// Create a new Parser struct instance
	parser := &Parser{
		lexer:  lex,
		errors: []string{},
	}

	// 	After the first call to [nextToken()],
	// 	curToken: 	peekToken = nil,
	// 	peekToken:  NextToken() = <First Token>
	parser.nextToken()

	// 	After the second call to [nextToken()],
	// 	curToken: 	peekToken = <First Token>,
	// 	peekToken:  NextToken() = <Second Token>
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
	parser.registerPrefix(token.LPAREN, parser.parseGroupedExpression)
	parser.registerPrefix(token.IF, parser.parseIfExpression)
	parser.registerPrefix(token.FUNCTION, parser.parseFunctionExpression)

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
	parser.registerInfix(token.LPAREN, parser.parseCallExpression)

	return parser
}

// Register a prefixParseFn function for a token type.
// Updates the prefixParseFns map.
func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

// Register a infixParseFn function for a token type.
// Updates the infixParseFns map.
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

// Function to get the next token from the lexer.
// Sets the currentToken to the peekToken and updates the peekToken to the next token from the lexer.
func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

// Function to parse the program
//
// Returns an ast.Program
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

// Function to check if the current token is of the expected type
func (parser *Parser) curTokenIs(token token.TokenType) bool {
	return parser.curToken.Type == token
}

// Function to check if the peek token is of the expected type
func (parser *Parser) peekTokenIs(token token.TokenType) bool {
	return parser.peekToken.Type == token
}

// Function to update the curToken & peekToken pointer,
// if the next token is of the expected type.
//
// Add error to the list of errors,
// when the next token is not of the expected type.
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
