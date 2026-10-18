package main

import (
	"fmt"
	"strings"
)

type Alexer struct {
	source string
	pos    int
}

type TokenType int

const (
	EOF TokenType = iota
	LET
	PRINT
	IDENT
	NUMBER
	STRING
	ASSIGN
	ADD
	SUB
	STAR
	DIVIDE
	LPARAN
	RPARAN
	SEMICOLON
)


func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LET:
		return "LET"
	case PRINT:
		return "PRINT"
	case IDENT:
		return "IDENT"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case ASSIGN:
		return "ASSIGN"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case STAR:
		return "STAR"
	case DIVIDE:
		return "DIVIDE"
	case LPARAN:
		return "LPARAN"
	case RPARAN:
		return "RPARAN"
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
	"let":   LET,
	"print": PRINT,
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
		tok := al.peek(0)

		switch tok {
		// Ignore whitespace: spaces, tabs, carriage returns, and newlines
		case ' ', '\t', '\r', '\n':
			al.consume()

		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			num := al.buildNumber()
			numToken := Token{
				Literal: num,
				Type:    NUMBER,
			}
			tokens = append(tokens, numToken)

		case '*':
			starTok := Token{
				Literal: "*",
				Type:    STAR,
			}
			tokens = append(tokens, starTok)
			al.consume()

		case '+':
			plusTok := Token{
				Literal: "+",
				Type:    ADD,
			}
			tokens = append(tokens, plusTok)
			al.consume()

		case '-':
			subTok := Token{
				Literal: "-",
				Type:    SUB,
			}
			tokens = append(tokens, subTok)
			al.consume()

		case '/':
			divTok := Token{
				Literal: "/",
				Type:    DIVIDE,
			}
			tokens = append(tokens, divTok)
			al.consume()

		case '=':
			assignTok := Token{
				Literal: "=",
				Type:    ASSIGN,
			}
			tokens = append(tokens, assignTok)
			al.consume()

		case '(':
			lPTok := Token{
				Literal: "(",
				Type:    LPARAN,
			}
			tokens = append(tokens, lPTok)
			al.consume()

		case ')':
			rPTok := Token{
				Literal: ")",
				Type:    RPARAN,
			}
			tokens = append(tokens, rPTok)
			al.consume()

		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '_':
			ident := al.buildIdent()
			identTok := Token{
				Literal: ident,
				Type:    LookupIdent(ident), 
			}
			tokens = append(tokens, identTok)

		case '"':
			str := al.buildString()
			strTok := Token{
				Literal: str,
				Type:    STRING,
			}
			tokens = append(tokens, strTok)

		case ';':
			semiTok := Token{
				Literal: ";",
				Type:    SEMICOLON,
			}
			tokens = append(tokens, semiTok)
			al.consume()

		default:
			panic(fmt.Sprintf("Unknown token: %q at pos %d", tok, al.pos))
		}
	}

	// Add EOF token to mark the end of input
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

	return string(numStr)
}

func (al *Alexer) buildString() string {
	al.consume() // Consume the opening double quote

	var str strings.Builder
	for al.pos < len(al.source) && al.peek(0) != '"' {
		if al.peek(0) == '\n' {
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