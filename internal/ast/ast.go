package ast

import (
	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/internal/mem"
)

type K uint8

const (
	KDocument K = iota
	KSection
	KExpression
	KComment
)

var kindToString = []string{
	"DocumentDeclaration",
	"SectionDeclaration",
	"ExpressionDeclaration",
	"CommentDeclaration",
}

func (k K) String() string {
	return kindToString[k]
}

type Attribute struct {
	Key   string
	Value string
}

type Element interface {
	Kind() K
	Type() string
	Id() string
	Text() string
	Children() []Element
	ChildrenCount() int
	Attribute() Attribute
	AppendChild(node Element)
	AppendChilden(children []Element)
	Loc() lexer.Loc
}

type Node struct {
	kind  K
	id    string
	mem   *mem.Mem
	key   string
	value string
	text  string
	loc   lexer.Loc
}

func (n *Node) Kind() K {
	return n.kind
}

func (n *Node) Id() string {
	return n.id
}

func (n *Node) Text() string {
	return n.text
}

func (n *Node) Type() string {
	return n.kind.String()
}

func (n *Node) Loc() lexer.Loc {
	return n.loc
}

func (n *Node) Children() []Element {
	elements := make([]Element, 0, n.mem.Len())
	attr := n.mem.List().Head
	for attr != nil {
		elements = append(elements, n.mem.Get(attr.Element).(Element))
		attr = attr.Next()
	}
	return elements
}

func (n *Node) ChildrenCount() int {
	return n.mem.Len()
}

func (n *Node) Attribute() Attribute {
	if n.kind == KExpression {
		return Attribute{
			Key:   n.key,
			Value: n.value,
		}
	}
	return Attribute{}
}

func (n *Node) AppendChild(node Element) {
	if node.Kind() == KExpression {
		n.mem.Set(node.Attribute().Key, node)
	} else {
		n.mem.Set(node.Id(), node)
	}
}

func (n *Node) AppendChilden(mem []Element) {
	for _, child := range mem {
		n.AppendChild(child)
	}
}

func (n *Node) SetStringField(field, v string) {
	switch field {
	case "id":
		n.id = v
	case "text":
		n.text = v
	case "key":
		n.key = v
	case "value":
		n.value = v
	}
}

func (n *Node) SetLoc(loc lexer.Loc) {
	n.loc = loc
}

func NewNode(kind K) *Node {
	return &Node{
		kind: kind,
		mem:  mem.NewMap(),
	}
}

func UpdateNode(node *Node, attr map[string]interface{}) {
	for k, v := range attr {
		switch k {
		case "id", "text", "value", "key":
			if value, ok := v.(string); ok {
				node.SetStringField(k, value)
			}
		case "loc":
			if loc, ok := v.(lexer.Loc); ok {
				node.SetLoc(loc)
			}

		}
	}
}
