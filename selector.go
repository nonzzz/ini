package ini

import (
	"errors"
	"fmt"

	"github.com/nonzzz/ini/internal/ast"
)

// Ini node slector

const (
	SectionKind    = ast.KSection
	ExpressionKind = ast.KExpression
	CommentKind    = ast.KComment
)

type AttributeBindings struct {
	Id    string
	Text  string
	Key   string
	Value string
}

type Operate interface {
	Get() (ast.Element, error)
	Set(bindings AttributeBindings) bool
	Delete() bool
	InsertBefore(node ast.Element) bool
	InsertAfter(node ast.Element) bool
}

type Selector interface {
	Section(id string) Operate
	Comment(id string) Operate
	Expression(key string) Operate
}

type operate struct {
	node ast.Element
	Id   string
	pos  int
}

type selector struct {
	ast ast.Element
}

func isEmptyNode(node ast.Element) bool {
	return node.ChildrenCount() == 0
}

func getValue(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func serializationBindings(previousBindings, bindings AttributeBindings) map[string]interface{} {
	props := make(map[string]interface{}, 4)
	props["id"] = getValue(bindings.Id, previousBindings.Id)
	props["text"] = getValue(bindings.Text, previousBindings.Text)
	props["key"] = getValue(bindings.Key, previousBindings.Key)
	props["value"] = getValue(bindings.Value, previousBindings.Value, previousBindings.Value)
	return props
}

func NewSelector(ini *Ini) Selector {
	return &selector{
		ast: ini.document,
	}
}

func CreateNode(kind ast.K) ast.Element {
	return ast.NewNode(kind)
}

func UpdateNodeAttributeBindings(node ast.Element, bindings AttributeBindings) {
	ast.UpdateNode(node.(*ast.Node), serializationBindings(AttributeBindings{}, bindings))
}

func (selector *selector) Section(id string) Operate {
	var n ast.Element = nil
	var pos = 0
	traverse(selector.ast, pos, func(node ast.Element, position int) bool {
		if node.Kind() == ast.KSection && node.Id() == id {
			n = node
			pos = position
			return true
		}
		return false
	})
	return &operate{
		node: n,
		Id:   id,
		pos:  pos,
	}
}

func (selector *selector) Comment(id string) Operate {
	var n ast.Element = nil
	var pos = 0
	traverse(selector.ast, pos, func(node ast.Element, position int) bool {
		if node.Id() == id && node.Kind() == ast.KComment {
			n = node
			pos = position
			return true
		}
		return false
	})
	return &operate{
		node: n,
		Id:   id,
		pos:  pos,
	}
}

func (selector *selector) Expression(key string) Operate {
	var n ast.Element = nil
	var pos = 0
	traverse(selector.ast, pos, func(node ast.Element, position int) bool {
		if node.Attribute().Key == key && node.Kind() == ast.KExpression {
			n = node
			pos = position
			return true
		}
		return false
	})
	return &operate{
		node: n,
		Id:   key,
		pos:  pos,
	}
}

func (op *operate) Get() (ast.Element, error) {
	if op.node == nil {
		return nil, errors.New(fmt.Sprintf("%s%s", "[ini]: can't find node ", op.Id))
	}
	return op.node, nil
}

func (op *operate) Set(bindings AttributeBindings) bool {
	if op.node == nil {
		return false
	}

	ast.UpdateNode(op.node.(*ast.Node), serializationBindings(AttributeBindings{
		Id:    op.node.Id(),
		Text:  op.node.Text(),
		Key:   op.node.Attribute().Key,
		Value: op.node.Attribute().Value,
	}, bindings))
	return true
}

func (op *operate) Delete() bool {
	if op.node == nil {
		return false
	}
	return true
}

func (op *operate) InsertBefore(node ast.Element) bool {
	if op.node == nil {
		return false
	}
	return true
}

func (op *operate) InsertAfter(node ast.Element) bool {
	if op.node == nil {
		return false
	}
	return true
}
