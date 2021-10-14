package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetIssue(repo string, num int) (*Issue, error) {
	path := fmt.Sprintf(IssueGetPath, repo, num)
	resp, err := callAPI("GET", path, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue(repo string, post *PostIssue) (*Issue, error) {
	path := fmt.Sprintf(IssueCreatePath, repo)
	data, err := json.Marshal(post)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	resp, err := callAPI("POST", path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func UpdateIssue(repo string, post *PostIssue, num int) (*Issue, error) {
	path := fmt.Sprintf(IssueUpdatePath, repo, num)
	data, err := json.Marshal(post)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	resp, err := callAPI("PATCH", path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CloseIssue(repo string, num int) (*Issue, error) {
	path := fmt.Sprintf(IssueUpdatePath, repo, num)
	data, err := json.Marshal(&UpdateStateIssue{"closed"})
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	resp, err := callAPI("PATCH", path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
