package main

import (
	"fmt"

	"github.com/nonzzz/ini"
)

func main() {

	i := ini.New()
	i.LoadFile("../case/mem.ini")

	selector := ini.NewSelector(i)

	section1, err := selector.Section("*").Get()

	if err != nil {
		fmt.Println(err)
		return
	}
	op := ini.NewSelector(section1).Expression("Browser")
	op.Set(ini.AttributeBindings{
		Value: "Chrome Browser",
	})
	expr, _ := op.Get()
	commentNode := ini.CreateNode(ini.CommentKind)
	expr.AppendChild(commentNode)
	ini.UpdateNodeAttributeBindings(commentNode, ini.AttributeBindings{
		Id: "followed comment",
	})
	s, _ := i.Printer()

	fmt.Println(s)
	// fmt.Println(section1.Children()[1])
}
