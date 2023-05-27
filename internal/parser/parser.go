package parser

import (
	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/pkg/ast"
)

type parser struct {
	input    string
	start    int
	end      int
	tokens   []lexer.Token
	Document *ast.Document
}

func Parser(input string) *parser {
	tokens := lexer.Tokenizer(input)
	p := &parser{
		input:    input,
		end:      len(tokens),
		tokens:   tokens,
		Document: ast.NewDocument(),
	}
	p.parse()
	return p
}

func (parser *parser) parse() {
	for {
		switch parser.current().Kind {
		case lexer.TEof:
			break
		case lexer.TWhitesapce:
			parser.advance()
			continue
		case lexer.TIdent:
			//
		case lexer.TComment:
			//
		}
	}

}

// func (parser *parser) equal(kind lexer.T) {
// 	//
// }

func (parser *parser) current() lexer.Token {
	return parser.at(parser.start)
}

func (parser *parser) at(pos int) lexer.Token {
	if pos < parser.end {
		return parser.tokens[pos]
	}
	if parser.end < len(parser.tokens) {
		return lexer.Token{
			Kind: lexer.TEof,
			Loc:  parser.tokens[parser.end].Loc,
		}
	}
	return lexer.Token{
		Kind: lexer.TEof,
		Loc:  lexer.Loc{Start: int32(len(parser.input))},
	}
}

func (parser *parser) advance() {
	if parser.start < parser.end {
		parser.start++
	}
}

func (parser *parser) eat(kind lexer.T) {
	//
}

// func NewParser(input []byte) *Praser {

// 	p := &Praser{
// 		lexer:    lexer.Lexer(input),
// 		Document: ast.NewDocument(),
// 	}

// 	var currentSection *ast.Section

// 	var expression *ast.Expression

// 	for {
// 		if p.lexer.Token() == tokenizer.TEof {
// 			p.Document.Type = tokenizer.TDocument
// 			p.Document.Loc = *p.lexer.Loc()
// 			p.Document.Line = p.lexer.Line() + 1
// 			break
// 		}
// 		tok := p.lexer.Token()
// 		literal := p.lexer.Literal()
// 		line := p.lexer.Line()
// 		loc := *p.lexer.Loc()
// 		if tok == tokenizer.TSection {
// 			currentSection = ast.NewSection(literal, line, tok, loc)
// 			p.Document.AppendChild(p.Document, currentSection)
// 		}

// 		if tok == tokenizer.TKey {
// 			expression = ast.NewExpression(literal, "", line, tokenizer.TExpression, loc)
// 		}

// 		if tok == tokenizer.TValue && expression != nil {
// 			expression.Value = literal
// 			if currentSection != nil {
// 				currentSection.AppendChild(currentSection, expression)
// 			} else {
// 				p.Document.AppendChild(p.Document, expression)
// 			}
// 		}

// 		if tok == tokenizer.TComment {
// 			p.Document.AppendChild(p.Document, ast.NewComment(literal, line, tok, loc))
// 		}
// 		p.eat(tok)
// 	}
// 	return p
// }

// func (parser *Praser) eat(token tokenizer.T) {
// 	if parser.lexer.Token() == token {
// 		parser.lexer.Next()
// 		return
// 	}
// }
