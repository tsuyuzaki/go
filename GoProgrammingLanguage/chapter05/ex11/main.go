/**
 * 線形代数 (Linear algebra) の講座の講師が、これからは微積分学 (calculus) を事前条件にすると決めました。
 * 循環を報告するように topoSort 関数を拡張しなさい。
 */
package main

import (
	"fmt"
	"os"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},

	"linear algebra": {"calculus": true},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "topoSort error [%v]", err)
		return
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func contains(key string, keys []string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	related := make(map[string]bool)
	var visitAll func(items map[string]bool) error

	visitAll = func(items map[string]bool) error {
		for item, _ := range items {
			if related[item] {
				return fmt.Errorf("cirtulation error [%s]", item)
			}
			if !seen[item] {
				related[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
				seen[item] = true
				related[item] = false
			}
		}
		return nil
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	err := visitAll(keys)
	return order, err
}
