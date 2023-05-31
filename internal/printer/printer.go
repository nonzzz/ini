package printer

import (
	"fmt"
	"strings"

	"github.com/nonzzz/ini/pkg/ast"
)

// Ini printer

type printer struct {
	rs             []string
	currentSecPos  int
	currentExprPos int
}

func (p *printer) printExpression(node *ast.ExpressionNode, parent ast.Node) {
	switch parent.(type) {
	case *ast.Document:
		if len(node.Nodes) > 0 {
			switch node.Nodes[0].(type) {
			case *ast.CommentNode:
				p.rs = append(p.rs, node.Text)
			default:
				p.rs = append(p.rs, fmt.Sprintf("%s\r\n", node.Text))
			}
		} else {
			p.rs = append(p.rs, fmt.Sprintf("%s\r\n", node.Text))
		}

		p.currentExprPos = len(p.rs) - 1
	case *ast.SectionNode:
		s := strings.Builder{}
		s.WriteString(p.rs[p.currentSecPos])
		if len(node.Nodes) > 0 {
			switch node.Nodes[0].(type) {
			case *ast.CommentNode:
				s.WriteString(node.Text)
			default:
				s.WriteString(node.Text)
				s.WriteString("\r\n")
			}
		} else {
			s.WriteString(node.Text)
			s.WriteString("\r\n")
		}
		p.rs[p.currentSecPos] = s.String()
	}
}

func (p *printer) printSection(node *ast.SectionNode) {
	if len(node.Nodes) > 0 {
		switch node.Nodes[0].(type) {
		case *ast.CommentNode:
			p.rs = append(p.rs, fmt.Sprintf("[%s]", node.Name))
		default:
			p.rs = append(p.rs, fmt.Sprintf("[%s]\r\n", node.Name))
		}
	} else {
		p.rs = append(p.rs, fmt.Sprintf("[%s]\r\n", node.Name))
	}
	p.currentSecPos = len(p.rs) - 1
}

func (p *printer) printComment(node *ast.CommentNode, parent ast.Node) {
	s := strings.Builder{}
	switch parent.(type) {
	case *ast.ExpressionNode:
		if p.currentSecPos != -1 {
			s.WriteString(p.rs[p.currentSecPos])
			s.WriteString(node.Text)
			s.WriteString("\r\n")
			p.rs[p.currentSecPos] = s.String()
		} else {
			s.WriteString(p.rs[p.currentExprPos])
			s.WriteString(node.Text)
			s.WriteString("\r\n")
			p.rs[p.currentExprPos] = s.String()
		}

	case *ast.SectionNode:
		s.WriteString(p.rs[p.currentSecPos])
		s.WriteString(node.Text)
		s.WriteString("\r\n")
		p.rs[p.currentSecPos] = s.String()
	case *ast.Document:
		p.rs = append(p.rs, fmt.Sprintf("%s\r\n", node.Text))
	}
}

func Printer(tree *ast.Document) string {
	p := printer{
		currentSecPos:  -1,
		currentExprPos: -1,
	}
	ast.Walk(tree, func(node, parentNode ast.Node) {
		switch t := node.(type) {
		case *ast.CommentNode:
			p.printComment(t, parentNode)
		case *ast.ExpressionNode:
			p.printExpression(t, parentNode)
		case *ast.SectionNode:
			p.printSection(t)

		}
	})

	return strings.Join(p.rs, "")
}
