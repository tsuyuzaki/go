/**
 * > go test -v .
 */
package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"5 / 9 * (F++ - 1)", Env{"F": 9}, "5"},
		{"i++", Env{"i": 212}, "213"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}

		s := expr.String()
		fmt.Println(s)
		expr, err = Parse(s)
		if err != nil {
			t.Error(err)
			continue
		}
		got = fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s(orig:%s).Eval() in %v = %q, want %q\n",
				s, test.expr, test.env, got, test.want)
		}
	}
}
