package parser

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
	"github.com/nonzzz/ini/pkg/ast"
)

func TestExpression(t *testing.T) {
	input := "a  = 123\n b=[127.0.0.1,192.168.1]\r\n"

	p := Parser(input)
	n := p.Nodes
	test.AssertEqual(t, n[0].Text, "a  = 123")
	test.AssertEqual(t, n[1].Text, "b=[127.0.0.1,192.168.1]")
}

func TestComment(t *testing.T) {
	input := "#comment1 \r\n ;#comment2  \n #    comment3"
	p := Parser(input)
	n := p.Nodes
	test.AssertEqual(t, n[0].Text, "comment1 ")
	test.AssertEqual(t, n[1].Text, "#comment2  ")
	test.AssertEqual(t, n[2].Text, "    comment3")
}

func TestExpressionWithComment(t *testing.T) {
	input := "a  = 123 ;comment1\n b=[127.0.0.1,192.168.1]#comment2\r\n ; single line comment"

	p := Parser(input)
	n := p.Nodes
	test.AssertEqual(t, len(n), 3)
	test.AssertEqual(t, n[0].Text, "a  = 123 ")
	test.AssertEqual(t, n[0].Nodes[0].Type, ast.Comment)
	test.AssertEqual(t, n[1].Text, "b=[127.0.0.1,192.168.1]")
	test.AssertEqual(t, n[1].Nodes[0].Type, ast.Comment)
	test.AssertEqual(t, n[2].Text, " single line comment")
}

func TestSection(t *testing.T) {
	input := "[s1];This is a section \r\n node1 = 123 \n node2 = [127.0.0.1] \r\n [s2] \n [s3] #This is a section \n s = 1"
	p := Parser(input)
	n := p.Nodes
	test.AssertEqual(t, len(n), 3)
	test.AssertEqual(t, len(n[0].Nodes), 3)
}
