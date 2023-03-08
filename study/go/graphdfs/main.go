package main

import "fmt"

func dfs1(g [][]int) ([]int, []int) {
	stack := []int{}
	stack = append(stack, 0)

	// 0: 未, 1: 中, 2: 済
	colors := make([]int, len(g))

	d := make([]int, len(g))
	f := make([]int, len(g))
	time := 1

	d[0] = time

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		// まだ未探索のノードのみをnextsに追加
		nexts := []int{}
		for _, v := range g[top] {
			if colors[v] == 0 {
				nexts = append(nexts, v)
				time++
				d[v] = time
			}
		}

		if len(nexts) == 0 {
			stack = stack[:len(stack)-1]
			colors[top] = 2
			time++
			f[top] = time
		} else {
			stack = append(stack, nexts...)
			colors[top] = 1
		}
	}

	return d, f
}

func dfs2(g [][]int) ([]int, []int) {
	colors := make([]int, len(g))
	d := make([]int, len(g))
	f := make([]int, len(g))
	time := 1

	var fn func(u int)
	fn = func(u int) {
		for _, v := range g[u] {
			if colors[v] == 0 {
				time++
				colors[v] = 1
				d[v] = time
				fn(v)
			}
		}
		time++
		colors[u] = 2
		f[u] = time
	}
	d[0] = time
	fn(0)

	return d, f
}

func main() {
	g := [][]int{
		{1, 2}, {2, 3}, {4}, {5}, {5}, {},
	}
	{
		d, f := dfs1(g)
		for i := 0; i < len(g); i++ {
			fmt.Printf("%d %d %d\n", i+1, d[i], f[i])
		}
	}
	fmt.Println("===")
	{
		d, f := dfs2(g)
		for i := 0; i < len(g); i++ {
			fmt.Printf("%d %d %d\n", i+1, d[i], f[i])
		}
	}
}
