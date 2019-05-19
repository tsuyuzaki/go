package eval

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Env map[Var]float64

