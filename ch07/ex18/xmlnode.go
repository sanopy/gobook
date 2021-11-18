package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	return e.string(0)
}

func (e *Element) string(depth int) string {
	indent := strings.Repeat("  ", depth)
	s := fmt.Sprintf("%s<%s", indent, e.Type.Local)
	for _, attr := range e.Attr {
		s += fmt.Sprintf(" %s=\"%s\"", attr.Name.Local, attr.Value)
	}
	s += ">\n"
	for _, c := range e.Children {
		switch c.(type) {
		case CharData:
			s += fmt.Sprintf("  %s%s\n", indent, c)
		case *Element:
			c := c.(*Element)
			s += c.string(depth + 1)
		}
	}
	s += fmt.Sprintf("%s</%s>\n", indent, e.Type.Local)
	return s
}

func Parse(r io.Reader) (*Element, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, nil}
			if len(stack) > 0 {
				idx := len(stack) - 1
				stack[idx].Children = append(stack[idx].Children, e)
			}
			stack = append(stack, e) // push
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("invalid xml structure")
			}
			if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				return nil, fmt.Errorf("invalid xml structure")
			}
			idx := len(stack) - 1
			stack[idx].Children = append(stack[idx].Children, CharData(tok))
		}
	}
	return nil, fmt.Errorf("invalid xml structure")
}
