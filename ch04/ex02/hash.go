package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var algo = flag.String("algo", "", "specify hashing algorithm ('sha256', 'sha384' or 'sha512')")

func main() {
	flag.Parse()

	var s string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	switch *algo {
	case "sha384":
		c := sha512.Sum384([]byte(s))
		fmt.Printf("%x\n", c)
	case "sha512":
		c := sha512.Sum512([]byte(s))
		fmt.Printf("%x\n", c)
	default:
		c := sha256.Sum256([]byte(s))
		fmt.Printf("%x\n", c)
	}
}
