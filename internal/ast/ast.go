package ast

import "ini/internal/tokenizer"

type NodeType string

type Node interface {
	Type() NodeType
	NextSibling() Node
	PrevSibling() Node
	Parent() Node
	AppendChild(self, child Node)
	RemoveChild(self, child Node)
	HasChild() bool
}

type BaseNode struct {
	firstChild Node
}

type SectionNode struct {
	BaseNode
	Name tokenizer.Tokenizer
}

type Document struct {
	BaseNode
}

func NewDocument() *Document {

	return &Document{}
}
