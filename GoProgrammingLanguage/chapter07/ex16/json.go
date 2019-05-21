/**
 * ウェブベースの電卓プログラムを書きなさい。
 */
package main

import (
	"fmt"
	"encoding/json"
	"./eval"
)
func main() {
	env := make(eval.Env)
	if err := json.Unmarshal([]byte(`{"hoge":1,"foo":2}`), &env); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(env)
}