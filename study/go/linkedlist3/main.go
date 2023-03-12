package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	next  *Node
	prev  *Node
	value string
}

type List struct {
	dummy *Node
	size  int
}

func NewList() *List {
	n := &Node{}
	n.next = n
	n.prev = n
	return &List{
		dummy: n,
		size:  0,
	}
}

func (l *List) PushBack(value string) {
	newNode := &Node{value: value}
	newNode.prev = l.dummy.prev
	newNode.next = l.dummy
	l.dummy.prev.next = newNode
	l.dummy.prev = newNode
	l.size++
}

func (l *List) InsertAt(value string, index int) error {
	if l.size-1 < index {
		fmt.Println(l.size)
		return errors.New("out of range")
	}
	n := l.dummy
	// 目標の挿入場所の手前まで辿る
	for i := 0; i < index; i++ {
		n = n.next
	}
	newNode := &Node{value: value}
	newNode.prev = n
	newNode.next = n.next
	n.next.prev = newNode
	n.next = newNode
	l.size++
	return nil
}

func (l *List) DeleteAt(index int) error {
	if l.size-1 < index {
		return errors.New("out of range")
	}
	n := l.dummy
	// 削除するノードまで辿る
	for i := 0; i <= index; i++ {
		n = n.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
	return nil
}

func (l *List) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	node := l.dummy.next
	for {
		if node == l.dummy {
			sb.WriteRune(']')
			return sb.String()
		}
		if node != l.dummy.next {
			sb.WriteString(", ")
		}
		sb.WriteString(node.value)
		node = node.next
	}
}

func main() {
	l := NewList()
	l.PushBack("A") // A
	fmt.Println(l.String())
	l.PushBack("C") // A, C
	fmt.Println(l.String())
	l.InsertAt("B", 1) // A, B, C
	fmt.Println(l.String())
	l.PushBack("E") // A, B, C, E
	fmt.Println(l.String())
	l.PushBack("F") // A, B, C, E, F
	fmt.Println(l.String())
	l.InsertAt("D", 3) // A, B, C, D, E, F
	fmt.Println(l.String())

	l.DeleteAt(4) // A, B, C, D, F
	fmt.Println(l.String())
	l.DeleteAt(4) // A, B, C, D
	fmt.Println(l.String())
	l.DeleteAt(2) // A, B, D
	fmt.Println(l.String())
	l.DeleteAt(0) // B, D
	fmt.Println(l.String())
	l.DeleteAt(1) // B
	fmt.Println(l.String())
	l.DeleteAt(0)
	fmt.Println(l.String())
}
