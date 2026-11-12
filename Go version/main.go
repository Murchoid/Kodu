package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("../test.csk")
	if err != nil {
		log.Fatal(err)
	}

	// 1. Lexical Analysis (Tokens)
	alexer := NewLexer(string(data))
	tokens := alexer.ALex()

	DumpTokens(tokens)

}
