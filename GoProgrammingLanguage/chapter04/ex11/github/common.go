package github

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const csvPath = `Issue.csv`

type postData struct {
	url     string
	token   string
	jsonStr []byte
}

func postIssue(input map[string]string) {
	postData := createPostData(input)
	if postData == nil {
		os.Exit(1)
	}
	req, err := http.NewRequest(
		"POST",
		postData.url,
		bytes.NewBuffer([]byte(postData.jsonStr)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.NewRequest() %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", postData.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.Do() %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func createPostData(orgInput map[string]string) *postData {
	input := make(map[string]string)
	token := ""
	for k, v := range orgInput {
		if k != "token" {
			input[k] = v
		} else {
			token = v
		}
	}

	ok := true
	for token == "" {
		token, ok = ScanText("Please input token: ")
		if !ok {
			return nil
		}
	}

	url := input["URL"]
	if url == "" {
		fmt.Fprintf(os.Stderr, "No URL\n")
		return nil
	}

	jsonStr, err := json.Marshal(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal err[%v]\n", err)
		return nil
	}

	return &postData{url: url,
		token:   fmt.Sprintf("token %s", token),
		jsonStr: jsonStr}
}

func getFixedCSVInput(confirmMsg string) (map[string]string, bool) {
	openCSV()
	input := readCSV()
	answer := confirm(input, confirmMsg)
	if answer == "Modify" {
		input = nil
		return getFixedCSVInput(confirmMsg)
	}

	if err := os.Remove(csvPath); err != nil {
		fmt.Fprintf(os.Stderr, "os.Remove err[%v]\n", err)
		os.Exit(1)
	}
	if answer != "Yes" {
		input = map[string]string{}
	}
	return input, (answer == "Yes")
}

func writeCSV(formData string) {
	f, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.OpenFile err[%v]\n", err)
		os.Exit(1)
	}
	defer f.Close()

	f.Write([]byte(formData))
}

func openCSV() {
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

func readCSV() map[string]string {
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

func confirm(input map[string]string, confirmMsg string) string {
	jsonstr, _ := json.MarshalIndent(input, "", "    ")
	msg := fmt.Sprintf("Your input:\n%s\n\n%s (Yes/No/Modify): ", jsonstr, confirmMsg)
	txt, ok := ScanText(msg)
	if !ok {
		return "No"
	}
	if txt != "Yes" && txt != "No" && txt != "Modify" {
		return confirm(input, confirmMsg)
	}
	return txt
}

func ScanText(msg string) (string, bool) {
	fmt.Printf("%s", msg)
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "Scan error\n")
		return "", false
	}
	txt := s.Text()
	return txt, true
}
