package find

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	if n == nil {
		return nil
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, ElementsByTagName(c, name...)...)
	}

	for _, tag := range name {
		if n.Type == html.ElementNode && n.Data == tag {
			ret = append(ret, n)
			break
		}
	}

	return ret
}
