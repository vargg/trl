package wh

import (
	"golang.org/x/net/html"
)

func getAttribute(node *html.Node, key string) (string, bool) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func checkNode(node *html.Node, attr string, expectedValue string) bool {
	if node.Type == html.ElementNode {
		attrValue, ok := getAttribute(node, attr)
		if ok && attrValue == expectedValue {
			return true
		}
	}
	return false
}

func traverse(node *html.Node, attr string, expectedValue string) *html.Node {
	if checkNode(node, attr, expectedValue) {
		return node
	}

	for curr := node.FirstChild; curr != nil; curr = curr.NextSibling {
		res := traverse(curr, attr, expectedValue)
		if res != nil {
			return res
		}
	}
	return nil
}

func getElementById(n *html.Node, idValue string) *html.Node {
	return traverse(n, "id", idValue)
}

func getElementByClass(n *html.Node, classValue string) *html.Node {
	return traverse(n, "class", classValue)
}
