package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sanopy/gobook/ch02/ex02/lengthconv"
	"github.com/sanopy/gobook/ch02/ex02/tempconv"
	"github.com/sanopy/gobook/ch02/ex02/weightconv"
)

func main() {
	var (
		l = flag.Bool("l", false, "convert length")
		t = flag.Bool("t", false, "convert temperature")
		w = flag.Bool("w", false, "convert weight")
	)
	flag.Parse()
	for _, arg := range flag.Args() {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
		}
		if *l {
			f := lengthconv.Feet(v)
			m := lengthconv.Meter(v)
			fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
		}
		if *t {
			c := tempconv.Celsius(v)
			f := tempconv.Fahrenheit(v)
			fmt.Printf("%s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))
		}
		if *w {
			p := weightconv.Pound(v)
			k := weightconv.Kilogram(v)
			fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToKg(p), k, weightconv.KgToP(k))
		}
	}
}
