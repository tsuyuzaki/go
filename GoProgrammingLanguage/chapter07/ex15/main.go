/**
 * 標準入力から単一の式を読み込み、その式内の変数に対する値をユーザに問い合わせて、
 * それからその結果の環境のもとでその式を評価するプログラムを書きなさい。
 * すべてのエラーをきちんと処理しなさい。
 */
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"./eval"
)

func scan(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("Scan stoped")
		os.Exit(0)
	}
	return scanner.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input expr.\n")
	inputExpr := scan(scanner)
	expr, err := eval.Parse(inputExpr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eval.Parse() error. [%v]\n", err)
		return
	}
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "expr.Check() error. [%v]\n", err)
		return
	}
	
	env := make(eval.Env)
	for v, _ := range vars {
		fmt.Printf("Please input value of [%s]:", string(v))
		strF := scan(scanner)
		f, err := strconv.ParseFloat(strF, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "strconv.ParseFloat() error. str[%s] err[%v]", strF, err)
			return
		}
		env[v] = f
	}
	fmt.Printf("env is %v.\n", env)
	fmt.Printf("Expr(%s).Eval(env) is %f.\n", inputExpr, expr.Eval(env))
}
