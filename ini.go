package ini

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/nonzzz/ini/internal/parser"
	"github.com/nonzzz/ini/internal/printer"
	"github.com/nonzzz/ini/pkg/ast"
)

type Ini struct {
	document *ast.Document
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
	if i.document == nil {
		return nil
	}

	iniMap := make(map[string]interface{})

	currentSection := ""

	i.Walk(func(node ast.Node, _ ast.Node) {
		switch t := node.(type) {
		case *ast.SectionNode:
			currentSection = t.Name
			sectionMap := make(map[string]interface{})
			iniMap[currentSection] = sectionMap
		case *ast.ExpressionNode:
			if _, ok := iniMap[currentSection].(map[string]interface{}); !ok {
				iniMap[t.Key] = t.Value
				return
			}
			iniMap[currentSection].(map[string]interface{})[t.Key] = t.Value
		}
	})

	return iniMap
}

func (i *Ini) Marshal2Json() []byte {
	iniMap := i.Marshal2Map()
	if iniMap == nil {
		return nil
	}
	re, _ := json.Marshal(iniMap)
	return re
}

func (i *Ini) Printer() (string, error) {
	if i.document == nil {
		return "", errors.New("[ini]: Node is empty.Please called LoadFile or Parse")
	}
	return printer.Printer(i.document), nil
}

func (i *Ini) Walk(walker ast.Walker) {
	ast.Walk(i.document, walker)
}
