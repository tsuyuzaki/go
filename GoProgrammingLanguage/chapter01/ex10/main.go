/**
 * 大量のデータを生成するウェブサイトを見つけなさい。報告される時間が大きく変化するかを調べるために fetchall を2回続けて実行して、
 * キャッシュされているかどうかを調査しなさい。毎回同じ内容を得ているでしょうか。
 * fetchall を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい。
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
