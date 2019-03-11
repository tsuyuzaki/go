package github

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "os/exec"
    "encoding/json"
)

func getFixedCSVInput(csvPath, confirmMsg string) (map[string]string, bool) {
    openCSV(csvPath)
    input := readCSV(csvPath)
    answer := confirm(input, confirmMsg)
    if answer == "Modify" {
        input = nil
        return getFixedCSVInput(csvPath, confirmMsg)
    }
    
    clearCSV(csvPath, "")
    if answer != "Done" {
        input = map[string]string{}
    }
    return input, (answer == "Done")
}

func clearCSV(csvPath, formData string) {
    f, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.OpenFile err[%v]\n", err)
        os.Exit(1)
    }
    defer f.Close()
    
    f.Write([]byte(formData))
}

func openCSV(csvPath string) {
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

func readCSV(csvPath string) map[string]string {
    f, err := os.Open(csvPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Open() error [%v]\n", err)
        os.Exit(1)
    }
    defer f.Close()
    
    s := bufio.NewScanner(f)
    input := make(map[string]string)
    for s.Scan() {
        strs := strings.Split(s.Text(), ",")
        if len(strs) == 2 {
            input[strs[0]] = strs[1]
        }
    }
    return input
}

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