/**
 * gopl.io/ch4/treesort (4.4 節) の *tree 型に対して、ツリー内の値の列を見せる String メソッドを書きなさい。
 */
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	values := []int{5,2,4,6,8,7,3,9,}
	var t *tree
	for _, value := range values {
		t = add(t, value)
	}
	fmt.Println(t)
}

func (t *tree) String() string {
	return t.toStr("")
}

func (t* tree) toStr(indent string) string {
	if t == nil {
		return ""
	}
	s := fmt.Sprintf("%s%d\n", indent, t.value)
	indent += "  "
	s += t.right.toStr(indent)
	s += t.left.toStr(indent)
	return s
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
