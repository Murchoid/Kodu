package main

import (
	"fmt"
	"strings"
)

type Alexer struct {
	source string
	line    int
	start    int
	pos    int
}


func LookupIdent(ident string) TokenType {
	if tokType, ok := keywords[ident]; ok {
		return tokType
	}
	return IDENT
}

func NewLexer(source string) *Alexer {
	return &Alexer{
		source: source,
		pos:    0,
		line:    0,
		start:    0,
	}
}

func (al *Alexer) ALex() []Token {
	var tokens []Token
	for al.pos < len(al.source) {
		tok := al.peek(0)
		var thisToken Token
		switch tok {
		// Ignore whitespace: spaces, tabs, carriage returns, and newlines
		case ' ', '\t', '\r', '\n':
			if al.peek(0) == '\n' {
				al.line++
			}
			al.consume()
			continue

		case '*':
			thisToken = al.createAndConsumeToken("*", STAR)

		case '+':
			thisToken = al.createAndConsumeToken("+", ADD)

		case '-':
			thisToken = al.createAndConsumeToken("-", SUB)

		case '/':
			if al.peek(1) == '/' {
				for al.pos < len(al.source) && al.peek(0) != '\n' {
					al.consume()
				}
				al.line++
				continue
			} else if al.peek(1) == '*' {
				for al.pos < len(al.source) {
					if al.peek(0) == '*' && al.peek(1) == '/' {
						al.consume()
						al.consume()
						break
					}
					al.consume()
				}
				al.consume() // the last /
				continue
			}
			thisToken = al.createAndConsumeToken("/", DIVIDE)

		case '=':
			if al.peek(1) == '=' {
				al.consume()
				thisToken = al.createAndConsumeToken("==", EQTO)
			} else {
				thisToken = al.createAndConsumeToken("=", ASSIGN)
			}

		case '!':
			if al.peek(1) == '=' {
				al.consume()
				thisToken = al.createAndConsumeToken("!=", NOTEQ)
			} else {
				thisToken = al.createAndConsumeToken("!", NOT)
			}

		case '<':
			if al.peek(1) == '=' {
				al.consume()
				thisToken = al.createAndConsumeToken("<=", LEQTO)
			} else {
				thisToken = al.createAndConsumeToken("<", LESS)
			}

		case '>':
			if al.peek(1) == '=' {
				al.consume()
				thisToken = al.createAndConsumeToken(">=", GEQTO)
			} else {
				thisToken = al.createAndConsumeToken(">", GREATER)
			}

		case '(':
			thisToken = al.createAndConsumeToken("(", LPARAN)

		case ')':
			thisToken = al.createAndConsumeToken(")", RPARAN)

		case '{':
			thisToken = al.createAndConsumeToken("{", LBRACE)

		case '}':
			thisToken = al.createAndConsumeToken("}", RBRACE)

		case '.':
			thisToken = al.createAndConsumeToken(".", DOT)

		case ',':
			thisToken = al.createAndConsumeToken(",", COMMA)

		case ';':
			thisToken = al.createAndConsumeToken(";", SEMICOLON)

		case '"':
			str := al.buildString()
			thisToken = Token{Literal: str, Type: STRING}

		default:
			if isDigit(tok) {
				num := al.buildNumber()
				thisToken = Token{Literal: num, Type: NUMBER}
			} else if isAlphabet(tok) {
				ident := al.buildIdent()
				thisToken = Token{Literal: ident, Type: LookupIdent(ident)}
			} else {
				panic(fmt.Sprintf("Unknown token: %q at pos %d", tok, al.pos))
			}
		}
		tokens = append(tokens, thisToken)
	}

	// to mark the end of input
	eofTok := Token{
		Literal: "EOF",
		Type:    EOF,
	}
	tokens = append(tokens, eofTok)

	return tokens
}

// Helpers
func (al *Alexer) peek(idx int) byte {
	if al.pos+idx >= len(al.source) || al.pos+idx < 0 {
		return 0
	}
	return al.source[al.pos+idx]
}

func (al *Alexer) next() {
	if al.pos < len(al.source) {
		al.pos++
	}
}

func (al *Alexer) consume() byte {
	if al.pos >= len(al.source) {
		return 0
	}
	token := al.source[al.pos]
	al.next()
	return token
}

func (al *Alexer) buildNumber() string {
	var numStr []byte
	for al.pos < len(al.source) && isDigit(al.peek(0)) {
		numStr = append(numStr, al.consume())
	}
	if al.peek(0) == '.' && isDigit(al.peek(1)) {
		numStr = append(numStr, al.consume())
		for al.pos < len(al.source) && isDigit(al.peek(0)) {
			numStr = append(numStr, al.consume())
		}
	}

	return string(numStr)
}

func (al *Alexer) buildString() string {
	al.consume() // Consume the opening double quote

	var str strings.Builder
	for al.pos < len(al.source) && al.peek(0) != '"' {
		if al.peek(0) == '\n' {
			al.line++
			panic("No closing double quote for string before newline")
		}
		str.WriteByte(al.consume())
	}

	if al.pos >= len(al.source) {
		panic("No closing double quote for string")
	}

	al.consume() // Consume the closing double quote
	return str.String()
}

func (al *Alexer) buildIdent() string {
	var str strings.Builder

	for al.pos < len(al.source) && isAlphanumeric(al.peek(0)) {
		str.WriteByte(al.consume())
	}

	return str.String()
}

func (al *Alexer) createAndConsumeToken(literal string, tokType TokenType) Token {
	tok := Token{
		Literal: literal,
		Type:    tokType,
	}
	al.consume()

	return tok
}

func isDigit(tok byte) bool {
	return tok >= '0' && tok <= '9'
}

func isAlphabet(tok byte) bool {
	return (tok >= 'a' && tok <= 'z') || (tok >= 'A' && tok <= 'Z') || tok == '_'
}

func isAlphanumeric(tok byte) bool {
	return isDigit(tok) || isAlphabet(tok)
}

func DumpTokens(tokens []Token) {
	fmt.Printf("--- Tokens Dump (Total: %d) ---\n", len(tokens))
	for i, tok := range tokens {
		fmt.Printf("[%02d] %-12s %q\n", i, tok.Type, tok.Literal)
	}
	fmt.Println("------------------------------")
}

func (al *Alexer) DumpTokens(tokens []Token) {
	DumpTokens(tokens)
}
