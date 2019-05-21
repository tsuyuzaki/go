/**
 * ウェブベースの電卓プログラムを書きなさい。
 */
package main

import (
	"os"
	"fmt"
	"log"
	"html/template"
	"net/http"
	"encoding/json"
	"./eval"
)

func getExprAndEnv(r *http.Request) (string, string, bool) {
	query := r.URL.Query()
	value, ok := query["env"]
	var env string
	if ok && len(value) == 1 { // env is optional.
		env = value[0]
	}
	
	value, ok = query["expr"]
	if !ok {
		return "", env, false
	}
	if len(value) != 1 {
		return "", env, false
	}
	return value[0], env, true
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	strexpr, strenv, ok := getExprAndEnv(r)
	if !ok {
		fmt.Fprintf(w, "Please enter the formula in the URL\nEx: http://localhost:8000/?expr=a*b&env={\"a\":2,\"b\":4} (Ex means 2 * 4.)")
		return
	}
	expr, err := eval.Parse(strexpr)
	if err != nil {
		fmt.Fprintf(w, "Invalid expr(%s). [%v]\n", strexpr, err)
		return
	}
	env := make(eval.Env)
	if strenv != "" {
		if err = json.Unmarshal([]byte(strenv), &env); err != nil {
			fmt.Fprintf(w, "Invalid env expression. err[%v]\n", err)
			return
		}
	}

	vars := make(map[eval.Var]bool)
	if err = expr.Check(vars); err != nil {
		fmt.Fprintf(w, "expr check() error. [%v]\n", err)
		return
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
