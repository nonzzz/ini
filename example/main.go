package main

import (
	"fmt"

	"github.com/nonzzz/ini"
)

func main() {

	i := ini.New()
	_, _ = i.LoadFile("../case/mem.ini")

	selector := ini.NewSelector(i)

	section1, err := selector.Query("*", ini.SectionKind).Get()

	ini.UpdateNodeAttributeBindings(section1, ini.AttributeBindings{
		Id: "KANNO",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(section1.Children())
	op := ini.NewSelector(section1).Query("Browser", ini.ExpressionKind)
	op.Set(ini.AttributeBindings{
		Value: "Chrome Browser",
	})
	expr, _ := op.Get()
	commentNode := ini.NewNode(ini.CommentKind)
	ini.UpdateNodeAttributeBindings(commentNode, ini.AttributeBindings{
		Id: "followed comment",
	})
	expr.AppendChild(commentNode)
	co, _ := ini.NewSelector(section1).Query("followed comment", ini.CommentKind).Get()
	s, _ := i.Printer()
	fmt.Println(s)
	fmt.Println(co)
}
