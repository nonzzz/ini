package ast

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
	Children() []Element
	ChildrenCount() int
	Attribute() Attribute
	AppendChild(node Element)
}

type Node struct {
	Element  Element
	kind     K
	children []Element
	key      string
	value    string
	Text     string
}

func (n *Node) Kind() K {
	return n.kind
}

func (n *Node) Type() string {
	return n.kind.String()
}

func (n *Node) Children() []Element {
	return n.children
}

func (n *Node) ChildrenCount() int {
	return len(n.children)
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
	n.children = append(n.children, node)
}
