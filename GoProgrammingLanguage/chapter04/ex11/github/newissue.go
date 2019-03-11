package github

import (
    "net/http"
    "os"
    "fmt"
    "bytes"
    "encoding/json"
    "bufio"
    "os/exec"
    "strings"
)

func PostNewIssue() {
    clearNewIssueCSV()
    input, ok := getNewIssueInput()
    clearNewIssueCSV()
    if ok {
        postNewIssue(input)
    }
}

const csvPath = `NewIssue.csv`

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

func getNewIssueInput() (map[string]string, bool) {
    openNewIssueCSV()
    input := readNewIssueCSV()
    answer := confirm(input, "Would you like to create new issue?")
    if answer == "Done" {
        return input, true
    } else if answer == "Cancel" {
        return map[string]string{}, false
    } else if answer == "Modify" {
        input = nil
        return getNewIssueInput()
    } else {
        fmt.Fprintf(os.Stderr, "Invalid user input[%s]\n", answer)
        return map[string]string{}, false
    }
}

func clearNewIssueCSV() {
    const formData = `URL,https://api.github.com/repos/test4golang/test/issues
token,
title,
body,`
    file, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.OpenFile err[%v]\n", err)
        os.Exit(1)
    }
    defer file.Close()
    
    file.Write([]byte(formData))
}

func openNewIssueCSV() {
    cmd := exec.Command(
        `C:\Program Files (x86)\Microsoft Office\root\Office16\EXCEL.EXE`,
        csvPath)
    err := cmd.Start()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
    cmd.Wait()
}

func readNewIssueCSV() map[string]string {
    f, err := os.Open(csvPath)
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