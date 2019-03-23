/**
 * sum のように、可変個引数関数である max と min を書きなさい。
 * 引数なしで呼び出されたら、これらの関数は何をすべきでしょうか。
 * 少なくとも一つの引数が必要なバージョンも書きなさい。
 */
package main

import (
	"fmt"
)

func join(delimiter string, vals ...string) string {
	joined := ""
	for _, val := range vals {
		if joined != "" {
			joined += delimiter
		}
		joined += val
	}
	return joined
}

func main() {
	fmt.Println(join(",", []string{}...))
	fmt.Println(join(",", []string{"あめんぼ"}...))
	fmt.Println(join(",", []string{"あめんぼ", "赤いな", "あいうえお"}...))
	fmt.Println(join("__", []string{"あめんぼ", "赤いな", "あいうえお"}...))
}
