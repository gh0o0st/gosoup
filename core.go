package gosoup

import (
	"fmt"
	"log"
	"regexp"

	"golang.org/x/net/html"
)

//a depth first search algorithm recursively search through the html
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func getText(n *html.Node) string {
	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result = n.Data
		}
	}
	return result
}

func find(node *html.Node, tag string) *html.Node {
	var result *html.Node
	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			result = n
			return
		}
	}
	forEachNode(node, visit, nil)
	return result
}

func hasAttr(n *html.Node, attr Attr) bool {
	has := false
	for _, a := range n.Attr {
		switch val := attr.Val.(type) {
		case string:
			if a.Key == attr.Key && a.Val == val {
				has = true
			}
		case regexp.Regexp:
			if a.Key == attr.Key && val.FindString(a.Val) != "" {
				has = true
			}
		default:
			log.Fatal("Unsupported type")
		}
	}
	return has
}

func findWithAttr(node *html.Node, tag string, attrs []Attr) *html.Node {
	var result *html.Node

	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			match := true
			for _, attr := range attrs {
				if !hasAttr(n, attr) {
					match = false
				}
			}
			if match {
				result = n
			}
		}
	}

	forEachNode(node, visit, nil)
	return result
}

func findAll(node *html.Node, tag string) []*html.Node {
	var resultList []*html.Node
	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			resultList = append(resultList, n)
		}
	}
	forEachNode(node, visit, nil)
	return resultList
}

func findAllWithAttr(node *html.Node, tag string, attrs []Attr) []*html.Node {
	var resultList []*html.Node

	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			match := true
			for _, attr := range attrs {
				if !hasAttr(n, attr) {
					match = false
				}
			}
			if match {
				resultList = append(resultList, n)
			}
		}
	}

	forEachNode(node, visit, nil)
	return resultList
}

func child(n *html.Node, c string) (*html.Node, error) {
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		if i.Data == c {
			return i, nil
		}
	}
	return nil, fmt.Errorf("no such child: %s", c)
}

//get the string representation of the node
//for easy debug
