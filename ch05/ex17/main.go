package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sanopy/gobook/ch05/ex17/find"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [url] [tag]...\n", os.Args[0])
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "http request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse failed: %v\n", err)
		os.Exit(1)
	}
	nodes := find.ElementsByTagName(doc, os.Args[2:]...)

	for _, n := range nodes {
		html.Render(os.Stdout, n)
		fmt.Println()
	}
}
