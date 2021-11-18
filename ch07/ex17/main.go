// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node struct {
	Tag   string
	Id    string
	Class string
}

func (n Node) String() string {
	s := fmt.Sprintf("%s", n.Tag)
	if n.Id != "" {
		s += fmt.Sprintf("#%s", n.Id)
	}
	if n.Class != "" {
		s += fmt.Sprintf(".%s", n.Class)
	}
	return s
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Node // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var id, class string
			for _, attr := range tok.Attr {
				if attr.Name.Local == "id" {
					id = attr.Value
				} else if attr.Name.Local == "class" {
					class = attr.Value
				}
			}
			stack = append(stack, Node{tok.Name.Local, id, class}) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%v: %s\n", stack, tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []Node, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if len(y[0]) > 0 && y[0][0] == '#' && x[0].Id == y[0][1:] ||
			len(y[0]) > 0 && y[0][0] == '.' && x[0].Class == y[0][1:] ||
			x[0].Tag == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
