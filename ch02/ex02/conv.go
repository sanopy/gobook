package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
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
		switch {
		case *l:
			f := lengthconv.Feet(v)
			m := lengthconv.Meter(v)
			fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
			fallthrough
		case *t:
			c := tempconv.Celsius(v)
			f := tempconv.Fahrenheit(v)
			fmt.Printf("%s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))
			fallthrough
		case *w:
			p := weightconv.Pound(v)
			k := weightconv.Kilogram(v)
			fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToKg(p), k, weightconv.KgToP(k))
		}
	}
}
