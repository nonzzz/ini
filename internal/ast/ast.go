package ast

import "ini/internal/tokenizer"

type NodeType int

type Node interface {
	Type() NodeType
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
	// ReplaceChild(new, old Node)
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
	Name tokenizer.Tokenizer
}

type VNode struct {
	BaseNode
	Key   tokenizer.Tokenizer
	Value tokenizer.Tokenizer
	Index int
	Line  int
}

type CommentNode struct {
	BaseNode
	Name tokenizer.Tokenizer
}

type Document struct {
	BaseNode
}

func NewDocument() *Document {

	return &Document{}
}

func (node *BaseNode) Type() NodeType {
	return 0
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

// func (node *BaseNode) ReplaceChild() {}
