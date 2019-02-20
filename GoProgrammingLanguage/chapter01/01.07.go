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

func writeBody(resp *http.Response) {
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
        writeBody(resp)
    }
}