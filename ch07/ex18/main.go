// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"fmt"
	"os"
)

func main() {
	node, err := Parse(os.Stdin)
	if err != nil {
		fmt.Printf("xml parse failed: %v", err)
		os.Exit(1)
	}
	fmt.Println(node)
}
