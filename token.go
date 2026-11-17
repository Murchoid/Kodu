package main

import "fmt"

type TokenType int

const (
	EOF TokenType = iota

	//keywords
	VAR
	PRINT
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	RETURN
	SUPER
	THIS
	TRUE
	WHILE

	//literals
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
	LBRACE
	RBRACE
	DOT
	COMMA
	LESS    //<
	LTHAN   //<
	LEQTO   //<=
	GREATER //>
	GTHAN   //>
	GEQTO   //>=
	NOT     //!
	NOTEQ   //!=
	EQTO    //==
	SEMICOLON
)

const LET = VAR

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case VAR:
		return "VAR"
	case PRINT:
		return "PRINT"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case WHILE:
		return "WHILE"
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
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case LESS:
		return "LESS"
	case LTHAN:
		return "LTHAN"
	case LEQTO:
		return "LEQTO"
	case GREATER:
		return "GREATER"
	case GTHAN:
		return "GTHAN"
	case GEQTO:
		return "GEQTO"
	case NOT:
		return "NOT"
	case NOTEQ:
		return "NOTEQ"
	case EQTO:
		return "EQTO"
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
	"let":    VAR,
	"print":  PRINT,
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"while":  WHILE,
}