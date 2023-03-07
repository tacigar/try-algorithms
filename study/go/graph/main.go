package main

import (
	"fmt"
	"strings"
)

func TransformAdjListToAdjMatrices(adjList [][]int) [][]bool {
	res := make([][]bool, len(adjList))
	for i := 0; i < len(res); i++ {
		res[i] = make([]bool, len(adjList))
	}

	for i := 0; i < len(adjList); i++ {
		for j := 0; j < len(adjList[i]); j++ {
			res[i][adjList[i][j]] = true
		}
	}
	return res
}

func PrintMatrices(m [][]bool) {
	var s strings.Builder
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] {
				s.WriteRune('1')
			} else {
				s.WriteRune('0')
			}
			if j != len(m[i])-1 {
				s.WriteRune(' ')
			} else {
				s.WriteRune('\n')
			}
		}
	}
	fmt.Println(s.String())
}

func main() {
	res := TransformAdjListToAdjMatrices([][]int{
		{1, 3},
		{3},
		{},
		{2},
	})
	PrintMatrices(res)
}
