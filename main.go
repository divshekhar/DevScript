package main

import (
	"devscript/src/eval"
	"devscript/src/lexer"
	"devscript/src/object"
	"devscript/src/parser"
	"devscript/src/repl"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	argsLen := len(os.Args)

	if argsLen > 1 {
		runCommand(os.Args[1])
	}

	fmt.Printf("Hello %s! This is the DevScript programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runCommand(command string) {
	switch {
	case command == "--version" || command == "-v":
		fmt.Println("DevScript version 0.0.1")
		os.Exit(0)
	case command == "--help" || command == "-h":
		fmt.Println("Usage: devscript [options]")
		fmt.Println("Options:")
		fmt.Println("--version | -v\t\tPrints the current version of DevScript")
		fmt.Println("--help | -h\t\tPrints the help message")
		os.Exit(0)
	case checkPath(command):
		runFile(command)
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s, use --help for more information", command)
		os.Exit(1)
	}
}

func checkPath(path string) bool {
	// check relative path exists or not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File not found: ", path)
		return false
	}
	return true
}

func runFile(path string) {
	// convert to absolute path
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	// check if the file has .ds extension
	if filepath.Ext(path) != ".ds" {
		fmt.Println("Invalid file extension: ", absPath)
		fmt.Println("Expected: .ds (DevScript file extension)")
		os.Exit(1)
	}

	// create a new environment
	env := object.NewEnvironment()

	// get the file content
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	// Tokenize
	lex := lexer.New(string(content))
	parser := parser.New(lex)
	program := parser.ParseProgram()

	// evaluate the program
	evaluated := eval.Eval(program, env)
	fmt.Println(evaluated.Inspect())
}
