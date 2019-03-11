package github

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
)

func confirm(input map[string]string, msg string) bool {
    jsonstr, _ := json.MarshalIndent(input, "", "    ")
    fmt.Printf("Your input:\n%s\n\n%s (Y/N): ", jsonstr, msg)
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