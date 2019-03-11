package github

import (
    "net/http"
    "os"
    "fmt"
    "bytes"
    "encoding/json"
)

func PostNewIssue() {
    const (
        newIssueCSVPath = `NewIssue.csv`
        formData = 
`URL,https://api.github.com/repos/test4golang/test/issues
token,
title,
body,`
        confirmMsg = "Would you like to create new issue?"
    )
    clearCSV(newIssueCSVPath, formData)
    input, ok := getFixedCSVInput(newIssueCSVPath, confirmMsg)
    if ok {
        postNewIssue(input)
    }
}

type newIssue struct {
    url     string
    token   string
    jsonStr []byte
}

func createNewIssue(input map[string]string) *newIssue {
    url := input["URL"]
    if url == "" {
        fmt.Fprintf(os.Stderr, "No URL\n")
        return nil
    }
    token, ok := input["token"]
    if token == "" {
        fmt.Fprintf(os.Stderr, "No token\n")
        return nil
    }
    if ok {
        delete(input, "token")
    }
    
    jsonStr, err := json.Marshal(input)
    if err != nil {
        fmt.Fprintf(os.Stderr, "json.Marshal err[%v]\n", err)
        return nil
    }
    
    return &newIssue{url: url, 
        token: fmt.Sprintf("token %s", token), 
        jsonStr: jsonStr}
}

func postNewIssue(input map[string]string) {
    issue := createNewIssue(input)
    if issue == nil {
        os.Exit(1)
    }
    req, err := http.NewRequest(
        "POST", 
        issue.url,
        bytes.NewBuffer([]byte(issue.jsonStr)))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.NewRequest() %v\n", err)
        os.Exit(1)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", issue.token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(os.Stderr, "client.Do() %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    
    fmt.Println(resp.Status)
}