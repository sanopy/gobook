package pretty

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func Print(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(os.Stdout, doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

var depth int

func startElement(w io.Writer, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s=\"%s\"", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprintf(w, " />\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
		depth++
	case html.DoctypeNode:
		fmt.Fprintf(w, "%*s<!doctype %s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s \"%s\"", a.Key, a.Val)
		}
		fmt.Fprintf(w, ">\n")
	case html.CommentNode:
		fmt.Fprintf(w, "%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		if strings.TrimSpace(n.Data) != "" {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", n.Data)
		}
	}
}

func endElement(w io.Writer, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	case html.DoctypeNode:
	case html.CommentNode:
	case html.TextNode:
		// do nothing...
	}
}
