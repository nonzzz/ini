package parser

import (
	"testing"

	"github.com/nonzzz/ini/internal/ast"
	"github.com/nonzzz/ini/internal/test"
)

func TestExpression(t *testing.T) {
	input := "a  = 123\n b=[127.0.0.1,192.168.1]\r\n"
	p := Parser(input)
	children := p.Children()
	expectContent := func(node ast.Element) string {
		t.Helper()
		test.AssertEqual(t, node.Type(), ast.KExpression.String())
		test.AssertEqual(t, node.Kind(), ast.KExpression)
		return node.Text()
	}
	test.AssertEqual(t, expectContent(children[0]), "a  = 123")
	test.AssertEqual(t, expectContent(children[1]), "b=[127.0.0.1,192.168.1]")
}

func TestComment(t *testing.T) {
	input := "#comment1 \r\n ;#comment2  \n #    comment3"
	p := Parser(input)
	children := p.Children()
	test.AssertEqual(t, children[0].Id(), "comment1 ")
	test.AssertEqual(t, children[1].Id(), "#comment2  ")
	test.AssertEqual(t, children[2].Id(), "    comment3")
}

func TestExpressionWithComment(t *testing.T) {
	input := "a  = 123 ;comment1\n b=[127.0.0.1,192.168.1]#comment2\r\n ; single line comment"

	p := Parser(input)
	children := p.Children()
	test.AssertEqual(t, p.ChildrenCount(), 3)
	test.AssertEqual(t, children[0].Text(), "a  = 123 ")
	test.AssertEqual(t, children[0].Children()[0].Kind(), ast.KComment)
	test.AssertEqual(t, children[1].Text(), "b=[127.0.0.1,192.168.1]")
	test.AssertEqual(t, children[1].Children()[0].Kind(), ast.KComment)
	test.AssertEqual(t, children[2].Kind(), ast.KComment)
}

func TestSection(t *testing.T) {
	input := "[s1];This is a section \r\n node1 = 123 \n node2 = [127.0.0.1] \r\n [s2] \n [s3] #This is a section2 \n s = 1"
	p := Parser(input)
	children := p.Children()
	test.AssertEqual(t, p.ChildrenCount(), 3)
	test.AssertEqual(t, children[0].Kind(), ast.KSection)
	test.AssertEqual(t, children[0].Id(), "s1")
	test.AssertEqual(t, children[0].ChildrenCount(), 3)
	test.AssertEqual(t, children[0].Children()[0].Kind(), ast.KComment)
	test.AssertEqual(t, children[0].Children()[1].Kind(), ast.KExpression)
	test.AssertEqual(t, children[0].Children()[2].Kind(), ast.KExpression)
	test.AssertEqual(t, children[1].Kind(), ast.KSection)
	test.AssertEqual(t, children[2].Kind(), ast.KSection)
	test.AssertEqual(t, children[2].Children()[0].Kind(), ast.KComment)
	test.AssertEqual(t, children[2].Children()[1].Kind(), ast.KExpression)
}
