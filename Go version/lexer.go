package main

import (
	"fmt"
)

type Alexer struct {
	source string
	pos    int
	line   int
	start  int
}

type TokenType int

const (
	EOF TokenType = iota
	VAR
	PRINT
	FOR
	IDENT
	NUMBER
	STRING
	ADD
	SUB
	STAR
	COMMA
	DOT
	DIVIDE
	LPARAN
	RPARAN
	LBRACE
	RBRACE
	CLASS
	SUPER
	NOT
	NOT_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	AND
	ELSE
	FALSE
	FUN
	IF
	NIL
	OR
	RETURN
	THIS
	TRUE
	WHILE
	SEMICOLON
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case VAR:
		return "VAR"
	case PRINT:
		return "PRINT"
	case FOR:
		return "FOR"
	case IDENT:
		return "IDENT"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case STAR:
		return "STAR"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case DIVIDE:
		return "DIVIDE"
	case LPARAN:
		return "LPARAN"
	case RPARAN:
		return "RPARAN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case CLASS:
		return "CLASS"
	case SUPER:
		return "SUPER"
	case NOT:
		return "NOT"
	case NOT_EQUAL:
		return "NOT_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case AND:
		return "AND"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case RETURN:
		return "RETURN"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case WHILE:
		return "WHILE"
	case SEMICOLON:
		return "SEMICOLON"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	Literal string
	Type    TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %q)", t.Type, t.Literal)
}

var keywords = map[string]TokenType{
	"var":    VAR,
	"print":  PRINT,
	"fun":    FUN,
	"for":    FOR,
	"class":  CLASS,
	"super":  SUPER,
	"and":    AND,
	"else":   ELSE,
	"false":  FALSE,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"return": RETURN,
	"this":   THIS,
	"true":   TRUE,
	"while":  WHILE,
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
	}
}

func (al *Alexer) ALex() []Token {
	var tokens []Token
	for al.pos < len(al.source) {
		var token Token
		tok := al.peek(0)
		al.start = 0
		switch tok {
		// Ignore whitespace: spaces, tabs, carriage returns, and newlines
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
			if isAlphabet(tok) {
				str := al.buildIdent()
				t := LookupIdent(str)
				token = al.prepareToken(str, t)
			} else if isDigit(tok) {
				num := al.buildNumber()
				token = al.prepareToken(num, NUMBER)
			} else {
				panic(fmt.Sprintf("Unknown token: %q at pos %d", tok, al.pos))
			}
		}
		tokens = append(tokens, token)
	}

	// Add EOF token to mark the end of input
	tokens = append(tokens, al.prepareToken("EOF", EOF))

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
			panic("No closing double quote for string before newline")
		}
		al.consume()
	}

	if al.pos >= len(al.source) {
		panic("No closing double quote for string")
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
