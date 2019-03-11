/**
 * コマンドラインからユーザがGitHubのイシューを作成、読み出し、更新、クローズできるツールを構築しなさい。
 * 大量のテキストを入力する必要がある場合には、ユーザの好みのテキストエディタを起動するようにしなさい。
 */
package main

import (
/*    "log"
    "time"*/

    "net/http"
    "os"
    "fmt"
    "bytes"
    "./github"
    "encoding/json"
    "io/ioutil"
    //"os/exec"
)

func main() {
    /*if err := exec.Command("EmEditor", "issues.txt").Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
    }*/

    createIssue()
}

func createIssue() {
    body := github.Issue{Title:"golang test", Body:"golang test body"}
    data, err := json.Marshal(body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v", err)
        return
    }
    fmt.Printf("%s\n", data)

    req, err := http.NewRequest(
        "POST", 
        "https://api.github.com/repos/test4golang/test/issues",
        bytes.NewBuffer([]byte("{\"title\":\"golang test\",\"body\":\"golang test body\"}")))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v", err)
        return
    }

    // req.SetBasicAuth("test4golang", "testgithub00")

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "token d6cf41526d51465950e5fa9bc29ed1073b0c1041")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(os.Stderr, "client.Do() %v", err)
        return
    }
    defer resp.Body.Close()
    
    var result github.IssuesSearchResult
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        fmt.Fprintf(os.Stderr, "Decode() %v", err)
        return
    }
    fmt.Println(ioutil.ReadAll(resp.Body))
}