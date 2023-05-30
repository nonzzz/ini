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

type Node struct {
	Type  N
	Nodes []Node
	Loc   lexer.Loc
	Text  string
}

type Walker func(node *Node)

func walkImpl(node *Node, walker Walker) {
	walker(node)
	for i := range node.Nodes {
		walkImpl(&node.Nodes[i], walker)
	}
}

func Walk(node *Node, walker Walker) {
	walkImpl(node, walker)
}
