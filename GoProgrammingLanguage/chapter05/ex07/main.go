/**
 * 汎用の HTML プリティプリンタとなるような startElement と endElement を開発しなさい。
 * コメントノード、テキストノード、個々の王その属性 (<a href='...'>) を表示しなさい。
 * 要素が子を持たない場合には、<img></img> ではなく、</img> のような短い形式を使いなさい。
 * 出力をきちんとパースできることを保証するためのテストを書きなさい (第11章を参照)。
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.TextNode {
		txt := strings.TrimSpace(html.UnescapeString(n.Data))
		if txt != "" {
			fmt.Printf("%*s%s\n", depth*2, "", txt)
		}
	} else if n.Type == html.ElementNode {
		line := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			line += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			line += "/>"
		} else {
			line += ">"
			depth++
		}
		fmt.Println(line)
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
