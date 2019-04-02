/**
 * visit 関数を拡張して、画像、スクリプト、スタイルシートなどの他の種類のリンクをドキュメントから抽出するようにしなさい。
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	vs := visit(nil, doc)
	fmt.Println(strings.Join(vs, "\n"))
}

func getHrefValues(vs []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			vs = getValues(vs, "href", n)
		case "script", "img":
			vs = getValues(vs, "src", n)
		}
	}
	return vs
}

func getValues(vs []string, k string, n *html.Node) []string {
	for _, a := range n.Attr {
		if a.Key == k {
			vs = append(vs, a.Val)
		}
	}
	return vs
}

func visit(links []string, n *html.Node) []string {
	links = getHrefValues(links, n)
	for cur := n.FirstChild; cur != nil; cur = cur.NextSibling {
		links = visit(links, cur)
	}
	return links
}
