/**
 * UTF-8でエンコードされた文字列を表す []byte スライスの文字を、そのスライス内で逆順にするようにreverseを修正しなさい。
 * 新たなメモリを割り当てることなく行えるでしょうか。
 */
package main

import (
    "fmt"
    "unicode/utf8"
)

func reverse(bs []byte) {
    if len(bs) == 0 {
        return
    }

    _, size := utf8.DecodeRune(bs)
    reverseBytes(bs[:size])
    reverseBytes(bs[size:])
    reverseBytes(bs)
    
    reverse(bs[:len(bs) - size])
}

func reverseBytes(bs []byte) {
    for i, j := 0, len(bs) - 1; i < j; i, j = i + 1, j - 1 {
        bs[i], bs[j] = bs[j], bs[i]
    }
}

func main() {
    s := "あいうえおかきくけこ"
    bs := []byte(s)
    fmt.Println(string(bs))

    reverse(bs)
    fmt.Println(string(bs))
}
