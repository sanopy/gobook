package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	categories := map[string]string{
		"C": "Other",
		"L": "Letter",
		"M": "Mark",
		"N": "Number",
		"P": "Punctuation",
		"S": "Symbol",
		"Z": "Separator",
	}

	counts := make(map[string]int) // counts of Unicode categories
	invalid := 0                   // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for key := range categories {
			if unicode.Is(unicode.Categories[key], r) {
				counts[key]++
			}
		}
	}
	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s:%s\t%d\n", c, categories[c], n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
