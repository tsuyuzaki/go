/**
 * Expr インタフェースを満足する新たな具象型を定義し、オペランドの最小値を計算するなどの新たな操作を提供しなさい。
 * Parse 関数はその新たな型のインスタンスを生成しないので、それを使うためには構文ツリーを直接構築する (あるいはパーサを拡張する) 必要があります。
 */
package main

import (
	"os"
	"fmt"
	"math"
	"./eval"
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
