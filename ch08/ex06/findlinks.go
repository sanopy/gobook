// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type List struct {
	depth int
	links []string
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	depth := flag.Int("depth", 1, "max depth of crawl")
	flag.Parse()
	worklist := make(chan List)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- List{0, flag.Args()} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > *depth {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- List{depth, crawl(link)}
				}(list.depth+1, link)
			}
		}
	}
}
