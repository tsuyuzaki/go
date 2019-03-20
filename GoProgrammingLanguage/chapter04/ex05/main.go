/**
 * []stringスライス内で隣接している重複をスライス内で除去する関数を書きなさい。
 */
package main

import (
	"fmt"
)

func removeIfNextIsSame(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	ret := []string{strs[0]}
	for _, str := range strs[1:] {
		if str != ret[len(ret)-1] {
			ret = append(ret, str)
		}
	}
	return ret
}

func main() {
	s := []string{"hoge", "hoge", "foo", "hoge", "bar", "bar", "foo", "foo"}
	fmt.Println(removeIfNextIsSame(s))
}
