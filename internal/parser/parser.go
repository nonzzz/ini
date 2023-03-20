package parser

import (
	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/internal/tokenizer"
)

type Praser struct {
	ch       tokenizer.Tokenizer
	lexer    lexer.Lexical
	Document *ast.Document
}

func NewParser(input []byte) *Praser {

	p := &Praser{
		lexer:    lexer.Lexer(input),
		Document: ast.NewDocument(),
	}

	p.ch = p.lexer.Next()

	var currentSection *ast.SectionNode

	var VNode *ast.VariableNode

	for {
		if p.ch.Kind == tokenizer.TEof {
			break
		}
		if p.ch.Kind == tokenizer.TSection {
			currentSection = nil
			currentSection = &ast.SectionNode{
				Token: p.ch,
			}
			p.Document.AppendChild(p.Document, currentSection)
		}
		if p.ch.Kind == tokenizer.TKey {
			VNode = nil
			VNode = &ast.VariableNode{
				Key: p.ch,
			}

		}
		if p.ch.Kind == tokenizer.TValue {

			if VNode != nil {
				VNode.Value = p.ch
				if currentSection != nil {
					currentSection.AppendChild(currentSection, VNode)
				} else {
					p.Document.AppendChild(p.Document, VNode)
				}
			}

		}

		if p.ch.Kind == tokenizer.TComment {
			comment := &ast.CommentNode{
				Token: p.ch,
			}
			p.Document.AppendChild(p.Document, comment)

		}
		p.eat(p.ch.Kind)
	}

	return p
}

func (parser *Praser) eat(token string) {
	if parser.ch.Kind == token {
		parser.ch = parser.lexer.Next()
		return
	}
}
