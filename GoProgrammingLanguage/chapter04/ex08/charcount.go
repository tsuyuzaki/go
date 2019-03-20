package main

/**
 * unicode.IsLetterなどの関数を使って、Unicode分類に従って文字や数字などを数えるようにcharcountを修正しなさい。
 */
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func increment(r rune, counts map[string]map[rune]int) {
	key := "others"
	if unicode.IsLetter(r) {
		key = "letters"
	} else if unicode.IsMark(r) {
		key = "marks"
	} else if unicode.IsNumber(r) {
		key = "numbers"
	} else if unicode.IsPunct(r) {
		key = "punctuation"
	} else if unicode.IsSymbol(r) {
		key = "symbols"
	} else if unicode.IsSpace(r) {
		key = "spaces"
	}

	rmap, ok := counts[key]
	if !ok {
		rmap = make(map[rune]int)
		counts[key] = rmap
	}
	rmap[r]++
}

func printCounts(counts map[string]map[rune]int) {
	for k, v := range counts {
		for c, n := range v {
			fmt.Printf("%s\t%q\t%d\n", k, c, n)
		}
	}
}

func main() {
	counts := make(map[string]map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		increment(r, counts)
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	printCounts(counts)
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
