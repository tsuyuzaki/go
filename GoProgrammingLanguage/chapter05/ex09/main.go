/**
 * 文字列 s 内のそれぞれの部分文字列 "$foo" を f("foo") が返すテキストで置換する関数 expand(s string, f func(string) string) string を書きなさい。
 * $ で始まる任意の単語を探して、 $以降の文字列で関数 f を呼び出した結果のテキストです。
 */
package main

import (
    "fmt"
    "sort"
    "bufio"
    "strings"
)

func main() {
    fmt.Println(expand("$hoge $hogeg $ho", strings.ToUpper))
}

func contains(s string, ss []string) bool {
    for _, str := range ss {
        if str == s {
            return true
        }
    }
    return false
}

type targets []string

func (t targets) Len() int {
    return len(t)
}

func (t targets) Swap(i, j int) {
    t[i], t[j] = t[j], t[i]
}

func (t targets) Less(i, j int) bool {
    return t[i] > t[j]
}

func expand(s string, f func(string) string) string {
    targets := targets{}
    in := bufio.NewScanner(strings.NewReader(s))
    in.Split(bufio.ScanWords)
    for in.Scan() {
        txt := in.Text()
        if txt[0] == '$' && ! contains(txt, targets) {
            targets = append(targets, txt)
        }
    }
    sort.Sort(targets)
    for _, old := range targets {
        s = strings.Replace(s, old, f(old[1:]), -1)
    }
    return s
}
