package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	head := strings.Index(url, "://") + 3
	var tail int
	if strings.Contains(url[head:], "/") {
		tail = head + strings.Index(url[head:], "/")
	} else {
		tail = len(url)
	}
	filename := url[head:tail] + "_" + time.Now().Format("20060102150405") + ".dump"
	dst, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: %v\n", err)
		return
	}

	nbytes, err := io.Copy(dst, resp.Body)
	dst.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
