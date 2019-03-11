package github

import (
    "fmt"
    "bufio"
    "os"
)

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