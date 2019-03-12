/**
 * 人気があるウェブコミックxkcdはJSONインタフェースを持っています。たとえば、https://xkcd.com/571/info.0.jsonに対するリクエストは、
 * 多くのお気に入りのうちの一つであるコミック571の詳細な説明を生成します。それぞれのURLを (一度だけ!) ダウンロードして、オンラインインデックスを作成しなさい。
 * そのインデックスを使って、コマンドラインで提供された検索語と一致するコミックのそれぞれのURLと内容 (transcript) を表示するツール xkcd を書きなさい。
 */
package main

import (
    "net/http"
    "fmt"
    "os"
    "io/ioutil"
)

func fetch(ch chan<- string) {
    resp, err := http.Get("https://imgs.xkcd.com/comics/cant_sleep.png")
    if err != nil {
        ch<- fmt.Sprint(err)
        return
    }
    defer resp.Body.Close()
    
    buf, _ := ioutil.ReadAll(resp.Body)
    ch<-string(buf)
}

func main() {
    ch := make(chan string)
    go fetch(ch)
    
    f, err := os.OpenFile(`PNG.png`, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.OpenFile() error [%v]\n", err)
        os.Exit(1)
    }
    defer f.Close()

    f.Write([]byte(<-ch))
}
