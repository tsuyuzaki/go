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
    "bufio"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type XKCD struct {
    Month      string `json:"month"`
    Num        int
    Link       string
    Year       string
    SafeTitle  string `json:"safe_title"`
    Transcript string
    Alt        string
    Img        string
    Title      string
    Day        string
}

func fetch(url string, ch chan<- string) {
    resp, err := http.Get(url)
    if err != nil {
        ch<- fmt.Sprint(err)
        return
    }
    defer resp.Body.Close()
    
    buf, _ := ioutil.ReadAll(resp.Body)
    ch<-string(buf)
}

func getXKCDURL(i int) string {
    return fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
}

func getXKCDs() []*XKCD {
    const maxcnt = 1000
    
    ch := make(chan string)
    for i := 1; i < maxcnt; i++ {
        url := getXKCDURL(i)
        go fetch(url, ch)
    }
    xkcds := []*XKCD{}
    for i := 1; i < maxcnt; i++ {
        var xkcd XKCD
        if err := json.Unmarshal([]byte(<-ch), &xkcd); err != nil {
            fmt.Fprintf(os.Stderr, "json.Unmarshal() error [%v]\n", err)
            continue
        }
        xkcds = append(xkcds, &xkcd)
    }
    return xkcds
}

func search(query string, xkcds []*XKCD) {
    if query == "" {
        return
    }
    for _, xkcd := range xkcds {
        if ! strings.Contains(xkcd.Title, query) {
            continue
        }
        fmt.Println("--------\nTITLE:[", xkcd.Title, "]\nURL:", 
            getXKCDURL(xkcd.Num), "\nTRANSCRIPT:", xkcd.Transcript, "\n--------")
    }
}

func main() {
    xkcds := getXKCDs()
    
    s := bufio.NewScanner(os.Stdin)
    fmt.Println("Please input search query:")
    for s.Scan() {
        search(s.Text(), xkcds)
    }
}
