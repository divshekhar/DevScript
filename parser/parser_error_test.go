package parser

import "testing"

/*
Check if there are any errors in the parser,
if there are, print the errors and fail the test
*/
func checkParserErrors(testing *testing.T, parser *Parser) {
	// Get the errors list from the parser
	errors := parser.Errors()

	// If there are no errors, return
	if len(errors) == 0 {
		return
	}

	// print the number of errors
	testing.Errorf("parser has %d errors", len(errors))

	// print the errors
	for _, msg := range errors {
		testing.Errorf("parser error: %q", msg)
	}

	// Fail the test
	testing.FailNow()
}
