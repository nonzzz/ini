package lexer

import (
	"ini/internal/tokenizer"
)

type lexer struct {
	source []byte
	cp     byte
	pos    int
	line   int
}

func Lexer(input []byte) *lexer {
	l := &lexer{
		source: input,
		pos:    0,
	}
	l.cp = l.source[l.pos]
	return l
}

func (lexer *lexer) step() {
	lexer.pos += 1
	if lexer.pos > len(lexer.source)-1 {
		lexer.cp = 0
		return
	}
	lexer.cp = lexer.source[lexer.pos]
}

func (lexer *lexer) scanStatement() []byte {
	pos := lexer.pos
	for lexer.cp != '\n' && lexer.cp != '=' && lexer.cp != ' ' && lexer.cp != 0 {
		lexer.step()
	}
	return lexer.source[pos:lexer.pos]
}

func (lexer *lexer) Next() tokenizer.Tokenizer {
	for lexer.cp != 0 {
		switch lexer.cp {
		case ' ', '\t':
			lexer.step()
			continue
		case '\r', '\n':
			lexer.step()
			lexer.line++
			continue
		case '[':
			lexer.step()
			return tokenizer.NewToken(tokenizer.TOpenBrace, "[", lexer.line)
		case ']':
			lexer.step()
			return tokenizer.NewToken(tokenizer.TCloseBrace, "]", lexer.line)
		case ';':
			return tokenizer.NewToken(tokenizer.TSection, ";", lexer.line)
		case '#':
			return tokenizer.NewToken(tokenizer.TComment, "#", lexer.line)
		case '=':
			lexer.step()
			return tokenizer.NewToken(tokenizer.TAssign, "=", lexer.line)
		default:
			return tokenizer.NewToken(tokenizer.TKey, string(lexer.scanStatement()), lexer.line)
		}

	}
	return tokenizer.NewToken(tokenizer.TEof, "Eof", lexer.line)
}
