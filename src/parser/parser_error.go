package parser

import (
	"devscript/src/token"
	"fmt"
)

// Returns the list of errors
func (parser *Parser) Errors() []string {
	return parser.errors
}

// Function adds a peekError to the list of errors, if the next token is not of the expected type
func (parser *Parser) peekError(nextToken token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", nextToken, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}
