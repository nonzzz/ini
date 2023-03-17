package lexer

import (
	"fmt"
	"ini/internal/tokenizer"
	"testing"
)

func TestTokens(t *testing.T) {
	c1 := `a = 123
	b=456`
	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TKey, Value: "a", Line: 0},
		{Kind: tokenizer.TAssign, Value: "=", Line: 0},
		{Kind: tokenizer.TKey, Value: "123", Line: 0},
		{Kind: tokenizer.TKey, Value: "b", Line: 1},
		{Kind: tokenizer.TAssign, Value: "=", Line: 1},
		{Kind: tokenizer.TKey, Value: "456", Line: 1},
	}
	l := Lexer([]byte(c1))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			fmt.Println(tok)
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}

}
