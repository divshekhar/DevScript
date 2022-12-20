package repl

import (
	"bufio"
	"devscript/src/eval"
	"devscript/src/lexer"
	"devscript/src/object"
	"devscript/src/parser"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// New environment
	env := object.NewEnvironment()

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

		evaluatedResult := eval.Eval(program, env)
		if evaluatedResult != nil {
			io.WriteString(out, evaluatedResult.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Oh you found out an error!!! \n")
	io.WriteString(out, "Parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
