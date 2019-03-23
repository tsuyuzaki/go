/**
 * HTML ノードツリーと 0 個以上の名前が与えられたら、
 * それらの名前の一つと一致する要素をすべて返す可変個引数関数 ElementsByTagName を書きなさい。
 * 二つの呼び出し例を次に示します。
 * 
 * func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
 * 
 * images := ElementsByTagName(doc, "img")
 * headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	doc, err := getHTMLDoc("https://golang.org/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getHTMLDoc() error [%v]\n", err)
		return
	}
	
	nodes := ElementsByTagName(doc, "a", "div", "span")
	for _, node := range nodes {
		fmt.Println(node)
	}
}

func getHTMLDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	
	var forEachNode func(n *html.Node, name ...string)
	forEachNode = func(n *html.Node, name ...string) {
		if n.Type == html.ElementNode && contains(n.Data, name...) {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, name...)
		}
	}
	forEachNode(n, name...)

	return nodes
}

func contains(name string, names ...string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}