package main

import "fmt"

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