package main

import "fmt"

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}
	var fn func(n *Node)
	fn = func(n *Node) {
		if n.Value > value {
			if n.Left == nil {
				newNode := &Node{Value: value, Parent: n}
				n.Left = newNode
				return
			}
			fn(n.Left)
		} else {
			if n.Right == nil {
				newNode := &Node{Value: value, Parent: n}
				n.Right = newNode
				return
			}
			fn(n.Right)
		}
	}
	fn(t.Root)
}

func (t *Tree) PreorderWalk() []int {
	var result []int

	var fn func(n *Node)
	fn = func(n *Node) {
		result = append(result, n.Value)
		if n.Left != nil {
			fn(n.Left)
		}
		if n.Right != nil {
			fn(n.Right)
		}
	}
	fn(t.Root)
	return result
}

func (t *Tree) InorderWalk() []int {
	var result []int

	var fn func(n *Node)
	fn = func(n *Node) {
		if n.Left != nil {
			fn(n.Left)
		}
		result = append(result, n.Value)
		if n.Right != nil {
			fn(n.Right)
		}
	}
	fn(t.Root)
	return result
}

func (t *Tree) Find(value int) *Node {
	var fn func(n *Node) *Node
	fn = func(n *Node) *Node {
		if n.Value == value {
			return n
		}
		if n.Left != nil {
			r := fn(n.Left)
			if r != nil {
				return r
			}
		}
		if n.Right != nil {
			r := fn(n.Right)
			if r != nil {
				return r
			}
		}
		return nil
	}
	return fn(t.Root)
}

func (t *Tree) Delete(value int) {
	n := t.Find(value)
	fmt.Println(n, n.Parent)

	// case1: nがLeafの場合
	if n.Left == nil && n.Right == nil {
		// nがルートの場合は木を空にして終了
		if n.Parent == nil {
			t.Root = nil
			return
		}
		if n.Parent.Left != nil && n.Parent.Left == n {
			n.Parent.Left = nil
			return
		}
		if n.Parent.Right != nil && n.Parent.Right == n {
			n.Parent.Right = nil
			return
		}
		return
	}

	// case2: nの子が一つの場合
	if n.Left == nil && n.Right != nil {
		n.Right.Parent = n.Parent
		// nがルートの場合は子をルートにして終了
		if n.Parent == nil {
			t.Root = n.Right
			return
		}
		if n.Parent.Left != nil && n.Parent.Left == n {
			n.Parent.Left = n.Right
			return
		}
		if n.Parent.Right != nil && n.Parent.Right == n {
			n.Parent.Right = n.Right
			return
		}
		return
	}
	if n.Left != nil && n.Right == nil {
		n.Left.Parent = n.Parent
		// nがルートの場合は子をルートにして終了
		if n.Parent == nil {
			t.Root = n.Left
			return
		}
		if n.Parent.Left != nil && n.Parent.Left == n {
			n.Parent.Left = n.Left
			return
		}
		if n.Parent.Right != nil && n.Parent.Right == n {
			n.Parent.Right = n.Left
			return
		}
		return
	}

	// nの子が二つの場合、次節点を探す
	var next *Node
	if n.Right != nil {
		next = n.Right
		for n.Left != nil {
			next = next.Left
		}
	}
	// データを移し、次節点を削除
	n.Value = next.Value

	if next.Parent.Left != nil && next.Parent.Left == next {
		next.Parent.Left = nil
		return
	}
	if next.Parent.Right != nil && next.Parent.Right == next {
		next.Parent.Right = nil
		return
	}
}

func PrintYesOrNo(b bool) {
	if b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	t := NewTree()
	t.Insert(8)
	t.Insert(2)
	t.Insert(3)
	t.Insert(7)
	t.Insert(22)
	t.Insert(1)
	PrintYesOrNo(t.Find(1) != nil)
	PrintYesOrNo(t.Find(2) != nil)
	PrintYesOrNo(t.Find(3) != nil)
	PrintYesOrNo(t.Find(4) != nil)
	PrintYesOrNo(t.Find(5) != nil)
	PrintYesOrNo(t.Find(6) != nil)
	PrintYesOrNo(t.Find(7) != nil)
	PrintYesOrNo(t.Find(8) != nil)
	fmt.Printf("%d\n", t.InorderWalk())
	fmt.Printf("%d\n", t.PreorderWalk())
	t.Delete(3)
	t.Delete(7)
	fmt.Printf("%d\n", t.InorderWalk())
	fmt.Printf("%d\n", t.PreorderWalk())
}
