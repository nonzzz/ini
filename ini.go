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

type Visitor interface {
	Section(*ast.Section)
	Expression(*ast.Expression)
	Comment(*ast.Comment)
}

type IniVisitor struct {
	Visitor
}

func (b *IniVisitor) Section(node *ast.Section) {}

func (b *IniVisitor) Comment(node *ast.Comment) {}

func (b *IniVisitor) Expression(node *ast.Expression) {}

func New() *Ini {
	return &Ini{}
}

func traverse(node ast.Node, v Visitor) {
	traverseHelper(node, v)
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		traverse(child, v)
	}
}

func traverseHelper(node interface{}, visitor Visitor) {
	switch n := node.(type) {
	case *ast.Section:
		visitor.Section(n)
	case *ast.Expression:
		visitor.Expression(n)
	case *ast.Comment:
		visitor.Comment(n)
	}
}

func (ini *Ini) Parse(input string) *Ini {
	parser := parser.NewParser([]byte(input))
	ini.document = parser.Document
	return ini
}

type mapVisitor struct {
	maps map[string]interface{}
	IniVisitor
}

func (m *mapVisitor) Section(node *ast.Section) {
	m.maps[node.Literal] = make(map[string]interface{})
}

func (m *mapVisitor) Expression(node *ast.Expression) {
	if section, ok := node.Parent().(*ast.Section); ok {
		m.maps[section.Literal].(map[string]interface{})[node.Key] = node.Value
		return
	}
	m.maps[node.Key] = node.Value

}

func (ini *Ini) Marshal2Map() map[string]interface{} {

	if ini.document == nil {
		return nil
	}

	v := &mapVisitor{
		maps: make(map[string]interface{}),
	}

	ini.Accept(v)
	return v.maps
}

func (ini *Ini) Marshal2Json() []byte {
	maps := ini.Marshal2Map()
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

func (ini *Ini) Accept(v Visitor) {
	doc := ini.document
	traverse(doc, v)
}
