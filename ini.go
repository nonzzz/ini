package ini

import (
	"encoding/json"
	"os"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/parser"
)

type Ini struct {
	document *ast.Document
	err      error
}

func New() *Ini {
	return &Ini{}
}

func (ini *Ini) Parse(input string) *Ini {
	parser := parser.NewParser([]byte(input))
	ini.document = parser.Document
	return ini
}

func (ini *Ini) Marshl2Map() map[string]interface{} {

	if ini.document == nil {
		return nil
	}

	maps := make(map[string]interface{})
	for n := ini.document.FirstChild(); n != nil; n = n.NextSibling() {
		if expression, ok := n.(*ast.Expression); ok {
			maps[expression.Key] = expression.Value
		}
		if section, ok := n.(*ast.Section); ok {
			sectionMap := make(map[string]interface{})
			for bn := section.FirstChild(); bn != nil; bn = bn.NextSibling() {
				if nest, ok := bn.(*ast.Expression); ok {
					sectionMap[nest.Key] = nest.Value
				}
			}
			maps[section.Literal] = sectionMap
			continue
		}
	}
	return maps
}

func (ini *Ini) Marshl2Json() []byte {
	maps := ini.Marshl2Map()
	if maps == nil {
		return nil
	}
	re, _ := json.Marshal(maps)
	return re
}

func (ini *Ini) LoadFile(file string) *Ini {

	buf, err := os.ReadFile(file)

	if err != nil {
		ini.err = err
		return ini
	}
	return ini.Parse(string(buf))
}

func (ini *Ini) Err() error {
	return ini.err
}

func (ini *Ini) Ast() *ast.Document {
	return ini.document
}
