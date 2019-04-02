/**
 * countWordsAndImages を実装しなさい(単語の分割については練習問題 4.9 を参照)。
 */
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Please specify url.\n\tUsage: go run main.go https://golang.org\n")
		return
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountWordsAndImages error [%v]\n", err)
		return
	}
	fmt.Println("Words:", words, "\nImages:", images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	for cur := n; cur != nil; cur = cur.NextSibling {
		ws, is := getWordsAndImagesCount()
		words += ws
		images += is
		ws, is = countWordsAndImages(cur.FirstChild)
		words += ws
		images += is
	}
	return
}

func getWordsAndImagesCount(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			words++
		}
	}
	return
}
