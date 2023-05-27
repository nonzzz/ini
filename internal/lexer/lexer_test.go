package lexer

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
)

func lexToken(input string) (T, string) {
	tokens := Tokenizer(input)
	if len(tokens) > 0 {
		return tokens[0].Kind, tokens[0].DecodedText(input)
	}
	return TEof, ""

}

func TestTokens(t *testing.T) {
	expected := []struct {
		content string
		text    string
		token   T
	}{
		{"ident", "<ident-token>", TIdent},
		{" ", "<whitespace-token>", TWhitesapce},
		{"[", "<[-token>", TOpenBrace},
		{"]", "<]-token>", TCloseBrace},
		{";", "<comment-token>", TComment},
		{"", "<eof-token>", TEof},
	}

	for _, expr := range expected {
		t.Run(expr.content, func(t *testing.T) {
			tok, _ := lexToken(expr.content)
			test.AssertEqual(t, tok, expr.token)
			test.AssertEqual(t, tok.String(), expr.text)
		})
	}
}

func TestStrings(t *testing.T) {

	expected := []struct {
		content string
		text    string
	}{
		{"ident", "ident"},
		{"[", "["},
		{"]", "]"},
		{";", ";"},
	}
	for _, expr := range expected {
		t.Run(expr.content, func(t *testing.T) {
			_, text := lexToken(expr.content)
			test.AssertEqual(t, text, expr.content)
		})
	}
}

func expectComment(t *testing.T, input string, expect string) {
	t.Helper()
	t.Run(input, func(t *testing.T) {
		tok := Tokenizer(input)
		if len(tok) > 0 {
			test.AssertEqual(t, tok[0].DecodedText(input), expect)
		}
	})
}

func TestComment(t *testing.T) {
	expectComment(t, ";123456", "123456")
	expectComment(t, "#456789", "456789")
	expectComment(t, "# 1 2 3 4", " 1 2 3 4")
	expectComment(t, "; 1 2 3 4", " 1 2 3 4")
	expectComment(t, "#;1111", ";1111")
}

func expectMultipleComment(t *testing.T, input string, expect string) {
	t.Helper()
	t.Run(input, func(t *testing.T) {
		tok := Tokenizer(input)
		if len(tok) > 0 {
			test.AssertEqual(t, tok[1].DecodedText(input), expect)
		}
	})
}

func TestMultipleComment(t *testing.T) {
	expectMultipleComment(t, ";123\n# 456", " 456")
	expectMultipleComment(t, ";123\n#; 456", "; 456")
}

func TestExpression(t *testing.T) {

	input := "a=1 2 3 45 #   comment"

	expected := []struct {
		content string
		token   T
	}{
		{"a", TIdent},
		{"=", TEqual},
		{"1 2 3 45 ", TIdent},
		{"   comment", TComment},
	}
	tok := Tokenizer(input)
	for i, expr := range expected {
		test.AssertEqual(t, tok[i].DecodedText(input), expr.content)
		test.AssertEqual(t, tok[i].Kind, expr.token)
	}
}
