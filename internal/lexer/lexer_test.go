package lexer

import (
	"fmt"
	"ini/internal/tokenizer"
	"testing"
)

func TestTokens(t *testing.T) {
	c1 := `a = 123
	b=456
	[s1]

	[s8]

	#123456789

	;anothr testvvv

	`
	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TKey, Value: "a", Line: 0},
		{Kind: tokenizer.TAssign, Value: "=", Line: 0},
		{Kind: tokenizer.TValue, Value: "123", Line: 0},
		{Kind: tokenizer.TKey, Value: "b", Line: 1},
		{Kind: tokenizer.TAssign, Value: "=", Line: 1},
		{Kind: tokenizer.TValue, Value: "456", Line: 1},
		{Kind: tokenizer.TSection, Value: "s1", Line: 2},
		{Kind: tokenizer.TSection, Value: "s8", Line: 4},
		{Kind: tokenizer.TComment, Value: "123456789", Line: 6},
		{Kind: tokenizer.TComment, Value: "anothr testvvv", Line: 8},
	}
	l := Lexer([]byte(c1))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			fmt.Println(tok)
			t.Fatalf("%s != %s", tok.Value, ident.Value)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}

}

func TestCommentFollowVariable(t *testing.T) {
	txt := `a=123; This is a  comment variants.
 b=456# This is comment.
 `

	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TKey, Value: "a", Line: 0},
		{Kind: tokenizer.TAssign, Value: "=", Line: 0},
		{Kind: tokenizer.TValue, Value: "123", Line: 0},
		{Kind: tokenizer.TComment, Value: " This is a  comment variants.", Line: 0},
		{Kind: tokenizer.TKey, Value: "b", Line: 1},
		{Kind: tokenizer.TAssign, Value: "=", Line: 1},
		{Kind: tokenizer.TValue, Value: "456", Line: 1},
		{Kind: tokenizer.TComment, Value: " This is comment.", Line: 1},
	}
	l := Lexer([]byte(txt))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			t.Fatalf("%s != %s", tok.Value, ident.Value)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}
}

func TestCommentFollowSection(t *testing.T) {
	txt := `
	[S1];test1
	[S2]#test2
	`

	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TSection, Value: "S1", Line: 1},
		{Kind: tokenizer.TComment, Value: "test1", Line: 1},
		{Kind: tokenizer.TSection, Value: "S2", Line: 2},
		{Kind: tokenizer.TComment, Value: "test2", Line: 2},
	}
	l := Lexer([]byte(txt))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			t.Fatalf("%s != %s", tok.Value, ident.Value)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}
}

func TestSectionWidthEdge(t *testing.T) {
	txt := `
	[S1#1];test1
	[S2#2]#test2
	`
	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TSection, Value: "S1#1", Line: 1},
		{Kind: tokenizer.TComment, Value: "test1", Line: 1},
		{Kind: tokenizer.TSection, Value: "S2#2", Line: 2},
		{Kind: tokenizer.TComment, Value: "test2", Line: 2},
	}
	l := Lexer([]byte(txt))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			t.Fatalf("%s != %s", tok.Value, ident.Value)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}
}

func TestWithOutEndOfLine(t *testing.T) {

	txt := `[s1]
	a=3`
	expected := []tokenizer.Tokenizer{
		{Kind: tokenizer.TSection, Value: "s1", Line: 0},
		{Kind: tokenizer.TKey, Value: "a", Line: 1},
		{Kind: tokenizer.TAssign, Value: "=", Line: 1},
		{Kind: tokenizer.TValue, Value: "3", Line: 1},
		{Kind: tokenizer.TEof, Value: "Eof", Line: 1},
	}
	l := Lexer([]byte(txt))
	for _, ident := range expected {
		tok := l.Next()
		if tok.Kind != ident.Kind {
			t.Fatalf("%s != %s", tok.Kind, ident.Kind)
		}
		if tok.Value != ident.Value {
			t.Fatalf("%s != %s", tok.Value, ident.Value)
		}
		if tok.Line != ident.Line {
			t.Fatalf("%d != %d", tok.Line, ident.Line)
		}
	}
}
