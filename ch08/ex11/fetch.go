// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Response struct {
	filename string
	url      string
	n        int64
	err      error
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string, done <-chan struct{}) Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{"", url, 0, err}
	}

	req.Cancel = done

	resp, err := http.Get(url)
	if err != nil {
		return Response{"", url, 0, err}
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return Response{"", url, 0, err}
	}
	n, err := io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return Response{local, url, n, err}
}

func mirroredQuery(url []string) Response {
	resp := make(chan Response, len(url))
	done := make(chan struct{})
	for _, u := range url {
		go func(url string) { resp <- fetch(url, done) }(u)

	}

	res := <-resp
	close(done)
	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [url]...\n", os.Args[0])
		os.Exit(1)
	}

	resp := mirroredQuery(os.Args[1:])
	if resp.err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", resp.url, resp.err)
	}
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", resp.url, resp.filename, resp.n)
}
