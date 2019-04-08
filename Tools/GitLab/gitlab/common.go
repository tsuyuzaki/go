package gitlab

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

type SharedUser struct {
	ID          string
	UserID      int    `json:"user_id"`
	AccessLevel int    `json:"access_level"`
	ExpiresAt   string `json:"expires_at"`
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

func postRequest(token, url string, jsonStr []byte) bool {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
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
