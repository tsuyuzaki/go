/**
 * UTF-8でエンコードされた []byte スライス内で隣接しているUnicodeスペース
 * (unicode.IsSpaceを参照) を、もとのスライス内で一つのASCIIスペースへ圧縮する関数を書きなさい。
 */
package main

import (
    "fmt"
    "bytes"
    "unicode"
)

func removeIfNextIsUnicodeSpace(bs []byte) []byte {
    runes := bytes.Runes(bs)
    if len(runes) == 0 {
        return bs
    }
    
    prev := runes[0]
    var buf bytes.Buffer
    buf.WriteRune(prev)
    for _, r := range runes[1:] {
        if ! unicode.IsSpace(r) {
            buf.WriteRune(r)
        } else { // r is unicode space.
            if ! unicode.IsSpace(prev) {
                buf.WriteByte(' ')
            }
        }
        prev = r
    }
    return buf.Bytes()
}

func main() {
    s := "あめんぼ  赤いな　 　あいうえお"
    bs := []byte(s)
    fmt.Println(string(removeIfNextIsUnicodeSpace(bs)))
}
