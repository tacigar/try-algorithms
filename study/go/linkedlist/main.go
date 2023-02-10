package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	prev  *Node
	next  *Node
	value string
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func NewList() *List {
	return &List{head: nil, tail: nil, length: 0}
}

func (l *List) InsertAt(value string, index int) error {
	prevLength := l.length
	if index > prevLength {
		return errors.New("out of range")
	}
	l.length += 1

	if prevLength == 0 {
		newNode := &Node{value: value}
		l.head = newNode
		l.tail = newNode
		return nil
	}

	if index == 0 {
		prevHead := l.head
		newNode := &Node{value: value, next: prevHead}
		prevHead.prev = newNode
		l.head = newNode
		return nil
	}

	if index == prevLength {
		prevTail := l.tail
		newNode := &Node{value: value, prev: prevTail}
		prevTail.next = newNode
		l.tail = newNode
		return nil
	}

	if index == prevLength-1 {
		prevTail := l.tail
		newNode := &Node{value: value, next: prevTail, prev: prevTail.prev}
		prevTail.prev.next = newNode
		prevTail.prev = newNode
		return nil
	}

	prevNode := l.head
	for i := 0; i < index-1; i++ {
		prevNode = prevNode.next
	}
	newNode := &Node{value: value, prev: prevNode, next: prevNode.next}
	prevNode.next.prev = newNode
	prevNode.next = newNode

	return nil
}

func (l *List) DeletedAt(index int) error {
	prevLength := l.length
	if index >= prevLength {
		return errors.New("out of range")
	}
	l.length -= 1

	if prevLength == 1 {
		l.head = nil
		l.tail = nil
		return nil
	}

	if index == 0 {
		l.head.next.prev = nil
		l.head = l.head.next
		return nil
	}

	if index == prevLength-1 {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
		return nil
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev

	return nil
}

func (l *List) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	node := l.head
	for {
		if node == nil {
			sb.WriteRune(']')
			return sb.String()
		}
		if node != l.head {
			sb.WriteString(", ")
		}
		sb.WriteString(node.value)
		node = node.next
	}
}

func main() {
	l := NewList()
	l.InsertAt("A", 0)
	fmt.Println(l.String())
	l.InsertAt("C", 1)
	fmt.Println(l.String())
	l.InsertAt("B", 1)
	fmt.Println(l.String())
	l.InsertAt("F", 3)
	fmt.Println(l.String())
	l.InsertAt("E", 3)
	fmt.Println(l.String())
	l.InsertAt("D", 3)
	fmt.Println(l.String())

	l.DeletedAt(4)
	fmt.Println(l.String())
	l.DeletedAt(2)
	fmt.Println(l.String())
	l.DeletedAt(3)
	fmt.Println(l.String())
	l.DeletedAt(0)
	fmt.Println(l.String())
	l.DeletedAt(1)
	fmt.Println(l.String())
	l.DeletedAt(0)
	fmt.Println(l.String())

}
