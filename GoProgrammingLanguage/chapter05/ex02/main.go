/**
 * p、div、spanなどの要素名に、HTML ドキュメントツリー内でその要素名を持つ要素の数を対応させるマッピングを行う関数を書きなさい。
 *
 * $ go run fetch.go https://golang.org | go run main.go
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	cnt := map[string]int{}
	countElements(cnt, doc)
	fmt.Println(cnt)
}

func countElements(cnt map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		cnt[n.Data]++
	}
	for cur := n.FirstChild; cur != nil; cur = cur.NextSibling {
		countElements(cnt, cur)
	}
}
