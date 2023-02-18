package main

import "fmt"

// ノードのIDに連番が振られるという前提での実装

const None = -1

type Node struct {
	Parent       int
	LeftChild    int
	RightSibling int
}

func CreateNodes(size int) []Node {
	nodes := make([]Node, size)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Parent = None
		nodes[i].LeftChild = None
		nodes[i].RightSibling = None
	}
	return nodes
}

func Insert(nodes []Node, id int, children []int) {
	if children == nil || len(children) == 0 {
		return
	}
	nodes[id].LeftChild = children[0]
	for i, child := range children {
		nodes[child].Parent = id
		if i != len(children)-1 {
			nodes[child].RightSibling = children[i+1]
		}
	}
}

func Print(nodes []Node) {
	for id, node := range nodes {
		current := node
		depth := 0
		for {
			if current.Parent == None {
				break
			}
			depth++
			current = nodes[current.Parent]
		}

		tp := ""
		if node.Parent == None {
			tp = "root"
		} else if node.LeftChild == None {
			tp = "leaf"
		} else {
			tp = "internal node"
		}

		var children []int
		if node.LeftChild != None {
			children = append(children, node.LeftChild)
			current := nodes[node.LeftChild]
			for {
				if current.RightSibling == None {
					break
				}
				children = append(children, current.RightSibling)
				current = nodes[current.RightSibling]
			}
		}

		fmt.Printf(
			"node %2d: parent = %2d, depth = %2d, %13s %d\n",
			id,
			node.Parent,
			depth,
			tp,
			children,
		)
	}
}

func main() {
	nodes := CreateNodes(13)
	Insert(nodes, 0, []int{1, 4, 10})
	Insert(nodes, 1, []int{2, 3})
	Insert(nodes, 2, []int{})
	Insert(nodes, 3, []int{})
	Insert(nodes, 4, []int{5, 6, 7})
	Insert(nodes, 5, []int{})
	Insert(nodes, 6, []int{})
	Insert(nodes, 7, []int{8, 9})
	Insert(nodes, 8, []int{})
	Insert(nodes, 9, []int{})
	Insert(nodes, 10, []int{11, 12})
	Insert(nodes, 11, []int{})
	Insert(nodes, 12, []int{})
	fmt.Println(nodes)
	Print(nodes)
}
