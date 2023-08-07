package ini

import (
	"testing"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/test"
)

func CreateINIParse() *Ini {
	i := New()
	i.LoadFile("./case/str.ini")
	return i
}

func TestSection(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	s1, _ := selector.Section("s1").Get()
	s2, _ := selector.Section("s2").Get()
	test.AssertEqual(t, s1.Kind(), ast.KSection)
	test.AssertEqual(t, s2.Kind(), ast.KSection)
	test.AssertEqual(t, s1.ChildrenCount(), 3)
}

func TestExpression(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	expr, _ := selector.Expression("a").Get()
	test.AssertEqual(t, expr.Kind(), ast.KExpression)
}

func TestComemnt(t *testing.T) {
	ini := CreateINIParse()
	selector := NewSelector(ini)
	expr, _ := selector.Comment("follow comment").Get()
	test.AssertEqual(t, expr.Kind(), ast.KComment)
}
