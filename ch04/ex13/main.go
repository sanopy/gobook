// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sanopy/gobook/ch04/ex13/omdb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected search words for xkcd")
		os.Exit(1)
	}

	movie, err := omdb.SearchMovie(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	downloadImage(movie.Title, movie.Poster)
}

func downloadImage(title, url string) {
	// 拡張子の取得
	pos := strings.LastIndex(url, ".")
	ext := url[pos:]

	// 画像の取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("download failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 画像の書き込み
	filename := title + ext
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.Write(body)
}
