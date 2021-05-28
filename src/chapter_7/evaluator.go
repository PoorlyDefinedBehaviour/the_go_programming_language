package main

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(environment Environment) float64
	Check(vars map[Var]bool) error
}

type Var string

func (v Var) Eval(environment Environment) float64 {
	return environment[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

type Literal float64

func (literal Literal) Eval(environment Environment) float64 {
	return float64(literal)
}

func (literal Literal) Check(vars map[Var]bool) error {
	return nil
}

type Unary struct {
	operator rune
	operand  Expr
}

func (unary Unary) Eval(environment Environment) float64 {
	switch unary.operator {
	case '+':
		return +unary.operand.Eval(environment)
	case '-':
		return -unary.operand.Eval(environment)
	}

	panic(fmt.Sprintf("unsupported unary operator: %q", unary.operator))
}

func (unary Unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", unary.operator) {
		return fmt.Errorf("unexpected unary operator: %q", unary.operator)
	}

	return unary.operand.Check(vars)
}

type Binary struct {
	left     Expr
	operator rune
	right    Expr
}

func (binary Binary) Eval(environment Environment) float64 {
	switch binary.operator {
	case '+':
		return binary.left.Eval(environment) + binary.right.Eval(environment)
	case '-':
		return binary.left.Eval(environment) - binary.right.Eval(environment)
	case '*':
		return binary.left.Eval(environment) * binary.right.Eval(environment)
	case '/':
		return binary.left.Eval(environment) / binary.right.Eval(environment)
	}

	panic(fmt.Sprintf("unsupported binary operator: %q", binary.operator))
}

func (binary Binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", binary.operator) {
		return fmt.Errorf("unsupported binary operator: %q", binary.operator)
	}

	if err := binary.left.Check(vars); err != nil {
		return err
	}

	if err := binary.right.Check(vars); err != nil {
		return err
	}

	return nil
}

type Call struct {
	function  string
	arguments []Expr
}

func (call Call) Eval(environment Environment) float64 {
	switch call.function {
	case "pow":
		return math.Pow(call.arguments[0].Eval(environment), call.arguments[1].Eval(environment))
	case "sin":
		return math.Sin(call.arguments[0].Eval(environment))
	case "sqrt":
		return math.Sqrt(call.arguments[0].Eval(environment))
	}

	panic(fmt.Sprintf("undefined function: %s", call.function))
}

var (
	arities = map[string]int{
		"pow":  2,
		"sin":  1,
		"sqrt": 1,
	}
)

func (call Call) Check(vars map[Var]bool) error {
	arity, ok := arities[call.function]
	if !ok {
		return fmt.Errorf("unknown function: %q", call.function)
	}

	if len(call.arguments) != arity {
		return fmt.Errorf("call to %s has %d arguments, expected %d", call.function, len(call.arguments), arity)
	}

	for _, argument := range call.arguments {
		if err := argument.Check(vars); err != nil {
			return err
		}
	}

	return nil
}

type Environment map[Var]float64

func main() {

}
