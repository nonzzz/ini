package ast

import (
	"github.com/nonzzz/ini/internal/lexer"
)

type N uint8

const (
	Doc N = iota
	Expr
	Sec
	Comment
)

var nodeToString = []string{
	"doc",
	"expression",
	"section",
	"comment",
}

func (n N) String() string {
	return nodeToString[n]
}

type Node interface{}

type Base struct {
	Type  N
	Nodes []Node
	Loc   lexer.Loc
	Text  string
	Node
}

type ExpressionNode struct {
	Base
	Key   string
	Value string
}

type Document struct {
	Base
}

type SectionNode struct {
	Base
	Name string
}

type CommentNode struct {
	Base
	Comma string
}

type Walker func(node Node)

func walkImpl(node Node, walker Walker) {
	walker(node)
	switch t := node.(type) {
	case *Document:
		for _, c := range t.Nodes {
			walkImpl(c, walker)
		}
	case *SectionNode:
		for _, c := range t.Nodes {
			walkImpl(c, walker)
		}
	case *CommentNode:
		for _, c := range t.Nodes {
			walkImpl(c, walker)
		}
	case *ExpressionNode:
		for _, c := range t.Nodes {
			walkImpl(c, walker)
		}
	}
}

func Walk(node Node, walker Walker) {
	walkImpl(node, walker)
}
