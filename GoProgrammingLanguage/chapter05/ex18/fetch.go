/**
 * 振る舞いを変えることなく、書き込み可能なファイルを閉じるために defer を使うよう fetch を書き直しなさい。
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`Please specify 1 url
  Usage: go run fetch.go https://golang.org/`)
		return
	}
	path, err := fetch(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}
}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get(%s) error [%v]\n", url, err)
		return "", err
	}
	defer resp.Body.Close()

	dirpath := "./result" + resp.Request.URL.Path
	if err := os.MkdirAll(dirpath, 0777); err != nil {
		fmt.Fprintf(os.Stderr, "os.MkdirAll(%s) error [%v]\n", dirpath, err)
		return dirpath, err
	}
	path := dirpath + "/index.html"
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Create(%s) error [%v]\n", path, err)
		return path, err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy() error [%v]\n", err)
		return path, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Close() error [%s]\n", err)
		}
	}()

	return path, err
}
