/**
 * 文字列の連結の代わりに bytes.Buffer を使って、再起呼び出しを行わない comma を作成しなさい。
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	len := len(s)
	if len != utf8.RuneCountInString(s) {
		fmt.Fprintf(os.Stderr, "String contains not ASCII character.\n")
		return s
	}
	if len <= 3 {
		return s
	}
	var buf bytes.Buffer
	countForComma := len % 3
	for i := 0; i < len; i++ {
		buf.WriteByte(s[i])
		countForComma--
		if countForComma == 0 && i != (len-1) {
			buf.WriteByte(',')
			countForComma = 3
		}
	}
	return buf.String()
}
