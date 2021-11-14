package main

import (
	"flag"
	"fmt"

	"github.com/sanopy/gobook/ch07/ex06/tempflag"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
