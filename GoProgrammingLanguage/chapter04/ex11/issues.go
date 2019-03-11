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
    /*body := github.Issue{Title:"golang test", Body:"golang test body"}
    data, err := json.Marshal(body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v", err)
        return
    }
    fmt.Printf("%s\n", data)*/

    req, err := http.NewRequest(
        "POST", 
        "https://api.github.com/repos/test4golang/test/issues",
        bytes.NewBuffer([]byte("{\"title\":\"golang test hoge\",\"body\":\"golang test body hogehoge\"}")))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", "token 40100310077b2ac52363fff59af4504952adabf9")
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
    fmt.Println(resp.Status)
}