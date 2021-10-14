// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/sanopy/gobook/ch04/ex12/xkcd"
)

const XkcdDumpPath = "./xkcd.dump.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected search words for xkcd")
		os.Exit(1)
	}

	search(os.Args[1:])
}

func search(words []string) {
	items := loadItems()
	fmt.Printf("Number\t%-32s\tImage\n", "Title")
	for _, item := range items {
		hit := true
		for _, word := range words {
			word = strings.ToLower(word)
			if !(strings.Contains(strings.ToLower(item.Title), word) ||
				strings.Contains(strings.ToLower(item.Transcript), word) ||
				strings.Contains(strings.ToLower(item.Alt), word)) {
				hit = false
				break
			}
		}
		if hit {
			fmt.Printf("%4d\t%-32s\t%s\n", item.Num, item.Title, item.Img)
		}
	}
}

func loadItems() []*xkcd.XkcdItem {
	// ダンプファイルの存在チェック
	_, err := os.Stat(XkcdDumpPath)
	if os.IsNotExist(err) {
		// ダンプファイルが存在しない場合は生成
		return createIndex()
	}

	r, err := os.Open(XkcdDumpPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}

	items := []*xkcd.XkcdItem{}
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	return items
}

func createIndex() []*xkcd.XkcdItem {
	const jobs = 100
	items := []*xkcd.XkcdItem{}
	isFinished := false
	num := 1
	for !isFinished {
		ch := make(chan *xkcd.XkcdItems)
		for i := 0; i < jobs; i++ {
			xkcd.Wg.Add(1)
			go xkcd.CallInfoAPI(num, ch)
			num++
		}
		for i := 0; i < jobs; i++ {
			item := <-ch
			if item.Ok {
				items = append(items, item.Item)
			} else {
				if item.Num == 404 {
					continue
				}
				isFinished = true
			}
		}
		xkcd.Wg.Wait()
	}

	json, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	w, err := os.Create(XkcdDumpPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}
	defer w.Close()
	_, err = w.Write(json)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %q", err)
		os.Exit(1)
	}

	return items
}
