package gosoup

import (
	"io"

	"golang.org/x/net/html"
)

//Soup represent a HTML node
type Soup struct {
	*html.Node
}

//child soup of string c
func (s *Soup) Child(c string) (*Soup, error) {
	n, err := child(s.Node, c)
	return newSoup(n), err
}

//get node atrributes
func (s *Soup) Attr(key string) (string, bool) {
	for _, attr := range s.Node.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}
func newSoup(n *html.Node) *Soup {
	return &Soup{Node: n}
}

//basicly a wrapper of html.Parse
func Parse(r io.Reader) (*Soup, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	s := newSoup(n)
	return s, nil
}

//get all the text of the html
func (s *Soup) GetText() string {
	return getText(s.Node)
}

//find html node with specific tag
func (s *Soup) Find(tag string) *Soup {
	node := find(s.Node, tag)
	return newSoup(node)
}

//find all the node of specific tag
func (s *Soup) FindAll(tag string) []*Soup {
	nList := findAll(s.Node, tag)
	var soupList []*Soup
	for _, n := range nList {
		soupList = append(soupList, newSoup(n))
	}
	return soupList
}

//find node with specific tag of certain attributes
func (s *Soup) FindWithAttr(tag, key string, val interface{}) *Soup {
	node := findWithAttr(s.Node, tag, key, val)
	return newSoup(node)
}

//find all the nodes with specific tag of certain attributes
func (s *Soup) FindAllWithAttr(tag, key string, val interface{}) []*Soup {
	var soupList []*Soup
	nList := findAllWithAttr(s.Node, tag, key, val)
	for _, n := range nList {
		soupList = append(soupList, newSoup(n))
	}
	return soupList
}
