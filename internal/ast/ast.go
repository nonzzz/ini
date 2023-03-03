package ast

type NodeKind uint8

const (
	Root = iota
	CommentNode
	KVNode
	Section
)
