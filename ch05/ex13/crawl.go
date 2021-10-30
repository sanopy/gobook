package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, targetUrl string) {
	u, err := url.Parse(targetUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url parse failed: %v\n", err)
		os.Exit(1)
	}
	targetHost := u.Host

	worklist := []string{targetUrl}
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				u, err := url.Parse(item)
				if err != nil || u.Host != targetHost {
					continue
				}
				download(u)
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func download(u *url.URL) {
	// get content
	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("download failed: %v", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("download failed: %s", resp.Status)
		return
	}
	defer resp.Body.Close()

	// path to download file
	basePath := "./dump"
	path := basePath + u.Path
	if u.Path == "" {
		path += "/"
	}
	if filepath.Ext(path) == "" {
		path += "index.html"
	}

	// make directory
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Printf("mkdir failed: %v", err)
		return
	}

	// write content to file
	f, err := os.Create(path)
	if err != nil {
		log.Printf("file open failed: %v", err)
		return
	}
	defer f.Close()
	io.Copy(f, resp.Body)
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Crawl the web breadth-first,
// starting from the command-line arguments.
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [url]\n", os.Args[0])
		os.Exit(1)
	}
	breadthFirst(crawl, os.Args[1])
}
