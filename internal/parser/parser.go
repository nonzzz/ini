package parser

import (
	"ini/internal/ast"
	"ini/internal/lexer"
	"ini/internal/tokenizer"
)

type Praser struct {
	ch       tokenizer.Tokenizer
	lexer    lexer.Lexical
	Document *ast.Document
}

func NewParser(input []byte) {

	p := &Praser{
		lexer:    lexer.Lexer(input),
		Document: ast.NewDocument(),
	}

	p.ch = p.lexer.Next()

	var currentSection *ast.SectionNode

	for {
		if p.ch.Kind == tokenizer.TEof {
			break
		}
		if p.ch.Kind == tokenizer.TSection {
			currentSection = nil
			currentSection = &ast.SectionNode{
				Name: p.ch,
			}
			if currentSection != nil {
				//
			}

		}
		p.eat(p.ch.Kind)
	}
}

func (parser *Praser) eat(token string) {
	if parser.ch.Kind == token {
		parser.lexer.Next()
		return
	}
}
