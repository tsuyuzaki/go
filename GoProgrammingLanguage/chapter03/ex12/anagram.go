/**
 * Ex)
 * $ go run anagram.go  banana NaNABa
 * Anagram!!
 *
 * $ go run anagram.go  banana Apple
 * Not anagram
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Please pass 2 arguments.\n")
		return
	}
	lhs := toRuneCounts(toLower(os.Args[1]))
	rhs := toRuneCounts(toLower(os.Args[2]))
	if isEqual(lhs, rhs) {
		fmt.Println("Anagram!!")
	} else {
		fmt.Println("Not anagram")
	}
}

func toLower(s string) string {
	rs := []rune(s)
	var buf bytes.Buffer
	for _, r := range rs {
		buf.WriteRune(unicode.ToLower(r))
	}
	return buf.String()
}

func toRuneCounts(s string) map[rune]int {
	rs := []rune(s)
	runeCounts := make(map[rune]int)

	for _, r := range rs {
		_, ok := runeCounts[r]
		if ok {
			continue
		}
		runeCounts[r] = strings.Count(s, string(r))
	}
	return runeCounts
}

func isEqual(lhs map[rune]int, rhs map[rune]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for r, cnt := range lhs {
		rhsCnt, ok := rhs[r]
		if !ok || rhsCnt != cnt {
			return false
		}
	}
	return true
}
