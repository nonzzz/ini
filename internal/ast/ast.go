package ast

import "github.com/nonzzz/ini/internal/tokenizer"

type Property struct {
	NodeType tokenizer.T
	Literal  string
	Line     int
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

type SectionNode struct {
	BaseNode
	Value Property
}

type ExpressionNode struct {
	BaseNode
	Key   Property
	Value Property
	Index int
	Line  int
}

type CommentNode struct {
	BaseNode
	Value Property
}

type Document struct {
	BaseNode
}

func NewDocument() *Document {
	d := &Document{}
	return d
}

func NewSection(literal string, line int) *SectionNode {
	prop := Property{
		NodeType: tokenizer.TSection,
		Literal:  literal,
		Line:     line,
	}
	return &SectionNode{
		Value: prop,
	}
}

func NewComment(literal string, line int) *CommentNode {
	prop := Property{
		NodeType: tokenizer.TComment,
		Literal:  literal,
		Line:     line,
	}
	return &CommentNode{
		Value: prop,
	}
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
