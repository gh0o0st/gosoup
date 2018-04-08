package gosoup

import (
	"io"

	"golang.org/x/net/html"
)

//Attr
type Attr struct {
	Key string
	Val interface{}
}

//Node represent a HTML node
type Node struct {
	*html.Node
}

//child Node of string c
func (s *Node) Child(c string) (*Node, error) {
	n, err := child(s.Node, c)
	return newNode(n), err
}

//get node atrributes
func (s *Node) Attr(key string) (string, bool) {
	for _, attr := range s.Node.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}
func newNode(n *html.Node) *Node {
	return &Node{Node: n}
}

//basicly a wrapper of html.Parse
func Parse(r io.Reader) (*Node, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	s := newNode(n)
	return s, nil
}

//get all the text of the html
func (s *Node) GetText() string {
	return getText(s.Node)
}

//Find finds node with specific tag and attributes, attrs can be nil, in which case just finding node for specific tag
func (s *Node) Find(tag string, attrs []Attr) *Node {
	if attrs == nil {
		//no attrs presented
		return newNode(find(s.Node, tag))
	} else {
		return newNode(findWithAttr(s.Node, tag, attrs))
	}
	return nil
}

//
func (s *Node) FindAll(tag string, attrs []Attr) (res []*Node) {
	var nodes []*html.Node
	if attrs == nil {
		//no attrs presented
		nodes = findAll(s.Node, tag)

	} else {
		nodes = findAllWithAttr(s.Node, tag, attrs)

	}
	for _, n := range nodes {
		res = append(res, newNode(n))
	}
	return res
}
