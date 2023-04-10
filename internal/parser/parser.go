package parser

import (
	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/internal/tokenizer"
)

type Praser struct {
	lexer    lexer.Lexical
	Document *ast.Document
}

func NewParser(input []byte) *Praser {

	p := &Praser{
		lexer:    lexer.Lexer(input),
		Document: ast.NewDocument(),
	}

	var currentSection *ast.Section

	var expression *ast.Expression

	for {
		if p.lexer.Token() == tokenizer.TEof {
			break
		}
		tok := p.lexer.Token()
		literal := p.lexer.Literal()
		line := p.lexer.Line()
		loc := *p.lexer.Loc()
		if tok == tokenizer.TSection {
			currentSection = ast.NewSection(literal, line, tok, loc)
			p.Document.AppendChild(p.Document, currentSection)
		}

		if tok == tokenizer.TKey {
			expression = ast.NewExpression(literal, "", line, tokenizer.TExpression, loc)
		}

		if tok == tokenizer.TValue && expression != nil {
			expression.Value = literal
			if currentSection != nil {
				currentSection.AppendChild(currentSection, expression)
			} else {
				p.Document.AppendChild(p.Document, expression)
			}
		}

		if tok == tokenizer.TComment {
			p.Document.AppendChild(p.Document, ast.NewComment(literal, line, tok, loc))
		}
		p.eat(tok)
	}
	return p
}

func (parser *Praser) eat(token tokenizer.T) {
	if parser.lexer.Token() == token {
		parser.lexer.Next()
		return
	}
}
