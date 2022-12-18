package parser

import (
	"devscript/ast"
	"devscript/token"
)

// Function to parse the boolean literals
//
// Eg:
//
//	true; false;
func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: parser.curToken,
		Value: parser.curTokenIs(token.TRUE),
	}
}
