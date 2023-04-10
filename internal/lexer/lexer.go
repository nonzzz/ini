package lexer

import (
	"unicode/utf8"

	"github.com/nonzzz/ini/internal/tokenizer"
)

type Location struct {
	Start int32
	Len   int32
}

func (l *Location) End() int32 {
	return l.Start + l.Len
}

type lexer struct {
	source  []byte
	cp      rune
	pos     int
	token   tokenizer.T
	line    int
	literal string
	loc     Location
}

type Lexical interface {
	Next()
	Token() tokenizer.T
	Line() int
	Literal() string
	Loc() *Location
}

func Lexer(input []byte) *lexer {
	l := &lexer{
		source: input,
	}
	l.step()
	l.Next()
	return l
}

func (lexer *lexer) step() {

	cp, width := utf8.DecodeRune(lexer.source[lexer.pos:])

	if width == 0 {
		cp = -1
	}

	lexer.cp = cp
	lexer.pos += width
}

func (lexer *lexer) Next() {

	for {
		lexer.token = 0
		switch lexer.cp {
		case -1:
			lexer.token = tokenizer.TEof
		case ' ', '\t':
			lexer.step()
			continue
		case '\r', '\n':
			if lexer.cp == '\n' {
				lexer.line++
			}
			lexer.step()
			continue
		case '[':
			pos := lexer.pos
			for {
				if lexer.cp == ']' {
					break
				}
				lexer.step()
			}
			lexer.literal = string(lexer.source[pos : lexer.pos-1])
			lexer.token = tokenizer.TSection
			lexer.loc.Start = int32(pos)
			lexer.loc.Len = int32(lexer.pos-1) - lexer.loc.Start
		case ']':
			lexer.step()
			continue
		case '=':
			lexer.literal = "="
			lexer.token = tokenizer.TAssign
			lexer.step()
		case '#', ';':
			pos := lexer.pos
			for {
				if lexer.cp == '\n' || lexer.cp == -1 {
					break
				}
				lexer.step()
			}
			lexer.literal = string(lexer.source[pos:lexer.pos])
			lexer.loc.Start = int32(pos)
			lexer.loc.Len = int32(lexer.pos) - lexer.loc.Start
			lexer.token = tokenizer.TComment
		default:
			pos := lexer.pos - 1
			space := 0
			for {
				if lexer.cp == '\n' || lexer.cp == '=' || lexer.cp == ';' || lexer.cp == '#' || lexer.cp == -1 || lexer.cp == '\r' {
					break
				}
				if lexer.cp == ' ' || lexer.cp == '\t' {
					space++
				}
				lexer.step()
			}
			if lexer.cp == '=' {
				lexer.token = tokenizer.TKey
			} else {
				lexer.token = tokenizer.TValue
			}
			if lexer.cp == -1 {
				lexer.literal = string(lexer.source[pos:lexer.pos])
			} else {
				if lexer.token == tokenizer.TKey {
					lexer.literal = string(lexer.source[pos : lexer.pos-1-space])
				} else {
					lexer.literal = string(lexer.source[pos : lexer.pos-1])
				}
			}
			lexer.loc.Start = int32(pos)
			lexer.loc.Len = int32(lexer.pos-1) - lexer.loc.Start
		}
		return
	}
}

func (lexer *lexer) Token() tokenizer.T {
	return lexer.token
}

func (lexer *lexer) Line() int {
	return lexer.line
}

func (lexer *lexer) Literal() string {
	return lexer.literal
}

func (lexer *lexer) Loc() *Location {
	return &lexer.loc
}
