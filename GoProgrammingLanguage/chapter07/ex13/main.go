/**
 * 構文ツリーを綺麗に表示する String メソッドを Expr に追加しなさい。
 * その結果を再び解析したら同等のツリーになることを検査しなさい。
 */
package main

import (
	"./eval"
	"fmt"
	"math"
	"os"
)

func main() {
	expr, err := eval.Parse("sqrt(A / pi)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "eval.Parse() error. [%v]\n", err)
		return
	}
	s := expr.String()
	fmt.Println(s)

	reExpr, err := eval.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eval.Parse() error. [%v]\n", err)
		return
	}

	env := eval.Env{"A": 87616, "pi": math.Pi}
	fmt.Printf("%f, %f", expr.Eval(env), reExpr.Eval(env))
}
