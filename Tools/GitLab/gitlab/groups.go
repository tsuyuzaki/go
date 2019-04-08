package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var accessLevelValues map[string]int

func init() {
	accessLevelValues = map[string]int{
		"10": 10,
		"20": 20,
		"30": 30,
		"40": 40,
	}
}

func AddMembers(token, gURL, strAccessLevel string, usernames []string) {
	accessLevel, ok := gAccessValues[strAccessLevel]
	if !ok {
		fmt.Fprintf(os.Stderr, "Invalid group_access value. access_level must be 10/20/30 or 40. [input is %s.]\n", strAccessLevel)
		return
	}
	parsedGURL, err := url.Parse(gURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url.Parse error[%v]\n", err)
		return
	}
	gID, ok := ToStrID(parsedGURL.Path)
	if !ok {
		return
	}

	for _, username := range usernames {
		if username == "" {
			continue
		}
		uID, ok := getUserID(token, parsedGURL.Host, username)
		if !ok {
			continue
		}
		if !addMember(token, parsedGURL.Host, gID, uID, accessLevel) {
			fmt.Println("ERROR:", gURL)
		}
	}
}

func getUserID(token, host, username string) (int, bool) {
	req, err := http.NewRequest("GET", "https://"+host+"/api/v4/users?username="+username, nil)
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

	var nodes []*Node
	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		fmt.Fprintf(os.Stderr, "Decode error[%v]\n", err)
		return 0, false
	}
	if len(nodes) != 1 {
		fmt.Fprintf(os.Stderr, "Some [%s] are exists.\n", username)
		return 0, false
	}
	return nodes[0].ID, true
}

func addMember(token, host, gID string, uID, accessLevel int) bool {
	input := SharedUser{ID: gID, UserID: uID, AccessLevel: accessLevel, ExpiresAt: ""}
	jsonStr, err := json.Marshal(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal error[%v]\n", err)
		return false
	}
	url := "https://" + host + "/api/v4/groups/" + gID + "/members"
	return postRequest(token, url, jsonStr)
}
