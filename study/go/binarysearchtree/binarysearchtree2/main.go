package main

import "fmt"

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) {
	// If the tree is empty yet, add a root node and return.
	if t.root == nil {
		t.root = &Node{
			value: value,
		}
		return
	}

	var fn func(n *Node)
	fn = func(n *Node) {
		// If the value is still registered, just return.
		if value == n.value {
			return
		}
		if value < n.value {
			if n.left == nil {
				n.left = &Node{value: value, parent: n}
				return
			}
			fn(n.left)
		} else {
			if n.right == nil {
				n.right = &Node{value: value, parent: n}
				return
			}
			fn(n.right)
		}
	}
	fn(t.root)
}

func (t *Tree) Find(value int) *Node {
	var fn func(n *Node) *Node
	fn = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if value == n.value {
			return n
		}
		if value < n.value {
			return fn(n.left)
		}
		return fn(n.right)
	}
	return fn(t.root)
}

func (t *Tree) PreorderWalk() []int {
	var result []int
	if t.root == nil {
		return result
	}
	var fn func(n *Node)
	fn = func(n *Node) {
		result = append(result, n.value)
		if n.left != nil {
			fn(n.left)
		}
		if n.right != nil {
			fn(n.right)
		}
	}
	fn(t.root)
	return result
}

func (t *Tree) InorderWalk() []int {
	var result []int
	if t.root == nil {
		return result
	}
	var fn func(n *Node)
	fn = func(n *Node) {
		if n.left != nil {
			fn(n.left)
		}
		result = append(result, n.value)
		if n.right != nil {
			fn(n.right)
		}
	}
	fn(t.root)
	return result
}

func (t *Tree) Delete(value int) {
	n := t.Find(value)
	if n == nil {
		return
	}

	// Case 1: n is a leaf node.
	if n.left == nil && n.right == nil {
		// Just remove n.
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
		return
	}

	// Case 2: the number of children of n is 1.
	if (n.left == nil && n.right != nil) || (n.left != nil && n.left == nil) {
		// Get the child node of n.
		var c *Node
		if n.left != nil {
			c = n.left
		} else {
			c = n.right
		}
		// Connect the parent node and child node, and remove n.
		if n.parent.left == n {
			n.parent.left = c
			c.parent = n.parent
		} else {
			n.parent.right = c
			c.parent = n.parent
		}
		return
	}

	// Case 3: the number of children of n is 2.
	next := n.right
	for {
		if next.left == nil {
			break
		}
		next = next.left
	}
	// Update the value of n with next.value.
	n.value = next.value

	// Case 3-1: next is a leaf node.
	if next.left == nil && next.right == nil {
		// Just remove next.
		if next.parent.left == next {
			next.parent.left = nil
		} else {
			next.parent.right = nil
		}
		return
	}

	// Case 3-2: the number of children of next is 1 (there is a right node of next)
	if next.parent.left == next {
		next.parent.left = next.right
		next.right = next.parent
	} else {
		next.parent.right = next.right
		next.right = next.parent
	}
	return
}

func main() {
	t := NewTree()
	t.Insert(8)
	t.Insert(2)
	t.Insert(3)
	t.Insert(7)
	t.Insert(22)
	t.Insert(1)

	fmt.Printf("%d\n", t.InorderWalk())
	fmt.Printf("%d\n", t.PreorderWalk())
	t.Delete(3)
	t.Delete(7)
	fmt.Printf("%d\n", t.InorderWalk())
	fmt.Printf("%d\n", t.PreorderWalk())
}
