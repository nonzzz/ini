package lexer

import (
	"unicode/utf8"

	"github.com/nonzzz/ini/internal/tokenizer"
)

type lexer struct {
	source  []byte
	cp      rune
	pos     int
	token   tokenizer.T
	line    int
	literal string
}

type Lexical interface {
	Next()
	Token() tokenizer.T
	Line() int
	Literal() string
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

	if cp == '\n' {
		lexer.line++

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
			lexer.token = tokenizer.TComment
		default:
			pos := lexer.pos - 1
			for {
				if lexer.cp == '\n' || lexer.cp == '=' || lexer.cp == ';' || lexer.cp == '#' || lexer.cp == ' ' || lexer.cp == -1 || lexer.cp == '\r' {
					break
				}
				lexer.step()
			}
			if lexer.cp == -1 {
				lexer.literal = string(lexer.source[pos:lexer.pos])
			} else {
				lexer.literal = string(lexer.source[pos : lexer.pos-1])
			}
			lexer.skipWhiteSpace()
			if lexer.cp == '=' {
				lexer.token = tokenizer.TKey
				return
			}
			lexer.token = tokenizer.TValue
			return
		}
		return
	}
}

func (lexer *lexer) skipWhiteSpace() {
	for {
		if lexer.cp != ' ' && lexer.cp != '\t' {
			break
		}
		lexer.step()
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
