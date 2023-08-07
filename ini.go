package ini

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/parser"
	"github.com/nonzzz/ini/internal/printer"
)

type Ini struct {
	document ast.Element
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

func traverse(node ast.Element, parentNode ast.Element, walker func(node, parentNode ast.Element) bool) {
	skip := walker(node, parentNode)
	if skip {
		return
	}
	if node.ChildrenCount() > 0 {
		for _, child := range node.Children() {
			traverse(child, node, walker)
		}
	}
}

func (i *Ini) Marshal2Map() map[string]interface{} {

	if i.document.ChildrenCount() == 0 {
		return nil
	}

	iniMap := make(map[string]interface{})

	var currentSection string

	traverse(i.document, nil, func(node, parentNode ast.Element) bool {
		switch node.Kind() {
		case ast.KSection:
			currentSection = node.Id()
			sectionMap := make(map[string]interface{}, node.ChildrenCount())
			iniMap[currentSection] = sectionMap
		case ast.KExpression:
			attr := node.Attribute()
			if _, ok := iniMap[currentSection]; !ok {
				iniMap[attr.Key] = attr.Value
			} else {
				iniMap[currentSection].(map[string]interface{})[attr.Key] = attr.Value
			}
		}
		return false
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
