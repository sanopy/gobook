// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"io"
	"net/http"
	"os"
	"time"
)

const BaseURL = "https://api.github.com"
const IssueGetPath = "/repos/%s/issues/%d"
const IssueCreatePath = "/repos/%s/issues"
const IssueUpdatePath = "/repos/%s/issues/%d"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type PostIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdateStateIssue struct {
	State string `json:"state"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func callAPI(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, nil
}
