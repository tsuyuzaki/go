package eval

import (
	"fmt"
)

type postfix struct {
	x  Expr
	op string
}

func (p postfix) Eval(env Env) float64 {
	switch p.op {
	case "++":
		return (p.x.Eval(env) + 1)
	case "--":
		return (p.x.Eval(env) - 1)
	}
	panic(fmt.Sprintf("unsupported postfix operator: %s", p.op))
}

func (p postfix) Check(vars map[Var]bool) error {
	if p.op != "++" && p.op != "--" {
		return fmt.Errorf("unexpected postfix op %s", p.op)
	}
	return p.x.Check(vars)
}

func (p postfix) String() string {
	return p.x.String() + p.op
}