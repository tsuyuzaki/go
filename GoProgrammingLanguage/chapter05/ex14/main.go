/**
 * 異なる構造を調べるために breadthFirst 関数を使いなさい。
 * たとえば、topoSort の例 (有向グラフ) の講座の依存関係、コンピュータ上のファイルシステムの階層、
 * 公共機関のウェブサイトからダウンロードしたパスや地下鉄の経路 (無向グラフ) のリストなどを利用できます。
 */
package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	topoSort(prereqs)
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func topoSort(m map[string][]string) {
	f := func(item string) []string {
		fmt.Println(item, ":")
		children := m[item]
		for _, child := range children {
			fmt.Println("  ", child)
		}
		return children
	}
	
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	
	breadthFirst(f, keys)
}