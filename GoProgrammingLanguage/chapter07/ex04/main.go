/**
 * strings.NewReader 関数は、その引数である文字列から読み込むことで io.Reader インタフェース (とほかのインタフェース) を満足する値を返します。
 * 皆さん自身で簡単な NewReader を実装し、HTML パーサ (5.2節) が文字列からの入力を受け取るようにしなさい。
 */
package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"golang.org/x/net/html"
	"./myreader"
)

func main() {
	bs, err := ioutil.ReadFile("input.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile error [%v]\n", err)
		return
	}
	r := myreader.NewReader(string(bs))
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse error [%v]\n", err)
		return
	}
	vs := visit(nil, doc)
	fmt.Println(strings.Join(vs, "\n"))
}

func visit(links []string, n *html.Node) []string {
	links = getHrefValues(links, n)
	for cur := n.FirstChild; cur != nil; cur = cur.NextSibling {
		links = visit(links, cur)
	}
	return links
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
