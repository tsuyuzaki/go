/**
 * CSS のように、名前だけではなくその属性でも要素が選択されるように xmlselect を拡張しなさい。
 * たとえば、<div id="page" class="wide"> などの要素は、その名前だけではなく一致する id や class によって選択できるようにします。
 * 
 * $ cat in.xml | go run main.go div class div
 */
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var elems []*xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elems = append(elems, &tok)
		case xml.EndElement:
			elems = elems[:len(elems)-1]
		case xml.CharData:
			names := toNames(elems)
			if containsAll(names, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

func toNames(elems []*xml.StartElement) []string {
	names := make([]string, 0)
	for _, elem := range elems {
		names = append(names, elem.Name.Local)
		for _, attr := range elem.Attr {
			names = append(names, attr.Name.Local)
		}
	}
	return names
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}