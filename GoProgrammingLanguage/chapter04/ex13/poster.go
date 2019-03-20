/**
 * Open Movie DatabaseのJSONに基づくウェブサービスは、
 * https://omdbapi.com/ から映画を名前で検索し、そのポスター画像をダウンロードさせてくれます。
 * コマンドラインで指定された映画のポスター画像をダウンロードするツール poster を書きなさい。
 */
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type MovieInfo struct {
	Title  string
	Poster string
}

func getPoster(query string) bool {
	info, ok := getMovieInfo(query)
	if !ok {
		return false
	}
	return writePoster(info)
}

func getMovieInfo(query string) (*MovieInfo, bool) {
	url := "http://www.omdbapi.com/?apikey=3e43c687&t=" + query
	ch := make(chan string)
	go fetch(url, ch)

	var info MovieInfo
	if err := json.Unmarshal([]byte(<-ch), &info); err != nil {
		fmt.Fprintf(os.Stderr, "json.Unmarshal() error [%v]\n", err)
		return &info, false
	}
	return &info, true
}

func writePoster(info *MovieInfo) bool {
	ch := make(chan string)
	go fetch(info.Poster, ch)

	f, err := os.OpenFile(`poster.jpg`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.OpenFile() error [%v]\n", err)
		return false
	}
	defer f.Close()

	_, err = f.Write([]byte(<-ch))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write() error [%v]\n", err)
		return false
	}
	return true
}

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	ch <- string(buf)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input search query:")
	if !s.Scan() {
		fmt.Fprintln(os.Stderr, "Failed to input scan")
		return
	}
	if !getPoster(s.Text()) {
		fmt.Fprintln(os.Stderr, "Failed to get image")
		return
	}
	err := exec.Command(
		`C:\Windows\system32\mspaint.exe`,
		`poster.jpg`).Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to run command [%v]", err)
	}
}
