/**
 * ループの代わりに visit への再起呼び出しを使って n.FirstChild リンクリストを走査するように findlinks プログラムを変更しなさい。
 *
 * $ go run fetch.go https://golang.org | go run findlinks.go
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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func getHrefValues(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return links
}

func visit(links []string, n *html.Node) []string {
	links = getHrefValues(links, cur)
	for cur := n; cur != nil; cur = cur.NextSibling {
		links = visit(links, cur.FirstChild)
	}
	return links
}
