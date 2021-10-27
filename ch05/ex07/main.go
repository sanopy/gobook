package main

import (
	"os"

	"github.com/sanopy/gobook/ch05/ex07/pretty"
)

func main() {
	for _, url := range os.Args[1:] {
		pretty.Print(url)
	}
}
