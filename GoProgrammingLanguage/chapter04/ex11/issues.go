/**
 * コマンドラインからユーザがGitHubのイシューを作成、読み出し、更新、クローズできるツールを構築しなさい。
 * 大量のテキストを入力する必要がある場合には、ユーザの好みのテキストエディタを起動するようにしなさい。
 */
package main

import (
    "net/http"
    "os"
    "fmt"
    "bytes"
    "./github"
    "encoding/json"
    "bufio"
    "os/exec"
    "strings"
)

func clearNewIssueForm() {
    const (
        formData = `URL,https://api.github.com/repos/test4golang/test/issues
token,
title,
body,`
    )
    file, err := os.OpenFile(`form\NewIssue.csv`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Open err[%v]", err)
        os.Exit(1)
    }
    defer file.Close()
    
    file.Write([]byte(formData))
}

func openNewIssueForm() {
    cmd := exec.Command(
        `C:\Program Files (x86)\Microsoft Office\root\Office16\EXCEL.EXE`,
        `form\NewIssue.csv`)
    err := cmd.Start()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
        os.Exit(1)
    }
    cmd.Wait()
}

func main() {
    clearNewIssueForm()
    openNewIssueForm()
    createNewIssue()
}

func readNewIssueCSV() map[string]string {
    f, err := os.Open(`form\NewIssue.csv`)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Open() error [%v]\n", err)
        os.Exit(1)
    }
    defer f.Close()
    
    s := bufio.NewScanner(f)
    values := make(map[string]string)
    for s.Scan() {
        strs := strings.Split(s.Text(), ",")
        if len(strs) == 2 {
            values[strs[0]] = strs[1]
        }
    }
    return values
}

func confirm(input map[string]string, msg string) bool {
    fmt.Printf("Your input:\n%s\n\n%s (Y/N): ", input, msg)
    s := bufio.NewScanner(os.Stdin)
    if ok := s.Scan(); ! ok {
        fmt.Fprintf(os.Stderr, "Scan error\n")
        return false
    }
    txt := s.Text()
    if txt == "Y" {
        return true
    } else if txt == "N" {
        return false
    } else {
        return confirm(input, msg)
    }
}

func createNewIssue() {
    input := readNewIssueCSV()
    if ! confirm(input, "Would you like to create new issue?") {
        return
    }
    url := input["URL"]
    if url == "" {
        fmt.Fprintf(os.Stderr, "Invalid URL\n")
        os.Exit(1)
    }

    newIssue := github.NewIssue{Title: input["title"], Body: input["body"]}
    jsonStr, err := json.Marshal(newIssue)

    req, err := http.NewRequest(
        "POST", 
        url,
        bytes.NewBuffer([]byte(jsonStr)))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v", err)
        os.Exit(1)
    }
    req.Header.Set("Content-Type", "application/json")
    token := input["token"]
    if token == "" {
        fmt.Fprintf(os.Stderr, "No token\n")
        os.Exit(1)
    }
    req.Header.Set("Authorization", fmt.Sprintf("token %s", token))


    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(os.Stderr, "client.Do() %v", err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    
    fmt.Println(resp.Status)
}