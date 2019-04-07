package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Node struct {
	ID     int
	WebURL string `json:"web_url"`
}

type SharedGroup struct {
	ID          string
	GroupID     int    `json:"group_id"`
	GroupAccess int    `json:"group_access"`
	ExpiresAt   string `json:"expires_at"`
}

var gAccessValues map[string]int

func init() {
	gAccessValues = map[string]int{
		"10": 10,
		"20": 20,
		"30": 30,
		"40": 40,
		"50": 50,
	}
}

func ToStrID(path string) (string, bool) {
	if path == "" {
		fmt.Println("Invalid Path")
		return "", false
	}
	if path[0] == '/' {
		return strings.Replace(path[1:], `/`, `%2F`, -1), true
	} else {
		return strings.Replace(path, `/`, `%2F`, -1), true
	}
}

func AddGroups(token, prjURL, strGAccess string, gURLs []string) {
	gAccess, ok := gAccessValues[strGAccess]
	if !ok {
		fmt.Fprintf(os.Stderr, "Invalid group_access value. group_access must be 10, 20, 30, 40 or 50. [input is %s.]\n", strGAccess)
		return
	}
	parsedPrjURL, err := url.Parse(prjURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url.Parse error[%v]\n", err)
		return
	}
	prjID, ok := ToStrID(parsedPrjURL.Path)
	if !ok {
		return
	}

	for _, gURL := range gURLs {
		if gURL == "" {
			continue
		}
		parsedGURL, err := url.Parse(gURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http.NewRequest error[%v]\n", err)
			continue
		}
		if parsedPrjURL.Host != parsedGURL.Host {
			fmt.Fprintf(os.Stderr, "Hosts are not same [%s] [%s]\n", prjURL, gURL)
			continue
		}
		gID, ok := getGroupID(token, parsedGURL)
		if !ok {
			continue
		}
		if !addGroup(token, parsedPrjURL.Host, prjID, gID, gAccess) {
			fmt.Println("ERROR:", gURL)
		}
	}
}

func getGroupID(token string, gURL *url.URL) (int, bool) {
	strGID, ok := ToStrID(gURL.Path)
	if !ok {
		return 0, false
	}
	req, err := http.NewRequest(
		"GET",
		"https://"+gURL.Host+"/api/v4/groups/"+strGID,
		nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.NewRequest error[%v]\n", err)
		return 0, false
	}

	resp, ok := sendRequest(token, req)
	if !ok {
		return 0, false
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "error[%s]\n", resp.Status)
		return 0, false
	}

	var node Node
	if err := json.NewDecoder(resp.Body).Decode(&node); err != nil {
		fmt.Fprintf(os.Stderr, "Decode error[%v]\n", err)
		return 0, false
	}
	return node.ID, true
}

func addGroup(token, host, prjID string, gID, gAccess int) bool {
	input := SharedGroup{ID: prjID, GroupID: gID, GroupAccess: gAccess, ExpiresAt: ""}
	jsonStr, err := json.Marshal(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal error[%v]\n", err)
		return false
	}

	req, err := http.NewRequest(
		"POST",
		"https://"+host+"/api/v4/projects/"+prjID+"/share",
		bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.NewRequest error[%v]\n", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	resp, ok := sendRequest(token, req)
	if !ok {
		return false
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	fmt.Println("\t", resp.Status)
	return (resp.StatusCode == http.StatusCreated)
}

func sendRequest(token string, req *http.Request) (*http.Response, bool) {
	req.Header.Set("Private-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.Do error[%v]\n", err)
		return nil, false
	}
	return resp, true
}
