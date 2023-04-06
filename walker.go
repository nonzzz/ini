package ini

import (
	"github.com/nonzzz/ini/internal/ast"
)

type Visitor interface {
	Section(*ast.Section)
	Expression(*ast.Expression)
	Comment(*ast.Comment)
}

type IniVisitor struct {
	Visitor
}

func (b *IniVisitor) Section(node *ast.Section) {}

func (b *IniVisitor) Comment(node *ast.Comment) {}

func (b *IniVisitor) Expression(node *ast.Expression) {}

func Accept(doc *ast.Document, v Visitor) {
	for n := doc.FirstChild(); n != nil; n = n.NextSibling() {
		if expression, ok := n.(*ast.Expression); ok {
			v.Expression(expression)
		}
		if section, ok := n.(*ast.Section); ok {
			for bn := section.FirstChild(); bn != nil; bn = bn.NextSibling() {
				if nest, ok := bn.(*ast.Expression); ok {
					v.Expression(nest)
				}
			}
			v.Section(section)
			continue
		}
		if comment, ok := n.(*ast.Comment); ok {
			v.Comment(comment)
		}
	}
}
