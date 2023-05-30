package ini

import (
	"os"

	"github.com/nonzzz/ini/internal/parser"
	"github.com/nonzzz/ini/pkg/ast"
)

type Ini struct {
	document *ast.Node
}

func New() *Ini {
	return &Ini{}
}

func (i *Ini) LoadFile(file string) (*Ini, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return i, err
	}
	return i.Parse(string(buf)), nil
}

func (i *Ini) Parse(input string) *Ini {
	p := parser.Parser(input)
	i.document = p
	return i
}

func (i *Ini) Marshal2Map() map[string]interface{} {
	//
	return nil
}

func (i *Ini) Marshal2Json() {}

func (i *Ini) Walk(walker ast.Walker) {
	ast.Walk(i.document, walker)
}
