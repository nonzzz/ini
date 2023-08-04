package printer

import (
	"github.com/nonzzz/ini/internal/ast"
)

type printer struct {
	code []byte
}

func (p *printer) printNewLine() {
	p.print("\r\n")
}

func (p *printer) print(text string) {
	p.code = append(p.code, text...)
}

func (p *printer) expression(node ast.Element) {

	attr := node.Attribute()
	p.print(attr.Key)
	p.print(" ")
	p.print("=")
	p.print(" ")
	p.print(attr.Value)
	if node.ChildrenCount() > 0 {
		children := node.Children()
		if children[0].Kind() == ast.KComment {
			p.comment(children[0])
		}
	} else {
		p.printNewLine()
	}
}

func (p *printer) comment(node ast.Element) {
	p.print("#")
	p.print(node.Id())
	p.printNewLine()
}

func (p *printer) section(node ast.Element) {
	p.print("[")
	p.print(node.Id())
	p.print("]")
	if node.ChildrenCount() > 0 {
		// ensure followed comment
		children := node.Children()
		if node.Loc().Column != children[0].Loc().Column {
			p.printNewLine()
		}
		for _, child := range children {
			switch child.Kind() {
			case ast.KComment:
				p.comment(child)
			case ast.KExpression:
				p.expression(child)
			}
		}
	} else {
		p.printNewLine()
	}
}

func Printer(tree ast.Element) string {
	p := printer{}
	for _, element := range tree.Children() {
		switch element.Kind() {
		case ast.KSection:
			p.section(element)
		case ast.KComment:
			p.comment(element)
		case ast.KExpression:
			p.expression(element)
		}
	}
	return string(p.code)
}
