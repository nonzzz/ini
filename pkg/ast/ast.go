package ast

import "github.com/nonzzz/ini/internal/lexer"

// import (
// 	"github.com/nonzzz/ini/internal/lexer"
// )

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
