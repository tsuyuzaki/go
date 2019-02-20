package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    const (
        prefixHttp = "http://"
        prefixHttps = "https://"
    )

    for _, url := range os.Args[1:] {
        if ! strings.HasPrefix(url, prefixHttp) && ! strings.HasPrefix(url, prefixHttps) {
            url = prefixHttp + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v¥n", err)
            os.Exit(1)
        }
        b, err := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}