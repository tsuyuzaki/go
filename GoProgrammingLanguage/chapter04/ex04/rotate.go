/**
 * 一回のパスで操作を行うrotateを書きなさい。
 *
 * rotateはおそらくp.97の以下の実装。
 * s := []int{0, 1, 2, 3, 4, 5}
 * reverse(s[2:])
 * reverse(s[:2])
 * reverse(s)  // "[2 3 4 5 0 1]"
 */
package main

import (
	"fmt"
)

func rotate(s []int, i int) {
	ret := make([]int, len(s), cap(s))
	copy(ret, s[i:])
	copy(ret[i:], s[:i])
	copy(s, ret)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}
