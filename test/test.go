package main

import (
	"devscript/src/ast"
	"devscript/src/eval"
	"devscript/src/lexer"
	"devscript/src/object"
	"devscript/src/parser"
	"devscript/src/token"
	"fmt"
)

func tokenizeInput(input string) {
	lex := lexer.New(input)
	fmt.Println("------LEXER OUTPUT-------")
	for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}

func parseInput(input string) {
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

func evalInput(input string) {
	lex := lexer.New(input)
	parser := parser.New(lex)
	program := parser.ParseProgram()
	env := object.NewEnvironment()
	result := eval.Eval(program, env)
	fmt.Println("------EVALUATOR OUTPUT-------")
	fmt.Printf("Result: %+v", result.Inspect())
}

func main() {
	input := `
	var x = 5;

	func add(x, y) {
		return x + y;
	}

	if (true) {
		add(x, 5);
	} else {
		return false;
	}
	`
	// tokenizeInput(input)
	// parseInput(input)
	evalInput(input)
}
