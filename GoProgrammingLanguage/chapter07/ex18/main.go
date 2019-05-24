/**
 * トークンに基づくデコーダの API を使って、任意の XML ドキュメントを読み込んで、
 * そのドキュメントを表す総称的なノードのツリーを構築するプログラムを書きなさい。
 * ノードには二種類あり、CharData ノードはテキスト文字列を表し、
 * Element ノードは名前付き要素とその属性を表します。それぞれの要素のノードは子ノードのスライスを持ちます。
 * 　次の宣言が役立つでしょう。
 * 
 * import "encoding/xml"
 * 
 * type Node interface{} // CharData あるいは *Element
 * 
 * type CharData string
 * 
 * type Element struct {
 *     Type     xml.Name
 *     Attr     []xml.Attr
 *     Children []Node
 * }
 * 
 * 
 * $ cat in.xml | go run main.go
 */
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData あるいは *Element

type CharData string

type Element struct {
    Type     xml.Name
    Attr     []xml.Attr
    Children []Node
}

func (e *Element) toStr(gen int) string {
	s := fmt.Sprintf("%*s<%s", gen*2, "", e.Type.Local)
	for _, attr := range e.Attr {
		s += fmt.Sprintf(` %s="%s"`, attr.Name.Local, attr.Value)
	}
	s += ">"
	
	for _, child := range e.Children {
		switch child := child.(type) {
		case CharData:
			s += string(child)
		case *Element:
			s += "\n" + child.toStr(gen+1) + "\n"
			s += fmt.Sprintf("%*s", gen*2, "")
		}
	}

	s += fmt.Sprintf("</%s>", e.Type.Local)
	return s
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	root := &Element{Type: xml.Name{Local:"root"}}
	stack := []*Element{root}
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
			elem := &Element{Type: tok.Name, Attr: tok.Attr}
			parent := stack[len(stack) - 1]
			parent.Children = append(parent.Children, Node(elem))
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			data := CharData(string(tok))
			parent := stack[len(stack) - 1]
			parent.Children = append(parent.Children, Node(data))
		}
	}

	fmt.Println(root.toStr(0))
}
