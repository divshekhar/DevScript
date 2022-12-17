package repl

import (
	"bufio"
	"devscript/lexer"
	"devscript/parser"
	"devscript/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// lexer instance
		var lex = lexer.New(line)

		fmt.Println("------LEXER OUTPUT-------")
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
		fmt.Println()

		// lexer instance
		lex = lexer.New(line)
		// parser instance
		parser := parser.New(lex)
		// parse the program
		program := parser.ParseProgram()

		fmt.Println("------PARSER OUTPUT-------")
		fmt.Printf("Number of statements after parsing: %d\n", len(program.Statements))
		for _, statement := range program.Statements {
			fmt.Printf("%+v\t\t%T\n", statement, statement)
		}
	}
}
