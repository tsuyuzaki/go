/**
 * 操作を続けるか否かを示すブーリアンの結果を pre 関数と post 関数が返すようにして、
 * それに対応するように forEachNode を修正しなさい。
 * 修正した forEachNode を使って、指定された id 属性を持つ最初の HTML 要素を見つけるような
 * 下記のシグニチャの関数 ElementByID を書きなさい。
 * ElementByID 関数は、一致が見つかったら走査を中止しなければなりません。
 *    func ElementByID(doc *html.Node, id string) *html.Node
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
	for _, id := range os.Args[1:] {
		fmt.Println(ElementByID(doc, id))
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

func ElementByID(doc *html.Node, id string) (*html.Node, bool) {
	return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) (*html.Node, bool) {
	if pre != nil {
		if !pre(n, id) {
			return n, true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret, found := forEachNode(c, id, pre, post)
		if found {
			return ret, true
		}
	}

	if post != nil {
		if !post(n, id) {
			return n, true
		}
	}
	return nil, false
}

func startElement(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return false
		}
	}
	return true
}

func endElement(n *html.Node, id string) bool {
	return true
}
