/**
 * 関数呼び出し io.Copy(dst, src) は、src から読み込み dst へ書き込みます。
 * ストリーム全体を保持するのに十分な大きさのバッファを要求することなくレスポンスの内容を os.Stdout へコピーするために、
 * ioutil.ReadAll の代わりにその関数を使いなさい。ない、io.Copy のエラー結果は必ず検査するようにしなさい。
 */
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)


func httpGet(url string) *http.Response {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v¥n", err)
        os.Exit(1)
    }
    return resp
}

func writeAndCloseResponse(resp *http.Response) {
    _, err := io.Copy(os.Stdout, resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    for _, url := range os.Args[1:] {
        resp := httpGet(url)
        writeAndCloseResponse(resp)
    }
}
