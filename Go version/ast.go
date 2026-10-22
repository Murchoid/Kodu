package main

import "fmt"

type Expr interface{ fmt.Stringer }
type Stmt interface{ stmtNode() }

// Expressions
type NumberLit struct{ Value float64 }
type StringLit struct{ Value string }
type Ident struct{ Name string }
type Unary struct {
	Op    string
	Right Expr
}
type Binary struct {
	Left  Expr
	Op    string
	Right Expr
}

// Statements
type LetStmt struct {
	Name  string
	Value Expr
}
type AssignStmt struct {
	Name  string
	Value Expr
}
type PrintStmt struct{ Value Expr }

type Program struct{ Stmts []Stmt }

func (*LetStmt) stmtNode()    {}
func (*AssignStmt) stmtNode() {}
func (*PrintStmt) stmtNode()  {}


func (n *NumberLit) String() string { return fmt.Sprint(n.Value) }
func (n *StringLit) String() string { return fmt.Sprint(n.Value) }
func (n *Ident) String() string     { return n.Name }
func (n *Unary) String() string     { return fmt.Sprintf("(%s%s)", n.Op, n.Right) }
func (n *Binary) String() string    { return fmt.Sprintf("(%s %s %s)", n.Left, n.Op, n.Right) }