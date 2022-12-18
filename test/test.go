package main

import (
	"devscript/ast"
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
	var z = 10000
	foobar;
	func
	5;
	10;
	-5;
	!-5;
	!5;
	a + b + c;
	a + b - c;
	a + b / c;
	5 > 4 == 3 < 4;
	true;
	true == true;
	false == false;
	true != false;
	10 < 1 != true;
	10 > 1 == true;
	10 == 10 == true;
	10 != 10 != false;
	`
	lex := lexer.New(input)
	parser := parser.New(lex)
	program := parser.ParseProgram()
	fmt.Println("------PARSER OUTPUT-------")
	fmt.Printf("Number of statements after parsing: %d\n", len(program.Statements))
	for _, statement := range program.Statements {
		fmt.Printf("%+v\t\t%T", statement, statement)

		// check if type of statement is *ast.ExpressionStatement
		expressionStatement, ok := statement.(*ast.ExpressionStatement)
		if ok {
			fmt.Printf(" -> %T\n", expressionStatement.Expression)
		} else {
			fmt.Println()
		}
	}
}

func main() {
	testLexerPhase()
	testParserPhase()
}
