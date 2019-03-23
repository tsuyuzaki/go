/**
 * strings.Join の可変個引数のバージョンを書きなさい。
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
