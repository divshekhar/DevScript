package parser

import (
	"devscript/ast"
	"fmt"
	"strconv"
)

// Function to parse the integer literals
//
//	5;
//	10 + 15;
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
