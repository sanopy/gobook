package main

import (
	"fmt"
	"strings"
)

func stringsPlusOperator(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func stringsJoin(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	arr := []string{"Austin", "Bexar", "Cactus"}

	fmt.Println(stringsPlusOperator(arr))
	fmt.Println(stringsJoin(arr))
}
