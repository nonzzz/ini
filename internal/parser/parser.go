package parser

import (
	"strings"

	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/pkg/ast"
)

type parser struct {
	input  string
	start  int
	end    int
	tokens []lexer.Token
}

func Parser(input string) *ast.Node {
	tokens := lexer.Tokenizer(input)
	p := &parser{
		input:  input,
		end:    len(tokens),
		tokens: tokens,
	}
	document := p.parse()
	return &ast.Node{
		Type:  ast.Doc,
		Loc:   lexer.Loc{Start: 0, Len: p.current().Loc.End()},
		Nodes: document,
	}
}

func (p *parser) parse() []ast.Node {

	var document []ast.Node
loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhitesapce:
			p.advance()
			continue
		case lexer.TComment:
			document = append(document, ast.Node{Type: ast.Comment, Loc: p.current().Loc, Text: p.decoded()})
			p.advance()
			continue
		case lexer.TIdent:
			// The token type is <ident-token> must be an expression
			// Because parseExpression will consume all tokens until the next <ident-token>
			expr := p.parseExpression()
			document = append(document, expr)
			continue
		case lexer.TOpenBrace:
			sec := p.parseSection()
			document = append(document, sec)
			continue
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

func (p *parser) parseExpression() (expr ast.Node) {
	expr = ast.Node{
		Type: ast.Expr,
		Loc:  lexer.Loc{Start: p.current().Loc.Start},
	}

	rs := strings.Builder{}
	rs.WriteString(p.decoded())

	p.advance()

loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhitesapce:
			rs.WriteString(p.decoded())
			p.advance()
			continue
		case lexer.TOpenBrace:
			rs.WriteString(p.decoded())
			p.advance()
			continue
		case lexer.TEqual:
			rs.WriteString(p.decoded())
			p.eat(lexer.TEqual)
			continue
		case lexer.TIdent:
			rs.WriteString(p.decoded())
			if p.at(p.start+1).Kind == lexer.TCloseBrace {
				p.advance()
				rs.WriteString(p.decoded())
			}
			if p.at(p.start+1).Kind == lexer.TComment {
				expr.Nodes = append(expr.Nodes, ast.Node{Type: ast.Comment, Loc: p.current().Loc, Text: p.decoded()})
				p.advance()
			}
			break loop
		}
	}

	expr.Text = rs.String()
	expr.Loc.Len = p.current().Loc.End()
	p.advance()
	return expr
}

func (p *parser) parseSection() (sec ast.Node) {
	sec = ast.Node{
		Type: ast.Sec,
		Loc:  lexer.Loc{Start: p.current().Loc.Start},
	}
	rs := strings.Builder{}
	rs.WriteString(p.decoded())
	// eat current token then turn the loop to consume all
	// tokens until the next token is <[-token>
	p.eat(lexer.TOpenBrace)
	return sec
	// loop:
	// 	for {
	// 		switch p.current().Kind {
	// 		case lexer.TOpenBrace:
	// 			break loop
	// 		case lexer.TCloseBrace:
	// 			rs.WriteString(p.decoded())
	// 			break loop
	// 		case lexer.TWhitesapce:
	// 			p.advance()
	// 			continue
	// 		case lexer.TIdent:
	// 			if p.eat(p.current().Kind) {
	// 				rs.WriteString(p.decoded())
	// 				if p.at(p.start+1).Kind == lexer.TCloseBrace {
	// 					// section is end
	// 					break loop
	// 				} else {
	// 					// parseExpression
	// 					// save the parse result to sec.Nodes
	// 					sec.Nodes = append(sec.Nodes, p.parseExpression())
	// 				}
	// 			}
	// 		case lexer.TComment:
	// 			sec.Nodes = append(sec.Nodes, ast.Node{Type: ast.Comment, Loc: p.current().Loc, Text: p.decoded()})
	// 			p.advance()
	// 			continue
	// 		default:
	// 			break loop
	// 		}

	// }
	// sec.Text = rs.String()
	// sec.Loc.Len = p.current().Loc.End()
	// p.advance()
	// return sec
}

// func (p *parser) parseSection() (sec ast.Node) {
// 	sec = ast.Node{
// 		Type: ast.Sec,
// 		Loc:  lexer.Loc{Start: p.current().Loc.Start},
// 	}
// 	rs := strings.Builder{}
// 	rs.WriteString(p.decoded())
// 	p.eat(lexer.TOpenBrace)
// sec:
// 	for {
// 		switch p.current().Kind {
// 		case lexer.TEof:
// 			break sec
// 		case lexer.TWhitesapce:
// 			rs.WriteString(p.decoded())
// 			p.advance()
// 			continue
// 		case lexer.TComment:
// 			sec.Nodes = append(sec.Nodes, p.parse()...)
// 			p.advance()
// 			continue
// 		case lexer.TCloseBrace:
// 			rs.WriteString(p.decoded())
// 			p.advance()
// 			continue
// 		case lexer.TIdent:
// 			if p.at(p.start+1).Kind == lexer.TCloseBrace {
// 				rs.WriteString(p.decoded())
// 				p.advance()
// 			} else {
// 				sec.Nodes = append(sec.Nodes, p.parse()...)
// 				return sec
// 			}
// 		default:
// 			if p.current().Kind == lexer.TOpenBrace {
// 				return sec
// 			}
// 			break sec
// 		}
// 	}
// 	sec.Text = rs.String()
// 	sec.Loc.Len = p.current().Loc.End()
// 	p.advance()
// 	return sec
// }
