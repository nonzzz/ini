package lexer

import (
	"unicode"
	"unicode/utf8"
)

type T uint8

const endOfFile = -1

const (
	TIdent T = iota
	TWhitesapce
	TEqual
	TOpenBrace
	TCloseBrace
	TComment
	TEof
)

var tokenToString = []string{
	"<ident-token>",
	"<whitespace-token>",
	"<=-token>",
	"<[-token>",
	"<]-token>",
	"<comment-token>",
	"<eof-token>",
}

func (t T) String() string {
	return tokenToString[t]
}

func isWhiteSpace(s rune) bool {
	return s == ' ' || s == '\t'
}

func isComment(s rune) bool {
	return s == ';' || s == '#'
}

func isNewLine(s rune) bool {
	return s == '\n' || s == '\r' || s == '\f'
}

func IsAlphaNumericDash(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s) || s == '_' || s == '-'
}

type Loc struct {
	Start int32
	Len   int32
}

func (loc Loc) End() int32 {
	return loc.Start + loc.Len
}

type Token struct {
	Kind T
	Loc  Loc
}

func (token Token) DecodedText(s string) string {
	raw := s[token.Loc.Start:token.Loc.End()]
	switch token.Kind {
	case TComment:
		return raw[1:]
	}
	return raw
}

type lexer struct {
	source string
	pos    int
	cp     rune
	token  Token
}

// process stream token
func Tokenizer(input string) []Token {
	l := &lexer{
		source: input,
	}
	var tokens []Token
	l.step()
	l.next()
	for l.token.Kind != TEof {
		tokens = append(tokens, l.token)
		l.next()
	}
	return tokens
}

func (lexer *lexer) step() {
	cp, width := utf8.DecodeRuneInString(lexer.source[lexer.pos:])
	if width == 0 {
		cp = -1
	}

	lexer.cp = cp
	lexer.token.Loc.Len = int32(lexer.pos) - lexer.token.Loc.Start
	lexer.pos += width
}

func (lexer *lexer) next() {
	for {
		lexer.token = Token{Loc: Loc{Start: lexer.token.Loc.End()}}
		switch lexer.cp {
		case endOfFile:
			lexer.token.Kind = TEof
		case ' ', '\t':
			lexer.step()
			for {
				if !isWhiteSpace(lexer.cp) {
					break
				}
				lexer.step()
			}
			lexer.token.Kind = TWhitesapce
		case '\r', '\n', '\f':
			if lexer.cp == '\r' {
				lexer.step()
			}
			lexer.step()
			continue
		case '[':
			lexer.step()
			lexer.token.Kind = TOpenBrace
		case ']':
			lexer.step()
			lexer.token.Kind = TCloseBrace
		case '#', ';':
			lexer.step()
			lexer.consumeComments()
			lexer.token.Kind = TComment
		case '=':
			lexer.step()
			lexer.token.Kind = TEqual
		default:
			// if IsAlphaNumericDash(lexer.cp) {

			// }
			lexer.token.Kind = lexer.consumeIdent()
		}
		return //nolint
	}
}

func (lexer *lexer) consumeComments() {
	for {
		if lexer.cp == -1 || lexer.cp == '\n' {
			break
		}
		lexer.step()
	}
}

func (lexer *lexer) consumeIdent() T {
	for {
		if lexer.cp == -1 || lexer.cp == '=' || isNewLine(lexer.cp) || isComment(lexer.cp) {
			break
		}
		lexer.step()
	}
	return TIdent
}
