/**
 * []stringスライス内で隣接している重複をスライス内で除去する関数を書きなさい。
 */
package main

import (
    "fmt"
)

func removeIfNextIsSame(s []string) []string {
    var prev string
    isLoopFirst := true
    dupCnt := 0
    for i := len(s) - 1; i >= 0; i-- {
        if isLoopFirst {
            isLoopFirst = false
            prev = s[i]
            continue
        }
        if prev == s[i] {
           swapToLast(s, i)
           dupCnt++
        }
        prev = s[i]
    }
    return s[:len(s) - dupCnt]
}

func swapToLast(s []string, i int) {
    str := s[i]
    copy(s[i:], s[i + 1:])
    s[len(s) - 1] = str
}

func main() {
    s := []string{"hoge", "hoge", "foo", "hoge", "bar", "bar", "foo", "foo"}
    fmt.Println(removeIfNextIsSame(s))
}
