package repl

import (
	"bufio"
	"devscript/lexer"
	"devscript/parser"
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
		lex := lexer.New(line)
		parser := parser.New(lex)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParseErrors(out, parser.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Oh you found out an error!!! \n")
	io.WriteString(out, "Parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
