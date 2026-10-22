package main

import (
	"fmt"
	"strconv"
)

// Interpreter evaluates statements and expressions in the AST
type Interpreter struct {
	env map[string]any // Variable environment mapping names to values (float64 or string)
}


func NewInterpreter() *Interpreter {
	return &Interpreter{
		env: make(map[string]any),
	}
}

// Interpret executes all statements in the program
func (in *Interpreter) Interpret(prog *Program) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("runtime error: %v", r)
		}
	}()

	for _, stmt := range prog.Stmts {
		in.execStmt(stmt)
	}
	return nil
}


func (in *Interpreter) execStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *LetStmt:
		val := in.evalExpr(s.Value)
		in.env[s.Name] = val

	case *AssignStmt:
		if _, ok := in.env[s.Name]; !ok {
			panic(fmt.Sprintf("variable %q is not defined", s.Name))
		}
		val := in.evalExpr(s.Value)
		in.env[s.Name] = val

	case *PrintStmt:
		val := in.evalExpr(s.Value)
		fmt.Println(val)

	default:
		panic(fmt.Sprintf("unknown statement type: %T", stmt))
	}
}

// evalExpr evaluates an expression and returns its runtime value (float64 or string)
func (in *Interpreter) evalExpr(expr Expr) any {
	switch e := expr.(type) {
	case *NumberLit:
		return e.Value

	case *StringLit:
		return e.Value

	case *Ident:
		val, ok := in.env[e.Name]
		if !ok {
			panic(fmt.Sprintf("undefined variable %q", e.Name))
		}
		return val

	case *Unary:
		val := in.evalExpr(e.Right)
		if e.Op == "-" {
			num, ok := val.(float64)
			if !ok {
				panic(fmt.Sprintf("cannot apply unary '-' to non-number (%T)", val))
			}
			return -num
		}
		panic(fmt.Sprintf("unknown unary operator %q", e.Op))

	case *Binary:
		left := in.evalExpr(e.Left)
		right := in.evalExpr(e.Right)
		return in.evalBinary(left, e.Op, right)

	default:
		panic(fmt.Sprintf("unknown expression type: %T", expr))
	}
}


func (in *Interpreter) evalBinary(left any, op string, right any) any {
	// Case 1: Both operands are numbers
	numLeft, isLeftNum := left.(float64)
	numRight, isRightNum := right.(float64)
	if isLeftNum && isRightNum {
		switch op {
		case "+":
			return numLeft + numRight
		case "-":
			return numLeft - numRight
		case "*":
			return numLeft * numRight
		case "/":
			if numRight == 0 {
				panic("division by zero")
			}
			return numLeft / numRight
		default:
			panic(fmt.Sprintf("unknown operator %q for numbers", op))
		}
	}

	// Case 2: Both operands are strings 
	strLeft, isLeftStr := left.(string)
	strRight, isRightStr := right.(string)
	if isLeftStr && isRightStr {
		if op == "+" {
			return strLeft + strRight
		}
		panic(fmt.Sprintf("operator %q is not supported for strings (only '+' concatenation is allowed)", op))
	}
 
	if isLeftStr && isRightNum {
		if op == "+" {
			rightStr := strconv.FormatFloat(numRight,'G', -1,64)
			return strLeft + rightStr
		}
		panic(fmt.Sprintf("operator %q is not supported for strings (only '+' concatenation is allowed)", op))
	}
	if isLeftNum && isRightStr {
		if op == "+" {
			leftStr := strconv.FormatFloat(numLeft,'G', -1,64)
			return leftStr + strRight
		}
		panic(fmt.Sprintf("operator %q is not supported for strings (only '+' concatenation is allowed)", op))
	}

	panic(fmt.Sprintf("operator %q is not supported for strings (only '+' concatenation is allowed)", op))
}
