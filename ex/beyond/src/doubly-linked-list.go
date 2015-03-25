package main

import (
	"errors" //<1>
	"fmt"
)

type Value int //<2>

type Node struct { //<3>
	Value
	prev, next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Front() *Node { //<4>
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Push(v Value) *List {
	n := &Node{Value: v} //<5>

	if l.head == nil { //<6>
		l.head = n
	} else {
		l.tail.next = n //<7>
		n.prev = l.tail //<8>
	}
	l.tail = n //<9>

	return l
}

var errEmpty = errors.New("List is empty")

func (l *List) Pop() (v Value, err error) {
	if l.tail == nil { //<10>
		err = errEmpty
	} else {
		v = l.tail.Value     //<11>
		l.tail = l.tail.prev //<12>
		if l.tail == nil {
			l.head = nil //<13>
		}
	}

	return v, err
}

func main() {
	l := new(List)

	l.Push(1)
	l.Push(2)
	l.Push(4)

	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v\n", n.Value)
	}

	fmt.Println()

	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("%v\n", v)
	}
}
