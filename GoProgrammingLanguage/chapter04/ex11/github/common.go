package github

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
)

func confirm(input map[string]string, msg string) string {
    jsonstr, _ := json.MarshalIndent(input, "", "    ")
    fmt.Printf("Your input:\n%s\n\n%s (Done/Cancel/Modify): ", jsonstr, msg)
    s := bufio.NewScanner(os.Stdin)
    if ok := s.Scan(); ! ok {
        fmt.Fprintf(os.Stderr, "Scan error\n")
        return "Cancel"
    }
    txt := s.Text()
    if txt == "Done" || txt == "Cancel" || txt == "Modify" {
        return txt
    } else {
        return confirm(input, msg)
    }
}