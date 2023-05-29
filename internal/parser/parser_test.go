package parser

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
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
