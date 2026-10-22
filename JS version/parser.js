package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser { return &Parser{tokens: tokens} }



func (p *Parser) peek() Token { return p.tokens[p.pos] }

func (p *Parser) next() Token {
	t := p.tokens[p.pos]
	if t.Type != EOF { 
		p.pos++
	}
	return t
}

func (p *Parser) expect(tt TokenType) Token {
	t := p.peek()
	if t.Type != tt {
		panic(fmt.Sprintf("parse error: unexpected %q (token #%d)", t.Literal, p.pos))
	}
	return p.next()
}

// Higher number = binds tighter. 0 means "not a binary operator", kinda like BODMAS
func precedence(t TokenType) int {
	switch t {
	case ADD, SUB:
		return 1
	case STAR, DIVIDE:
		return 2
	}
	return 0
}


func (p *Parser) ParseProgram() (prog *Program, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	prog = &Program{}
	for p.peek().Type != EOF {
		prog.Stmts = append(prog.Stmts, p.parseStmt())
	}
	return prog, nil
}

func (p *Parser) parseStmt() Stmt {
	switch p.peek().Type {
	case LET:
		p.next()
		name := p.expect(IDENT).Literal
		p.expect(ASSIGN)
		val := p.parseExpr(0)
		p.expect(SEMICOLON)
		return &LetStmt{Name: name, Value: val}

	case PRINT:
		p.next()
		val := p.parseExpr(0)
		p.expect(SEMICOLON)
		return &PrintStmt{Value: val}

	case IDENT:
		name := p.next().Literal
		p.expect(ASSIGN)
		val := p.parseExpr(0)
		p.expect(SEMICOLON)
		return &AssignStmt{Name: name, Value: val}
	}
	panic(fmt.Sprintf("parse error: unexpected %q at start of statement", p.peek().Literal))
}


func (p *Parser) parseExpr(minPrec int) Expr {
	left := p.parsePrefix()

	for {
		op := p.peek()
		prec := precedence(op.Type)
		if prec == 0 || prec <= minPrec {
			break 
		}
		p.next()                    // consume the operator
		right := p.parseExpr(prec) // parse the right side at THIS operator's level
		left = &Binary{Left: left, Op: op.Literal, Right: right}
	}
	return left
}

// Things that can start an expression: literals, names, -x, (expr)
func (p *Parser) parsePrefix() Expr {
	t := p.next()
	switch t.Type {
	case NUMBER:
		v, err := strconv.ParseFloat(t.Literal, 64)
		if err != nil {
			panic(fmt.Sprintf("bad number %q", t.Literal))
		}
		return &NumberLit{Value: v}

		case STRING:
		return &StringLit{Value: t.Literal}
	case IDENT:
		return &Ident{Name: t.Literal}
	case SUB:

		return &Unary{Op: "-", Right: p.parseExpr(3)}
	case LPARAN:
		e := p.parseExpr(0)
		p.expect(RPARAN)
		return e
	}
	panic(fmt.Sprintf("parse error: unexpected %q in expression", t.Literal))
}