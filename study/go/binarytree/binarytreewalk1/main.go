package main

import "fmt"

const None = -1

type Node struct {
	Parent int
	Left   int
	Right  int
}

type Tree []Node

func NewTree(size int) Tree {
	t := make(Tree, size)
	for i := 0; i < size; i++ {
		t[i].Parent = None
		t[i].Left = None
		t[i].Right = None
	}
	return t
}

func (t Tree) Insert(id, left, right int) {
	t[id].Left = left
	t[id].Right = right
	if left != None {
		t[left].Parent = id
	}
	if right != None {
		t[right].Parent = id
	}
}

func (t Tree) PreorderWalk(rootID int) []int {
	result := []int{}

	var fn func(id int)
	fn = func(id int) {
		result = append(result, id)
		if t[id].Left != None {
			fn(t[id].Left)
		}
		if t[id].Right != None {
			fn(t[id].Right)
		}
	}
	fn(rootID)

	return result
}

func (t Tree) InorderWalk(rootID int) []int {
	result := []int{}

	var fn func(id int)
	fn = func(id int) {
		if t[id].Left != None {
			fn(t[id].Left)
		}
		result = append(result, id)
		if t[id].Right != None {
			fn(t[id].Right)
		}
	}
	fn(rootID)

	return result
}

func (t Tree) PostorderWalk(rootID int) []int {
	result := []int{}

	var fn func(id int)
	fn = func(id int) {
		if t[id].Left != None {
			fn(t[id].Left)
		}
		if t[id].Right != None {
			fn(t[id].Right)
		}
		result = append(result, id)
	}
	fn(rootID)

	return result
}

func main() {
	t := NewTree(9)
	t.Insert(0, 1, 4)
	t.Insert(1, 2, 3)
	t.Insert(2, -1, -1)
	t.Insert(3, -1, -1)
	t.Insert(4, 5, 8)
	t.Insert(5, 6, 7)
	t.Insert(6, -1, -1)
	t.Insert(7, -1, -1)
	t.Insert(8, -1, -1)

	rootID := 0
	for {
		if t[rootID].Parent == None {
			break
		}
		rootID = t[rootID].Parent
	}

	fmt.Println(len(t), len(t.PreorderWalk(rootID)))

	fmt.Printf("%d\n", t.PreorderWalk(rootID))
	fmt.Printf("%d\n", t.InorderWalk(rootID))
	fmt.Printf("%d\n", t.PostorderWalk(rootID))
}
