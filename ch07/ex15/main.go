package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sanopy/gobook/ch07/ex15/eval"
)

func main() {
	fmt.Println("Input an expression")
	fmt.Println("ex) 'x + y', 'x < y ? y - x : x - y', ...")
	fmt.Print("expression: ")
	var s string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	vars := map[eval.Var]bool{}
	expr, err := eval.Parse(s)
	if err == nil {
		err = expr.Check(vars)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %v\n", err)
		os.Exit(1)
	}

	env := eval.Env{}
	if len(vars) > 0 {
		fmt.Println("Input variables")
	}
	for v, val := range vars {
		if val {
			fmt.Printf("%s: ", v)
			var in float64
			fmt.Scan(&in)
			env[v] = in
		}
	}

	res := expr.Eval(env)
	fmt.Printf("Result: %g\n", res)
}
