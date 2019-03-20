package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Issue struct {
	URL   string
	Title string
	State string
	Body  string
}

func ShowIssue() {
	const (
		formData = `URL,https://api.github.com/repos/test4golang/test/issues
number,`
		confirmMsg = "Would you like to get issue?"
	)
	writeCSV(formData)
	input, doGet := getFixedCSVInput(confirmMsg)
	if !doGet {
		return
	}
	issue := getIssue(input)
	input, doPost := showIssue(issue)
	if doPost {
		postIssue(input)
	}
}

func getIssue(input map[string]string) *Issue {
	url := input["URL"]
	if url == "" {
		fmt.Fprintf(os.Stderr, "No URL\n")
		os.Exit(1)
	}
	url += "/" + input["number"]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get() %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "%s\n", resp.Status)
		os.Exit(1)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		fmt.Fprintf(os.Stderr, "JSON decode error [%v]\n", err)
		os.Exit(1)
	}
	return &issue
}

func showIssue(issue *Issue) (map[string]string, bool) {
	txt := fmt.Sprintf("URL,%s\n", issue.URL)
	txt += fmt.Sprintf("token,\n")
	txt += fmt.Sprintf("title,%s\n", issue.Title)
	txt += fmt.Sprintf("state,%s\n", issue.State)
	txt += fmt.Sprintf("body,%s\n", issue.Body)

	writeCSV(txt)
	const confirmMsg = "Would you like to update issue?"
	return getFixedCSVInput(confirmMsg)
}
