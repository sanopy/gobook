package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/html"
	"gopl.io/ch5/links"
)

var targetHost string

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, targetUrl string) {
	u, err := url.Parse(targetUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url parse failed: %v\n", err)
		os.Exit(1)
	}
	targetHost = u.Host

	go func() { // debug
		for {
			fmt.Fprintf(os.Stderr, "[DEBUG] crawlTokens: %d\n", len(crawlTokens))
			fmt.Fprintf(os.Stderr, "[DEBUG] downloadTokens: %d\n", len(downloadTokens))
			time.Sleep(5 * time.Second)
		}
	}()

	worklist := make(chan []string)
	go func() { worklist <- []string{targetUrl} }()
	n := 1 // number of pending sends to worklist

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

				u, err := url.Parse(link)
				if err != nil || u.Host != targetHost { // skip if not target host
					continue
				}

				n++
				go func(link string) {
					download(u)
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var downloadTokens = make(chan struct{}, 20)

func download(u *url.URL) {
	downloadTokens <- struct{}{} // acquire a token
	defer func() {
		<-downloadTokens // release the token
	}()

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
		path += "/index.html"
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

	if err := printHtmlRewritePath(f, resp); err != nil {
		log.Printf("%v", err)
		return
	}
}

func printHtmlRewritePath(f io.Writer, resp *http.Response) error {
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("html parse failed: %v", err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if n.Data == "a" && a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host == targetHost { // rewrite path
					cwd, _ := os.Getwd()
					path := cwd + "/dump" + link.Path
					n.Attr[i].Val = path
				}
			}
		} else if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "link" || n.Data == "script") {
			for i, tag := range n.Attr {
				if (n.Data == "img" && tag.Key != "src") || (n.Data == "link" && tag.Key != "href") || (n.Data == "script" && tag.Key != "src") {
					continue
				}
				link, err := resp.Request.URL.Parse(tag.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host == targetHost { // image, css, jsはローカルに保存していない
					path := "https://" + targetHost + link.Path
					n.Attr[i].Val = path
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	html.Render(f, doc)
	return nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// crawlTokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var crawlTokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	crawlTokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-crawlTokens // release the token

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
