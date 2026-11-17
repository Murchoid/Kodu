package main

import (
	"fmt"
)

type Alexer struct {
	source    string
	hasErrors bool
	pos       int
	line      int
	start     int
	diagnostics []Diagnostics
}

func NewLexer(source string) *Alexer {
	return &Alexer{
		source:    source,
		pos:       0,
		hasErrors: false,
		line:      0,
		start:     0,
		diagnostics: []Diagnostics{},
	}
}

func (al *Alexer) ALex() ([]Token, []Diagnostics) {
	var tokens []Token
	for al.pos < len(al.source) {
		var token Token
		tok := al.peek(0)
		al.start = 0
		switch tok {

		case ' ', '\t', '\r':
			al.consume()
			continue

		case '\n':
			al.line++
			al.consume()
			continue

		case '*':
			token = al.prepareToken("*", STAR)
			al.consume()

		case '+':
			token = al.prepareToken("+", ADD)
			al.consume()

		case '-':
			token = al.prepareToken("-", SUB)
			al.consume()

		case ',':
			token = al.prepareToken(",", COMMA)
			al.consume()

		case '.':
			token = al.prepareToken(".", DOT)
			al.consume()

		case '/':
			if al.peek(1) == '/' {
				al.consume()
				al.consume()
				for al.pos < len(al.source) && al.peek(0) != '\n' {
					al.consume()
				}
				continue
			} else if al.peek(1) == '*' {
				al.consume()
				al.consume()
				for al.pos < len(al.source) && !(al.peek(0) == '*' && al.peek(1) == '/') {
					if al.peek(0) == '\n' {
						al.line++
					}
					al.consume()
				}
				al.consume() // '*'
				al.consume() // '/'
				continue
			} else {
				token = al.prepareToken("/", DIVIDE)
				al.consume()
			}

		case '!':
			if al.peek(1) == '=' {
				al.consume()
				al.consume()
				token = al.prepareToken("!=", NOT_EQUAL)
			} else {
				token = al.prepareToken("!", NOT)
				al.consume()
			}

		case '=':
			if al.peek(1) == '=' {
				al.consume()
				al.consume()
				token = al.prepareToken("==", EQUAL_EQUAL)
			} else {
				token = al.prepareToken("=", EQUAL)
				al.consume()
			}

		case '>':
			if al.peek(1) == '=' {
				al.consume()
				al.consume()
				token = al.prepareToken(">=", GREATER_EQUAL)
			} else {
				token = al.prepareToken(">", GREATER)
				al.consume()
			}

		case '<':
			if al.peek(1) == '=' {
				al.consume()
				al.consume()
				token = al.prepareToken("<=", LESS_EQUAL)
			} else {
				token = al.prepareToken("<", LESS)
				al.consume()
			}

		case '(':
			token = al.prepareToken("(", LPARAN)
			al.consume()

		case ')':
			token = al.prepareToken(")", RPARAN)
			al.consume()

		case '{':
			token = al.prepareToken("{", LBRACE)
			al.consume()

		case '}':
			token = al.prepareToken("}", RBRACE)
			al.consume()

		case '"':
			str := al.buildString()
			token = al.prepareToken(str, STRING)

		case ';':
			token = al.prepareToken(";", SEMICOLON)
			al.consume()

		default:
			al.start = al.pos
			if isAlphabet(tok) {
				str := al.buildIdent()
				t := LookupIdent(str)
				token = al.prepareToken(str, t)
			} else if isDigit(tok) {
				num := al.buildNumber()
				token = al.prepareToken(num, NUMBER)
			} else {
				al.hasErrors = true
				al.diagnostics = AppendDiagnostics(al.diagnostics, al.line, al.source[al.start:al.pos], fmt.Sprintf("Unknown token: %q at pos %d", tok, al.pos))
				al.consume()
			}
		}
		tokens = append(tokens, token)
	}

	// Add EOF token to mark the end of input
	tokens = append(tokens, al.prepareToken("EOF", EOF))

	return tokens, al.diagnostics
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
	al.start = al.pos
	for al.pos < len(al.source) && isDigit(al.peek(0)) {
		al.consume()
	}

	numStr := al.source[al.start:al.pos]
	return numStr
}

func (al *Alexer) buildString() string {
	al.consume() // Consume the opening double quote
	al.start = al.pos

	for al.pos < len(al.source) && al.peek(0) != '"' {
		if al.peek(0) == '\n' {
			al.line++
			al.hasErrors = true
			al.diagnostics = AppendDiagnostics(al.diagnostics, al.line, al.source[al.start:al.pos], "No closing double quote for string before endline")
			return al.source[al.start:al.pos]
		}
		al.consume()
	}

	if al.pos >= len(al.source) {
		end := al.start + (al.pos-al.start)/4
		al.hasErrors = true
		al.diagnostics = AppendDiagnostics(al.diagnostics, al.line, al.source[al.start:end], "No closing double quote for string")
		return al.source[al.start:end]
	}

	str := al.source[al.start:al.pos]
	al.consume() // Consume the closing double quote
	return str
}

func (al *Alexer) buildIdent() string {
	al.start = al.pos
	for al.pos < len(al.source) && isAlphanumeric(al.peek(0)) {
		al.consume()
	}

	str := al.source[al.start:al.pos]
	return str
}

func (al *Alexer) prepareToken(lit string, tok TokenType) Token {
	token := Token{
		Literal: lit,
		Type:    tok,
	}

	return token
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
