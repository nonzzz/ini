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
		if vnode, ok := c.(*ast.VNode); ok {
			maps[vnode.Key.Value] = vnode.Value.Value
		}
		if sect_node, ok := c.(*ast.SectionNode); ok {

			secMap := make(map[string]interface{})

			for kv := sect_node.FirstChild(); kv != nil; kv = kv.NextSibling() {
				if kvnode, ok := kv.(*ast.VNode); ok {
					secMap[kvnode.Key.Value] = kvnode.Value.Value
				}
			}

			maps[sect_node.Name.Value] = secMap
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
