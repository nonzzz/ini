package api

import (
	"encoding/json"
	"ini/internal/ast"
	"ini/internal/parser"
)

type Ini struct {
	document *ast.Document
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

	for c := ini.document.FirstChild(); c != nil; c = c.NextSibling() {

		if bn, ok := c.(*ast.VariableNode); ok {
			maps[bn.Key.Value] = bn.Value.Value
		}

		if sn, ok := c.(*ast.SectionNode); ok {
			secMap := make(map[string]interface{})
			for bn := sn.FirstChild(); bn != nil; bn = bn.NextSibling() {
				if nest, ok := bn.(*ast.VariableNode); ok {
					secMap[nest.Key.Value] = nest.Value.Value
				}
			}
			maps[sn.Token.Value] = secMap
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
