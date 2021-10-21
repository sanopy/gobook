package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	visit(count, doc)
	for tag, num := range count {
		fmt.Printf("<%s>: %d\n", tag, num)
	}
}

func visit(count map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	visit(count, n.FirstChild)
	visit(count, n.NextSibling)
	return
}
