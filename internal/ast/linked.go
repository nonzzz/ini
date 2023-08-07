package ast

type LinkedNode struct {
	Element string
	prev    *LinkedNode
	next    *LinkedNode
}

type LinkedList struct {
	Head *LinkedNode
	Tail *LinkedNode
	Cap  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Append(element string) {
	node := &LinkedNode{Element: element}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.prev = list.Tail
		list.Tail.next = node
		list.Tail = node
	}
	list.Cap++
}

func (list *LinkedList) Remove(element string) bool {
	cur := list.Head
	for {
		if cur.Element == element {
			prev := cur.prev
			next := cur.next
			if prev != nil {
				prev.next = next
			} else {
				list.Head = next
			}
			if next != nil {
				next.prev = prev
			} else {
				list.Tail = prev
			}
			cur.prev = nil
			cur.next = nil
			list.Cap--
			return true
		}
		cur = cur.Next()
	}
}

func (node *LinkedNode) Prev() *LinkedNode {
	return node.prev
}

func (node *LinkedNode) Next() *LinkedNode {
	return node.next
}
