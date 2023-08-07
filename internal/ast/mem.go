package ast

// This file is a ordered map implement.
// Ordered map is a memeory friendly data struct

// Principles
// Implement storage order using linked list
// hash map to store ast nodes

type Mem struct {
	list  *LinkedList
	paris map[string]Element
}

func NewMap() *Mem {
	m := &Mem{}
	m.paris = make(map[string]Element)
	m.list = NewLinkedList()
	return m
}

func (mem *Mem) Set(key string, value Element) {
	if !mem.Has(key) {
		mem.list.Append(key)
	}
	mem.paris[key] = value
}

func (mem *Mem) Get(key string) Element {
	return mem.paris[key]
}

func (mem *Mem) Delete(key string) bool {
	if mem.Has(key) {
		s := mem.list.Remove(key)
		if s {
			delete(mem.paris, key)
			return s
		}
	}
	return false
}

func (mem *Mem) Has(key string) bool {
	_, ok := mem.paris[key]
	return ok
}

func (mem *Mem) Len() int {
	return mem.list.Cap
}

func (mem *Mem) Paris() map[string]Element {
	return mem.paris
}

func (mem *Mem) List() *LinkedList {
	return mem.list
}
