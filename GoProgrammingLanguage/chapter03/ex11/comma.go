package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	if s == "" {
		fmt.Fprintf(os.Stderr, "String is empty.\n")
		return s
	}

	splitted := strings.Split(s, ".")
	len := len(splitted)
	if len > 2 {
		fmt.Fprintf(os.Stderr, "String contains multi decimal points.\n")
		return s
	}

	isStartWithSign := (s[0] == '-' || s[0] == '+')

	var commaInserded string
	if !isStartWithSign {
		commaInserded = commaBeforeDecimalPoint(splitted[0])
	} else {
		commaInserded = string(s[0]) + commaBeforeDecimalPoint(splitted[0][1:])
	}

	if len == 2 {
		commaInserded += ("." + splitted[1])
	}
	return commaInserded
}

func commaBeforeDecimalPoint(s string) string {
	len := len(s)
	if len != utf8.RuneCountInString(s) {
		fmt.Fprintf(os.Stderr, "String contains not ASCII character.\n")
		return s
	}
	countForComma := len % 3
	var buf bytes.Buffer
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
