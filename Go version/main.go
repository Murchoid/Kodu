package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("test.csk")
	if err != nil {
		log.Fatal(err)
	}

	// 1. Lexical Analysis (Tokens)
	alexer := NewLexer(string(data))
	tokens := alexer.ALex()


	DumpTokens(tokens)

	// 2. Syntax Analysis (Parser)
	parser := NewParser(tokens)
	prog, err := parser.ParseProgram()
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	// 3. Execution (Interpreter)
	interpreter := NewInterpreter()
	if err := interpreter.Interpret(prog); err != nil {
		log.Fatalf("Runtime error: %v", err)
	}
}