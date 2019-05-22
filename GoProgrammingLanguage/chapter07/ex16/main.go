/**
 * ウェブベースの電卓プログラムを書きなさい。
 */
package main

import (
	"./eval"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func calcHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	strExpr := q.Get("expr")
	if strExpr == "" {
		fmt.Fprintf(w, "Please enter the formula in the URL\nEx: http://localhost:8000/?expr=a*b&env={\"a\":2,\"b\":4} (Ex means 2 * 4.)")
		return
	}
	expr, err := eval.Parse(strExpr)
	if err != nil {
		fmt.Fprintf(w, "Invalid expr(%s). [%v]\n", strExpr, err)
		return
	}

	vars := make(map[eval.Var]bool)
	if err = expr.Check(vars); err != nil {
		fmt.Fprintf(w, "expr check() error. [%v]\n", err)
		return
	}

	env := make(eval.Env)
	strEnv := q.Get("env")
	if strEnv != "" { // env is optional.
		if err = json.Unmarshal([]byte(strEnv), &env); err != nil {
			fmt.Fprintf(w, "Invalid env expression. err[%v]\n", err)
			return
		}
	}

	result := expr.Eval(env)
	tpl := template.Must(template.ParseFiles("./template/index.html.tpl"))
	if err := tpl.Execute(w, result); err != nil {
		fmt.Fprintf(os.Stderr, "tpl.Execute error [%v]\n", err)
	}
}

func main() {
	http.HandleFunc("/", calcHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
