package parser

import (
	"strings"

	"github.com/nonzzz/ini/internal/lexer"
	"github.com/nonzzz/ini/pkg/ast"
)

type parser struct {
	input      string
	start      int
	end        int
	tokens     []lexer.Token
	secPos     int
	hasSec     bool
	approxLine int
}

func Parser(input string) *ast.Document {
	result := lexer.Tokenizer(input)

	p := &parser{
		input:      input,
		end:        len(result.Tokens),
		tokens:     result.Tokens,
		approxLine: result.ApproxLine,
	}
	document := p.parse()
	doc := &ast.Document{}
	doc.Type = ast.Doc
	doc.Loc = lexer.Loc{Start: 0, Len: p.current().Loc.End()}
	doc.Nodes = document

	return doc

}

func (p *parser) parse() []ast.Node {

	// It's just a pre-allocation.
	var document []ast.Node = make([]ast.Node, 0, p.approxLine)
loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhitesapce:
			p.advance()
			continue
		case lexer.TComment:
			comment := &ast.CommentNode{}
			comment.Type = ast.Comment
			comment.Loc = p.current().Loc
			comment.Text = p.decoded()
			comment.Comma = comment.Text[0:1]
			if p.hasSec {
				document[p.secPos].(*ast.SectionNode).Nodes = append(document[p.secPos].(*ast.SectionNode).Nodes, comment)
			} else {
				document = append(document, comment)
			}
			p.advance()
			continue
		case lexer.TIdent:
			// The token type is <ident-token> must be an expression
			// Because parseExpression will consume all tokens until the next <ident-token>
			expr := p.parseExpression()
			if p.hasSec {
				document[p.secPos].(*ast.SectionNode).Nodes = append(document[p.secPos].(*ast.SectionNode).Nodes, expr)
			} else {
				document = append(document, expr)
			}
			continue
		case lexer.TOpenBrace:
			sec := p.parseSection()
			document = append(document, sec)
			p.secPos = len(document) - 1
			continue
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

func (p *parser) parseExpression() (expr *ast.ExpressionNode) {
	expr = &ast.ExpressionNode{}
	expr.Type = ast.Expr
	expr.Loc = p.current().Loc
	expr.Key = strings.TrimSpace(p.decoded())
	rs := strings.Builder{}
	v := strings.Builder{}
	rs.WriteString(p.decoded())
	p.advance()
loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhitesapce, lexer.TOpenBrace:
			rs.WriteString(p.decoded())
			v.WriteString(p.decoded())
			p.advance()
			continue
		case lexer.TCloseBrace:
			rs.WriteString(p.decoded())
			v.WriteString(p.decoded())
			if p.at(p.start+1).Kind == lexer.TComment {
				p.advance()
				continue
			}
			break loop
		case lexer.TEqual:
			rs.WriteString(p.decoded())
			p.eat(lexer.TEqual)
			continue
		case lexer.TComment:
			comment := &ast.CommentNode{}
			comment.Type = ast.Comment
			comment.Loc = p.current().Loc
			comment.Text = p.decoded()
			comment.Comma = comment.Text[0:1]
			expr.Nodes = append(expr.Nodes, comment)
			break loop
		case lexer.TIdent:
			rs.WriteString(p.decoded())
			v.WriteString(p.decoded())
			if p.at(p.start+1).Loc.Column != expr.Loc.Column {
				break loop
			}
			p.advance()
			continue
		}
	}

	expr.Text = rs.String()
	expr.Loc.Len = p.current().Loc.End()
	expr.Value = v.String()
	p.advance()
	return expr
}

func (p *parser) parseSection() (sec *ast.SectionNode) {
	p.hasSec = false
	sec = &ast.SectionNode{}
	sec.Type = ast.Sec
	sec.Loc = p.current().Loc
	rs := strings.Builder{}
	v := strings.Builder{}
	rs.WriteString(p.decoded())
	for p.current().Kind != lexer.TCloseBrace {
		if p.current().Kind == lexer.TEof {
			break
		}
		p.eat(p.current().Kind)
		if p.current().Kind != lexer.TCloseBrace {
			v.WriteString(p.decoded())
		}
		rs.WriteString(p.decoded())
	}
	sec.Text = rs.String()
	sec.Name = v.String()
	sec.Loc.Len = p.current().Loc.End()
	p.advance()
	p.hasSec = true
	return sec
}
