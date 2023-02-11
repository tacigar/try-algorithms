package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrOutOfRange = errors.New("out of range")

type Node struct {
	prev  *Node
	next  *Node
	value string
}

type LinkedList struct {
	sentinel *Node
	length   int
}

func NewLinkedList() *LinkedList {
	n := &Node{}
	n.next = n
	n.prev = n
	return &LinkedList{sentinel: n}
}

func (l *LinkedList) InsertAt(value string, index int) error {
	prevLen := l.length
	if index > prevLen {
		return ErrOutOfRange
	}
	l.length += 1

	if index == prevLen {
		n := &Node{prev: l.sentinel.prev, next: l.sentinel, value: value}
		l.sentinel.prev.next = n
		l.sentinel.prev = n
		return nil
	}

	x := l.sentinel
	for i := 0; i < index; i++ {
		x = x.next
	}
	n := &Node{prev: x, next: x.next, value: value}
	x.next.prev = n
	x.next = n
	return nil
}

func (l *LinkedList) DeletedAt(index int) error {
	prevLen := l.length
	if index >= prevLen {
		return ErrOutOfRange
	}
	l.length -= 1

	x := l.sentinel
	for i := 0; i <= index; i++ {
		x = x.next
	}
	x.prev.next = x.next
	x.next.prev = x.prev
	return nil
}

func (l *LinkedList) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	node := l.sentinel.next
	for {
		if node == l.sentinel {
			sb.WriteRune(']')
			return sb.String()
		}
		if node != l.sentinel.next {
			sb.WriteString(", ")
		}
		sb.WriteString(node.value)
		node = node.next
	}
}

func main() {
	l := NewLinkedList()
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
