package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func httpGet(url string) *http.Response {
    const (
        prefixHttp = "http://"
        prefixHttps = "https://"
    )
    if ! strings.HasPrefix(url, prefixHttp) && ! strings.HasPrefix(url, prefixHttps) {
        url = prefixHttp + url
    }
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v¥n", err)
        os.Exit(1)
    }
    return resp
}

func writeBody(resp *http.Response) {
    _, err := io.Copy(os.Stdout, resp.Body)
    fmt.Printf("\n\nStatusCode [%d]\n", resp.StatusCode)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    for _, url := range os.Args[1:] {
        resp := httpGet(url)
        writeBody(resp)
    }
}