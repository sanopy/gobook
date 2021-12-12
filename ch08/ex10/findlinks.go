// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sanopy/gobook/ch08/ex10/links"
)

var done = make(chan struct{})

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read 1 byte
		close(done)
		wg.Wait()
		close(worklist)
	}()

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := crawl(link)
					wg.Add(1)
					go func() {
						defer wg.Done()
						select {
						case worklist <- foundLinks:
						case <-done:
							return
						}
					}()
				case <-done:
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				select {
				case unseenLinks <- link:
					seen[link] = true
				case <-done:
					continue
				}
			}
		}
	}
}
