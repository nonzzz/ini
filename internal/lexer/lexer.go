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

type Lexical interface {
	Next() tokenizer.Tokenizer
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

func (lexer *lexer) scanComment() []byte {
	pos := lexer.pos
	for {
		if lexer.cp == '\n' {
			break
		}
		lexer.step()
	}
	return lexer.source[pos:lexer.pos]
}

func (lexer *lexer) scanSection() []byte {
	lexer.step()
	literal := lexer.scanComment()
	return literal[0 : len(literal)-1]

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
			literal := string(lexer.scanSection())
			return tokenizer.NewToken(tokenizer.TSection, literal, lexer.line)
		case '#', ';':
			literal := string(lexer.scanComment())
			return tokenizer.NewToken(tokenizer.TComment, literal, lexer.line)
		case '=':
			lexer.step()
			return tokenizer.NewToken(tokenizer.TAssign, "=", lexer.line)
		default:
			literal := string(lexer.scanStatement())
			if lexer.cp == '\n' || lexer.cp == 0 {
				return tokenizer.NewToken(tokenizer.TValue, literal, lexer.line)
			}
			return tokenizer.NewToken(tokenizer.TKey, literal, lexer.line)

		}

	}
	return tokenizer.NewToken(tokenizer.TEof, "Eof", lexer.line)
}
