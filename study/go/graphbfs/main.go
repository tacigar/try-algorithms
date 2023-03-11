package main

import "fmt"

func dfs(g [][]int) []int {
	q := []int{}
	q = append(q, 0)

	distance := 0
	d := make([]int, len(g))
	d[0] = distance

	finished := make([]bool, len(g))

	for {
		if len(q) == 0 {
			break
		}
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if !finished[v] {
				finished[v] = true
				q = append(q, v)
				d[v] = d[u] + 1
			}
		}
	}
	return d
}

func main() {
	r := dfs([][]int{
		{1, 3},
		{3},
		{},
		{2},
	})
	fmt.Println(r)
}
