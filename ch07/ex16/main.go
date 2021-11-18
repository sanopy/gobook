package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sanopy/gobook/ch07/ex15/eval"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		expr, err := eval.Parse(r.FormValue("expr"))
		if err == nil {
			err = expr.Check(map[eval.Var]bool{})
		}
		if err != nil {
			fmt.Fprintf(w, "parse failed: %v\n", err)
			return
		}

		env := eval.Env{}
		for key := range r.Form {
			if key == "expr" {
				continue
			}
			v, err := strconv.ParseFloat(r.FormValue(key), 64)
			if err != nil {
				fmt.Fprintf(w, "parse float failed: %v\n", err)
			}
			env[eval.Var(key)] = v
		}

		res := expr.Eval(env)
		fmt.Fprintf(w, "Result: %g\n", res)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
