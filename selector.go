package ini

import (
	"fmt"
	"reflect"

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
}

type selector struct {
	ast ast.Element
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

func NewSelector(accept interface{}) Selector {
	t := reflect.TypeOf(accept)
	if t.Kind() == reflect.Ptr && t.Elem().Name() == "Ini" {
		return &selector{
			ast: accept.(*Ini).document,
		}
	} else if ast, ok := accept.(ast.Element); ok {
		return &selector{
			ast: ast,
		}
	} else {
		panic("invalid type")
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
	if node, ok := selector.ast.ChildrenParis()[id]; ok {
		if node.Kind() == ast.KSection {
			n = node
		}
	}
	return &operate{
		node: n,
		Id:   id,
	}
}

func (selector *selector) Comment(id string) Operate {
	var n ast.Element = nil
	if selector.ast.Kind() == ast.KSection || selector.ast.Kind() == ast.KExpression {
		if node, ok := selector.ast.ChildrenParis()[id]; ok {
			n = node
		}
	} else {
		traverse(selector.ast, nil, func(node, parentNode ast.Element) bool {
			if node.Kind() == ast.KComment {
				if _, ok := parentNode.ChildrenParis()[id]; ok {
					n = node
					return true
				}
			}

			return false
		})
	}
	return &operate{
		node: n,
		Id:   id,
	}
}

func (selector *selector) Expression(key string) Operate {
	var n ast.Element = nil
	if selector.ast.Kind() == ast.KSection {
		if node, ok := selector.ast.ChildrenParis()[key]; ok {
			n = node
		}
	} else {
		traverse(selector.ast, nil, func(node, parentNode ast.Element) bool {
			if node.Kind() == ast.KExpression {
				if _, ok := parentNode.ChildrenParis()[key]; ok {
					n = node
					return true
				}
			}
			return false
		})
	}
	return &operate{
		node: n,
		Id:   key,
	}
}

func (op *operate) Get() (ast.Element, error) {
	if op.node == nil {
		return nil, fmt.Errorf("%s%s", "[ini]: can't find node ", op.Id)
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
