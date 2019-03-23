/**
 * 振る舞いを変えることなく、書き込み可能なファイルを閉じるために defer を使うよう fetch を書き直しなさい。
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v促n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
