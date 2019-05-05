/**
 * sort.Interface 型は、他の利用にも適用できます。
 * 列 s が回文 (palindrome) であるか、つまり列を逆順にしても変わらないかを報告する関数 IsPalindrome(s sort.Interface) bool を書きなさい。
 * インデックス i と j の要素は、!s.Less(i, j) && !s.Less(j, i) であれば等しいと見なしなさい。
 */
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	halfL := l/2
	for i, j := 0, l-1; i < halfL; i++ {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1,2,3,2,1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1,2,3,3,2,1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1,2,3,4,2,1})))
}