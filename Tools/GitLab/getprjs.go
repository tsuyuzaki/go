package main

import (
	"./gitlab"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type GroupInfo struct {
	SharedProjects []*gitlab.Node `json:"shared_projects"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input your token and GroupURL.")
		return
	}
	url, err := url.Parse(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "url.Parse error [%v]", err)
		return
	}

	prjID, ok := gitlab.ToStrID(url.Path)
	if !ok {
		return
	}
	req, err := http.NewRequest(
		"GET",
		"https://"+url.Host+"/api/v4/groups/"+prjID,
		nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.NewRequest error[%v]\n", err)
		return
	}
	req.Header.Set("Private-Token", os.Args[1])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.Do error[%v]\n", err)
		return
	}
	defer resp.Body.Close()

	var result GroupInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "Decode error[%v]\n", err)
		return
	}
	for _, p := range result.SharedProjects {
		fmt.Println(p.WebURL)
	}
}
