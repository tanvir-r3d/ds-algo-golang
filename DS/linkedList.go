package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) insert(value int) {
	node := &Node{data: value, next: nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = l.tail.next
	}
}

func (l *List) traverse() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (l *List) removeAtLast() {
	current := l.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
	l.tail = current
}

func (l *List) removeAt(position int) {
	current := l.head
	if position == 1 {
		l.head = l.head.next
	} else if position > 1 {
		for count := 1; count < position-1; count++ {
			current = current.next
		}

		current.next = current.next.next
	}
}

func (l *List) reverse() {
	current := l.head
	var prev *Node
	next := l.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = l.tail
	l.tail = prev
}

func main() {
	list := &List{}
	list.insert(5)
	list.insert(9)
	list.insert(8)
	list.insert(4)
	list.traverse()
	list.reverse()
	list.traverse()
}
