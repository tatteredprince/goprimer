package main

import (
	"fmt"
	"strconv"
	"strings"
)

type List struct {
	data int
	next *List
}

func NewList(values ...int) *List {
	root := &List{data: values[0]}

	next := root
	for _, v := range values[1:] {
		next.next = &List{data: v}
		next = next.next
	}

	return root
}

func (l List) NextNode() *List {
	return l.next
}

func (l *List) String() string {
	str := "["

	for node := l; node != nil; node = node.NextNode() {
		str += strconv.Itoa(node.data) + " "
	}

	return strings.TrimRight(str, " ") + "]"
}

func (l *List) AddNode(value int) {
	var node *List
	for node = l; node.next != nil; node = node.NextNode() {
	}
	node.next = &List{data: value}
}

func (l *List) Reverse() *List {
	var prev *List
	curr := l

	for curr != nil {
		// prev, curr, curr.next = curr, curr.next, prev
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	list := NewList(1, 2, 3, 4, 5)
	list.AddNode(6)
	list.AddNode(7)

	fmt.Println(list)

	list = list.Reverse()

	fmt.Println(list)
}
