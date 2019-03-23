/**
 * スライスの代わりにマップを使うようにtopoSortを書き直して、最初のソートを削除しなさい。
 * 結果は非決定的ですが、結果が有効なトポロジカル順序になっていることを検証しなさい。
 */
package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra":  true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math":        true},
	"databases":             {"data structures":      true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math":        true},
	"networks":              {"operating systems":    true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	var i int
	for course, _ := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
		i++
	}
}

func topoSort(m map[string]map[string]bool) map[string]bool {
	order := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[item] = true
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}