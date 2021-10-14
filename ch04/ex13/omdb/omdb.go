package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ApiURL = "http://www.omdbapi.com/?apikey=%s&t=%s"

type Movie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Poster   string
}

func SearchMovie(title []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(title, " "))
	url := fmt.Sprintf(ApiURL, os.Getenv("APIKEY"), q)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
