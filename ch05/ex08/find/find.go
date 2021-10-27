package find

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, shouldContinue, nil)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil && !pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		v := forEachNode(c, id, pre, post)
		if v != nil {
			return v
		}
	}

	if post != nil && !post(n, id) {
		return n
	}
	return nil
}

func shouldContinue(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == id {
			return false
		}
	}
	return true
}
