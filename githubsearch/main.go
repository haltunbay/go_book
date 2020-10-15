package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json: "total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json: "html_url"`
}

func main() {
	terms := []string{"repo:golang/go", "is:open", "json", "decoder", "created:>2020-06-14"}
	result, err := searchIssues(terms)

	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}

	fmt.Printf("%d items found\n", len(result.Items))

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s \t %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s ", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
