package ini

import (
	"bytes"
	"strings"

	"github.com/nonzzz/ini/pkg/ast"
)

type visitor struct {
	s [][]string
	IniVisitor
}

func (v *visitor) Section(node *ast.Section) {
	if node.Line < len(v.s) {
		v.s[node.Line] = append(v.s[node.Line], "["+node.Literal+"]")
	}
}

func (v *visitor) Comment(node *ast.Comment) {
	if node.Line < len(v.s) {
		v.s[node.Line] = append(v.s[node.Line], ";"+node.Literal)
	}
}

func (v *visitor) Expression(node *ast.Expression) {
	if node.Line < len(v.s) {
		v.s[node.Line] = append(v.s[node.Line], node.Key+" = "+node.Value)
	}
}

func (ini *Ini) String() string {
	if ini.document == nil {
		return ""
	}

	// We preallocate a possible number of rows
	s := make([][]string, ini.document.Line)

	v := &visitor{
		s: s,
	}

	ini.Accept(v)

	var bf bytes.Buffer
	for line, row := range v.s {
		if len(row) > 0 {
			bf.WriteString(strings.Join(row, " "))
		}

		if line > 0 && line < len(v.s)-1 {
			bf.WriteString("\n")
		}
	}
	return bf.String()
}
