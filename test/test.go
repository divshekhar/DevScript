package main

import (
	"devscript/lexer"
	"devscript/parser"
	"devscript/token"
	"fmt"
)

func testLexerPhase() {
	input := "var x = 5;"
	lex := lexer.New(input)
	fmt.Println("------LEXER OUTPUT-------")
	for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}

func testParserPhase() {
	input := `
	var x = 5;
	var y = 10;
	var z = 10000;
	foobar;
	func
	`
	lex := lexer.New(input)
	parser := parser.New(lex)
	program := parser.ParseProgram()
	fmt.Println("------PARSER OUTPUT-------")
	fmt.Printf("Number of statements after parsing: %d\n", len(program.Statements))
	for _, statement := range program.Statements {
		fmt.Printf("%+v\t%T\n", statement, statement)
	}
}

func main() {
	testLexerPhase()
	testParserPhase()
}
