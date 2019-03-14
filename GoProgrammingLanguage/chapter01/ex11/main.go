/**
 * alexa.com にある上位100万件のウェブサイトのように、より長い引数リストで fetchall を試しなさい。
 * あるウェブサイトが応答しない場合には、プログラムはどのように振る舞うでしょうか。
 * (そのような場合に対処するための機構は8.9節で説明されています。)
 * 
 * main.exe https://www.alexa.com/siteinfo/google.com https://www.alexa.com/siteinfo/youtube.com https://www.alexa.com/siteinfo/facebook.com https://www.alexa.com/siteinfo/baidu.com https://www.alexa.com/siteinfo/wikipedia.org https://www.alexa.com/siteinfo/qq.com https://www.alexa.com/siteinfo/tmall.com https://www.alexa.com/siteinfo/yahoo.com https://www.alexa.com/siteinfo/taobao.com https://www.alexa.com/siteinfo/amazon.com
 * 以下、実行結果。

1.66s  106612 https://www.alexa.com/siteinfo/tmall.com
1.82s  123935 https://www.alexa.com/siteinfo/amazon.com
1.88s  110013 https://www.alexa.com/siteinfo/baidu.com
1.88s  111094 https://www.alexa.com/siteinfo/taobao.com
1.88s  108488 https://www.alexa.com/siteinfo/qq.com
1.88s  122354 https://www.alexa.com/siteinfo/yahoo.com
2.00s  121712 https://www.alexa.com/siteinfo/wikipedia.org
2.32s  118081 https://www.alexa.com/siteinfo/youtube.com
2.59s  125821 https://www.alexa.com/siteinfo/google.com
4.26s  130671 https://www.alexa.com/siteinfo/facebook.com
%.2fs elapsed
 4.2589867

並列で処理が行われている。
*/
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    file, err := os.Create(`log.txt`)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Create: %v¥n", err)
        os.Exit(1)
    }
    defer file.Close()

    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }
    for range os.Args[1:] {
        fmt.Fprintln(file, <-ch)
    }
    fmt.Fprintln(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
