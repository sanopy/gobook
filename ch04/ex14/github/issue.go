package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssues(repo string) (*ApiResult, error) {
	url := fmt.Sprintf(ApiURL, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &ApiResult{result}, nil
}
