package ini

import (
	"sort"
	"strings"

	"github.com/nonzzz/ini/internal/ast"
)

type visitor struct {
	lines map[int]string
	IniVisitor
}

func (v *visitor) Section(node *ast.Section) {
	v.lines[node.Line] = "[" + node.Literal + "]"
}

func (v *visitor) Comment(node *ast.Comment) {
	if node.PrevSibling() != nil {
		switch prev := node.PrevSibling().(type) {
		case *ast.Section:
			if prev.Line == node.Line {
				if s, ok := v.lines[node.Line]; ok {
					v.lines[node.Line] = s + ";" + node.Literal
				}
			}
		case *ast.Comment:
			if s, ok := v.lines[node.Line]; ok {
				v.lines[node.Line] = s + ";" + node.Literal
			}
		}
	} else {
		v.lines[node.Line] = ";" + node.Literal
	}
}

func (v *visitor) Expression(node *ast.Expression) {
	v.lines[node.Line] = node.Key + "=" + node.Value
}

func (ini *Ini) String() string {
	if ini.document == nil {
		return ""
	}
	v := &visitor{lines: make(map[int]string)}
	ini.Accept(v)
	lines := make([]int, 0, len(v.lines))
	str := make([]string, 0, len(v.lines))
	for k := range v.lines {
		lines = append(lines, k)
	}
	sort.Ints(lines)
	for _, i := range lines {

		if s, ok := v.lines[i]; ok {
			str = append(str, s)
		}
	}
	return strings.Join(str, "\n")
}
