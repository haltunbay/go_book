package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const APIURL = "https://api.github.com/repos/haltunbay/golangapi/issues"

type Issue struct {
	Title     string `json:"title, omitempty"`
	Body      string
	Assignees []string
	Milestone int
	State     string `json: "state, omitempty"`
	Labels    []string
}

func main() {
	var i = Issue{Title: "hello api", Body: "hello issue body"}
	//, Assignees: []string{}	, milestone: 1, labels: []string{"hello", "world", "api"}}
	createIssue(&i)
}

func createIssue(issue *Issue) {
	postBody, err := json.Marshal(issue)
	fmt.Println(string(postBody))
	if err != nil {
		log.Fatalf("failed to marshal issue: %s", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(postBody))

	if err != nil {
		log.Fatalf("error creating request %s", err)
	}
	req.Header.Add("Authorization", "token f3816411a88b0aa59815f90448dd355f5380e909")
	req.Header.Add("accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)

	// resp, err := http.Post(APIURL, "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Errorf("Unable to post: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("err reading response %s", err)
	}
	fmt.Println(string(body))
}
