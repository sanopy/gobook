package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [url]\n", os.Args[0])
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d\n", words)
	fmt.Printf("images: %d\n", images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}
	if n.Type == html.TextNode {
		words = wordsCount(n.Data)
	}

	w, i := countWordsAndImages(n.FirstChild)
	words += w
	images += i

	w, i = countWordsAndImages(n.NextSibling)
	words += w
	images += i

	return
}

func wordsCount(s string) (words int) {
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		words++
	}
	return
}
