package lexer

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
	"github.com/nonzzz/ini/internal/tokenizer"
)

func lexToken(input string) tokenizer.T {
	l := Lexer([]byte(input))
	return l.Token()

}

func TestTokens(t *testing.T) {
	expected := []struct {
		content string
		token   tokenizer.T
	}{
		{"", tokenizer.TEof},
		{";", tokenizer.TComment},
		{"#", tokenizer.TComment},
		{"=", tokenizer.TAssign},
	}

	for _, tok := range expected {
		t.Run(tok.content, func(t *testing.T) {
			test.AssertEqual(t, lexToken(tok.content), tok.token)
		})
	}
}

func expectComment(t *testing.T, input string, expect string) {
	t.Helper()
	t.Run(input, func(t *testing.T) {
		l := Lexer([]byte(input))
		test.AssertEqual(t, l.Literal(), expect)
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
		l := Lexer([]byte(input))
		l.Next()
		test.AssertEqual(t, l.Literal(), expect)
	})
}

func TestMultipleComment(t *testing.T) {
	expectMultipleComment(t, ";123\n# 456", " 456")
	expectMultipleComment(t, ";123\n#; 456", "; 456")
}

func expectCommentAndSection(t *testing.T, input, expect1, expect2 string) {
	t.Helper()
	t.Run(input, func(t *testing.T) {
		l := Lexer([]byte(input))
		test.AssertEqual(t, l.Literal(), expect1)
		l.Next()
		test.AssertEqual(t, l.Literal(), expect2)
	})
}

func TestCommentFollowSection(t *testing.T) {
	expectCommentAndSection(t, "[s1];comment1", "s1", "comment1")
	expectCommentAndSection(t, "[s2]\n;comment2", "s2", "comment2")
	expectCommentAndSection(t, "[s3]\n#;comment3", "s3", ";comment3")
}

func TestCommentFollowVariable(t *testing.T) {

	txt := "a = 1 \r\nb=2;999"

	expected := []struct {
		content string
		token   tokenizer.T
		line    int
	}{
		{"a", tokenizer.TKey, 0},
		{"=", tokenizer.TAssign, 0},
		{"1", tokenizer.TValue, 0},
		{"b", tokenizer.TKey, 1},
		{"=", tokenizer.TAssign, 1},
		{"2", tokenizer.TValue, 1},
		{"999", tokenizer.TComment, 1},
	}

	l := Lexer([]byte(txt))
	for _, tok := range expected {
		test.AssertEqual(t, l.Literal(), tok.content)
		test.AssertEqual(t, l.Line(), tok.line)
		test.AssertEqual(t, l.Token(), tok.token)
		l.Next()
	}
}
