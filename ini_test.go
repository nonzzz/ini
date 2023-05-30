package ini

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
	"github.com/nonzzz/ini/pkg/ast"
)

func TestWalker(t *testing.T) {
	i, _ := New().LoadFile("./str.ini")
	i.Walk(func(node ast.Node) {
		switch t := node.(type) {
		case *ast.SectionNode:
			if t.Name == "s2" {
				t.Name = "newNode"
			}
		}
	})
	expect := "{\"newNode\":{\"c\":\"4 \",\"d\":\"5\"},\"p\":\"0\",\"s1\":{\"a\":\" 1\",\"b\":\" 2\"}}"
	test.AssertEqual(t, string(i.Marshal2Json()), expect)
}
