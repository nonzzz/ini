package parser

import (
	"strings"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/lexer"
)

type parser struct {
	input      string
	start      int
	end        int
	tokens     []lexer.Token
	approxLine int
}

func Parser(input string) *ast.Node {
	result := lexer.Tokenizer(input)
	p := &parser{
		input:      input,
		end:        len(result.Tokens),
		tokens:     result.Tokens,
		approxLine: result.ApproxLine,
	}

	elements := p.parse()
	document := ast.NewNode(ast.KDocument)
	ast.UpdateNode(document, map[string]interface{}{
		"loc": lexer.Loc{Start: 0, Len: p.current().Loc.End()},
	})
	document.AppendChilden(elements)
	return document

}

func (p *parser) parse() []ast.Element {

	// It's just a pre-allocation.
	var document = make([]ast.Element, 0, p.approxLine)
loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhitesapce:
			p.advance()
			continue
		case lexer.TComment:
			document = append(document, p.parseComment())
			p.advance()
		case lexer.TIdent:
			document = append(document, p.parseExpression())
			if p.peek(lexer.TWhitesapce) {
				p.advance()
			}
		case lexer.TOpenBrace:
			p.eat(lexer.TOpenBrace)
			nested := p.convertSection()
			if nested != nil {
				document = append(document, nested)
			}
		default:
			break loop
		}
	}

	return document
}

func (p *parser) current() lexer.Token {
	return p.at(p.start)
}

func (p *parser) at(pos int) lexer.Token {
	if pos < p.end {
		return p.tokens[pos]
	}
	if p.end < len(p.tokens) {
		return lexer.Token{
			Kind: lexer.TEof,
			Loc:  p.tokens[p.end].Loc,
		}
	}
	return lexer.Token{
		Kind: lexer.TEof,
		Loc:  lexer.Loc{Start: int32(len(p.input))},
	}
}

func (p *parser) advance() {
	if p.start < p.end {
		p.start++
	}
}

func (p *parser) peek(kind lexer.T) bool {
	return kind == p.current().Kind
}

func (p *parser) eat(kind lexer.T) bool {
	if p.peek(kind) {
		p.advance()
		return true
	}
	return false
}

func (p *parser) decoded() string {
	return p.current().DecodedText(p.input)
}

func (p *parser) parseExpression() *ast.Node {
	var sb strings.Builder
	node := ast.NewNode(ast.KExpression)
	loc := p.current().Loc
	record := false
	key := ""

	for {
		if p.current().Loc.Column != loc.Column ||
			p.current().Kind == lexer.TEof ||
			p.current().Kind == lexer.TComment {
			break
		}
		if p.current().Kind == lexer.TEqual && !record {
			record = true
			key = sb.String()
		}
		sb.WriteString(p.decoded())
		p.advance()
	}

	if p.current().Kind == lexer.TComment && p.current().Loc.Column == loc.Column {
		node.AppendChild(p.parseComment())
		p.advance()
	}

	raw := sb.String()
	ast.UpdateNode(node, map[string]interface{}{
		"key":   strings.TrimSpace(key),
		"text":  raw,
		"loc":   lexer.Loc{Start: loc.Start, Column: loc.Column, Len: p.current().Loc.End()},
		"value": strings.TrimSpace(raw[len(key)+1:]),	})
	return node
}

func (p *parser) parseComment() *ast.Node {
	node := ast.NewNode(ast.KComment)
	raw := p.decoded()
	ast.UpdateNode(node, map[string]interface{}{
		"text": raw,
		"id":   raw[1:],
		"loc":  p.current().Loc,
	})
	return node
}

func (p *parser) convertSection() *ast.Node {

	var sb strings.Builder
	sb.WriteString("[")
	node := ast.NewNode(ast.KSection)
	loc := p.current().Loc
	for {
		if p.current().Loc.Column != loc.Column ||
			p.current().Kind == lexer.TEof ||
			p.current().Kind == lexer.TCloseBrace {
			if p.current().Kind == lexer.TCloseBrace {
				sb.WriteString("]")
				p.advance()
			}
			break
		}
		sb.WriteString(p.decoded())
		p.advance()
	}
	for {
		if p.at(p.start+1).Kind == lexer.TCloseBrace && p.current().Loc.Column == loc.Column {
			sb.WriteString(p.decoded())
			p.advance()
		} else {
			break
		}
	}

	s := sb.String()
	if strings.HasSuffix(s, "]") {
		ast.UpdateNode(node, map[string]interface{}{
			"text": s,
			"id":   s[1 : len(s)-1],
			"loc":  lexer.Loc{Start: loc.Start, Column: loc.Column, Len: p.current().Loc.End()},
		})
		for p.current().Kind == lexer.TWhitesapce {
			p.advance()
		}
		if p.current().Kind == lexer.TComment {
			node.AppendChild(p.parseComment())
			p.advance()
		}
	nesetd:
		for {
			switch p.current().Kind {
			case lexer.TWhitesapce:
				p.advance()
				continue
			case lexer.TOpenBrace:
				break nesetd
			case lexer.TIdent:
				node.AppendChild(p.parseExpression())
				if p.peek(lexer.TWhitesapce) {
					p.advance()
				}
			default:
				break nesetd
			}

		}
		return node
	}
	return nil
}
