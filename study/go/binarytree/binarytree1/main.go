package main

import (
	"fmt"
	"math"
)

const None = -1

type Node struct {
	Parent int
	Left   int
	Right  int
}

func NewNodes(size int) []Node {
	nodes := make([]Node, size)
	for i := 0; i < size; i++ {
		nodes[i].Parent = None
		nodes[i].Left = None
		nodes[i].Right = None
	}
	return nodes
}

func Insert(nodes []Node, id, left, right int) {
	nodes[id].Left = left
	nodes[id].Right = right
	if left != None {
		nodes[left].Parent = id
	}
	if right != None {
		nodes[right].Parent = id
	}
}

func GetDepthOfNodes(nodes []Node, rootID int) []int {
	depthOfNodes := make([]int, len(nodes))

	var fillDepthOfNodes func(id, depth int)
	fillDepthOfNodes = func(id, depth int) {
		depthOfNodes[id] = depth
		if nodes[id].Left != None {
			fillDepthOfNodes(nodes[id].Left, depth+1)
		}
		if nodes[id].Right != None {
			fillDepthOfNodes(nodes[id].Right, depth+1)
		}
	}
	fillDepthOfNodes(rootID, 0)

	return depthOfNodes
}

func GetHeightOfNodes(nodes []Node, rootID int) []int {
	heightOfNodes := make([]int, len(nodes))

	var fillHeightOfNodes func(id int) int
	fillHeightOfNodes = func(id int) int {
		if nodes[id].Left == None && nodes[id].Right == None {
			heightOfNodes[id] = 0
			return 0
		}

		lh, rh := 0, 0
		if nodes[id].Left != None {
			lh = fillHeightOfNodes(nodes[id].Left)
		}
		if nodes[id].Right != None {
			rh = fillHeightOfNodes(nodes[id].Right)
		}
		h := int(math.Max(float64(lh), float64(rh))) + 1
		heightOfNodes[id] = h
		return h
	}
	fillHeightOfNodes(rootID)

	return heightOfNodes
}

func Print(nodes []Node) {
	rootID := 0
	for {
		if nodes[rootID].Parent == None {
			break
		}
		rootID = nodes[rootID].Parent
	}
	depthOfNodes := GetDepthOfNodes(nodes, rootID)
	heightOfNodes := GetHeightOfNodes(nodes, rootID)

	for i := 0; i < len(nodes); i++ {
		sibling := None
		if nodes[i].Parent != None {
			if nodes[nodes[i].Parent].Left == i {
				sibling = nodes[nodes[i].Parent].Right
			} else {
				sibling = nodes[nodes[i].Parent].Left
			}
		}

		degree := 0
		if nodes[i].Left != None {
			degree++
		}
		if nodes[i].Right != None {
			degree++
		}

		tp := ""
		if nodes[i].Parent == None {
			tp = "root"
		} else if nodes[i].Left == None && nodes[i].Right == None {
			tp = "leaf"
		} else {
			tp = "internal node"
		}

		fmt.Printf(
			"node %d, parent = %2d, sibling = %2d, degree = %d, depth = %d, height = %d, type = %s\n",
			i,
			nodes[i].Parent,
			sibling,
			degree,
			depthOfNodes[i],
			heightOfNodes[i],
			tp,
		)
	}
}

func main() {
	nodes := NewNodes(9)
	Insert(nodes, 0, 1, 4)
	Insert(nodes, 1, 2, 3)
	Insert(nodes, 2, -1, -1)
	Insert(nodes, 3, -1, -1)
	Insert(nodes, 4, 5, 8)
	Insert(nodes, 5, 6, 7)
	Insert(nodes, 6, -1, -1)
	Insert(nodes, 7, -1, -1)
	Insert(nodes, 8, -1, -1)

	Print(nodes)
}
