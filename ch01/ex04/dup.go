package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	count int
	files map[string]struct{}
}

func main() {
	counts := make(map[string]*Line)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, c := range counts {
		if c.count > 1 {
			var keys []string
			for k := range c.files {
				keys = append(keys, k)
			}
			fmt.Printf("%d:\t%s\n%s\n\n", c.count, line, keys)
		}
	}
}

func countLines(f *os.File, counts map[string]*Line, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		_, ok := counts[s]
		if ok {
			counts[s].count++
		} else {
			counts[s] = &Line{1, make(map[string]struct{})}
		}
		counts[s].files[filename] = struct{}{}
	}
	// 注意：input.Err()からのエラーの可能性を無視している
}
