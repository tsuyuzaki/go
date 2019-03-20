/**
 * HTML ドキュメントツリー内のすべてのテキストノードの内容を表示する関数を書きなさい。
 * ウェブブラウザでは内容が表示されない <script> 要素と <style> 要素の中はしらべないようにしなさい。
 *
 * $ go run fetch.go https://golang.org | go run main.go
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func findFirstStartTag(z *html.Tokenizer) (html.Token, bool) {
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			return z.Token(), true
		} else if tt == html.ErrorToken {
			return html.Token{}, false
		}
	}
}

func getTexts(z *html.Tokenizer) []string {
	txts := []string{}
	tag, ok := findFirstStartTag(z)
	if !ok {
		fmt.Println("Start tag not found.")
		return txts
	}
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return txts
		} else if tt == html.StartTagToken {
			tag = z.Token()
		} else if tt == html.TextToken {
			if tag.Data == "script" || tag.Data == "style" {
				continue
			}
			txt := strings.TrimSpace(html.UnescapeString(string(z.Text())))
			if txt != "" {
				txts = append(txts, txt)
			}
		}
	}
}

func main() {
	z := html.NewTokenizer(os.Stdin)
	fmt.Println("-------")
	fmt.Println(strings.Join(getTexts(z), "\n-------\n"))
}
