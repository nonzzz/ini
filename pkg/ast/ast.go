package ast

import (
	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/internal/tokenizer"
)

type property struct {
	Type    tokenizer.T
	Literal string
	Line    int
	Loc     lexer.Location
}

type expressionProperty struct {
	Type  tokenizer.T
	Key   string
	Value string
	Line  int
	Loc   lexer.Location
}

type Node interface {
	NextSibling() Node
	PrevSibling() Node
	Parent() Node
	SetParent(Node)
	SetPrevSibling(Node)
	SetNextSibling(Node)
	HasChild() bool
	ChildCount() int
	FirstChild() Node
	LastChild() Node
	AppendChild(self, child Node)
	RemoveChild(self, child Node)
}

type BaseNode struct {
	firstChild    Node
	lastChild     Node
	parent        Node
	next          Node
	prev          Node
	childrenCount int
}

type Section struct {
	BaseNode
	property
}

type Expression struct {
	BaseNode
	expressionProperty
}

type Comment struct {
	BaseNode
	property
}

type Document struct {
	BaseNode
}

func newNode(literal string, line int, tok tokenizer.T, loc lexer.Location) property {
	return property{
		Literal: literal,
		Line:    line,
		Type:    tok,
		Loc:     loc,
	}
}

func NewSection(literal string, line int, tok tokenizer.T, loc lexer.Location) *Section {
	return &Section{
		property: newNode(literal, line, tok, loc),
	}
}

func NewComment(literal string, line int, tok tokenizer.T, loc lexer.Location) *Comment {
	return &Comment{
		property: newNode(literal, line, tok, loc),
	}
}

func NewExpression(key, value string, line int, tok tokenizer.T, loc lexer.Location) *Expression {
	return &Expression{
		expressionProperty: expressionProperty{
			Key:   key,
			Value: value,
			Line:  line,
			Type:  tok,
			Loc:   loc,
		},
	}
}

func NewDocument() *Document {
	return &Document{}
}

func (node *BaseNode) NextSibling() Node {
	return node.next
}

func (node *BaseNode) PrevSibling() Node {
	return node.prev
}

func (node *BaseNode) Parent() Node {
	return node.parent
}

func (node *BaseNode) SetParent(n Node) {
	node.parent = n
}

func (node *BaseNode) SetPrevSibling(n Node) {
	node.prev = n
}

func (node *BaseNode) SetNextSibling(n Node) {
	node.next = n
}

func (node *BaseNode) HasChild() bool {
	return node.firstChild != nil
}

func (node *BaseNode) ChildCount() int {
	return node.childrenCount
}

func (node *BaseNode) FirstChild() Node {
	return node.firstChild
}

func (node *BaseNode) LastChild() Node {
	return node.lastChild
}

func (node *BaseNode) AppendChild(self, child Node) {
	if node.firstChild == nil {
		node.firstChild = child
		child.SetNextSibling(nil)
		child.SetPrevSibling(nil)
	} else {
		last := node.lastChild
		last.SetNextSibling(child)
		child.SetPrevSibling(last)
	}
	child.SetParent(self)
	node.lastChild = child
	node.childrenCount++
}

func (node *BaseNode) RemoveChild(self, child Node) {
	if child.Parent() != self {
		return
	}
	prev := child.PrevSibling()
	next := child.NextSibling()
	if prev != nil {
		prev.SetNextSibling(next)
	} else {
		node.firstChild = next
	}
	if next != nil {
		next.SetPrevSibling(prev)
	} else {
		node.lastChild = prev
	}
	child.SetParent(nil)
	child.SetPrevSibling(nil)
	child.SetNextSibling(nil)
	node.childrenCount--

}
