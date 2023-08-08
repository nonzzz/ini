package ini

import (
	"testing"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/test"
)

func CreateINIParse() *Ini {
	i := New()
	_, _ = i.LoadFile("./case/str.ini")
	return i
}

func TestSection(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	s1, _ := selector.Query("s1", SectionKind).Get()
	s2, _ := selector.Query("s2", SectionKind).Get()
	test.AssertEqual(t, s1.Kind(), ast.KSection)
	test.AssertEqual(t, s2.Kind(), ast.KSection)
	test.AssertEqual(t, s1.ChildrenCount(), 3)
}

func TestExpression(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	expr, _ := selector.Query("a", ExpressionKind).Get()
	test.AssertEqual(t, expr.Kind(), ast.KExpression)
}

func TestComemnt(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	expr, _ := selector.Query("follow comment", ExpressionKind).Get()
	test.AssertEqual(t, expr.Kind(), ast.KComment)
}
