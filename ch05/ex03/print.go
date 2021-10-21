package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printTextNode(doc)
}

func printTextNode(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		fmt.Println(n.Data)
	}

	printTextNode(n.FirstChild)
	printTextNode(n.NextSibling)
	return
}
