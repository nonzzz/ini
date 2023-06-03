package printer

import "github.com/nonzzz/ini/pkg/ast"

type printer struct {
	code []byte
}

func (p *printer) printNewLine() {
	p.print("\r\n")
}

func (p *printer) print(text string) {
	p.code = append(p.code, text...)
}

func (p *printer) expression(node *ast.ExpressionNode) {

	if len(node.Nodes) > 0 {
		p.print(node.Key)
		p.print(" ")
		p.print("=")
		p.print(node.Value)
		for _, child := range node.Nodes {
			switch c := child.(type) {
			case *ast.CommentNode:
				p.comment(c)
			}
		}
		return
	}
	p.print(node.Key)
	p.print(" ")
	p.print("=")
	p.print(node.Value)
	p.printNewLine()
}

func (p *printer) comment(node *ast.CommentNode) {
	p.print(node.Text)
	p.printNewLine()
}

func (p *printer) section(node *ast.SectionNode) {
	if len(node.Nodes) > 0 {
		p.print("[")
		p.print(node.Name)
		p.print("]")
		for i, child := range node.Nodes {
			switch c := child.(type) {
			case *ast.ExpressionNode:
				if i == 0 {
					p.printNewLine()
				}
				p.expression(c)
			case *ast.CommentNode:
				p.comment(c)
			}
		}
		return
	}
	p.print("[")
	p.print(node.Name)
	p.print("]")
	p.printNewLine()
}

func Printer(tree *ast.Document) string {
	p := printer{}
	for _, node := range tree.Nodes {
		switch n := node.(type) {
		case *ast.SectionNode:
			p.section(n)
		case *ast.CommentNode:
			p.comment(n)
		case *ast.ExpressionNode:
			p.expression(n)
		}
	}
	return string(p.code)
}
