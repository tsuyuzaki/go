/**
 * 重複した行それぞれが含まれていたすべてのファイルの名前を表示するように dup2 を修正しなさい。
 *
 * 入力ファイルが Windowsの改行コード CRLF を利用していると、
 * "\n"のsplitが動作しないので注意！！
 * LFの改行の入力ファイルを利用すること。
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// https://qiita.com/daigo2010/items/d46975ad6decd8578c45
	counts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if line == "" {
				continue
			}
			value, exists := counts[line]
			if !exists {
				value = make(map[string]int)
				counts[line] = value
			}
			value[filename]++
		}
	}
	for line, filenames := range counts {
		total := 0
		fmt.Printf("%s:", line)
		for filename, cnt := range filenames {
			fmt.Printf("\t%s [%d times]\n", filename, cnt)
			total += cnt
		}
		fmt.Printf("\ttotal [%d times]\n", total)
	}
}
