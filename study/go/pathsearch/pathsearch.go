package main

import "fmt"

func searchPathWithMemoization(w, h int) int {
	memo := make([][]int, w)
	for i := range memo {
		memo[i] = make([]int, h)
	}

	var searchImpl func(x, y int) int

	searchImpl = func(x, y int) int {
		if x == w-1 || y == h-1 {
			return 1
		}
		if memo[x][y] > 0 {
			return memo[x][y]
		}
		memo[x][y] = searchImpl(x, y+1) + searchImpl(x+1, y)
		return memo[x][y]
	}

	return searchImpl(0, 0)
}

func searchPathWithDP(w, h int) int {
	dp := make([][]int, w)
	for i := range dp {
		dp[i] = make([]int, h)
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if x == 0 && y == 0 {
				dp[x][y] = 1
			} else if x == 0 {
				dp[x][y] = dp[x][y-1]
			} else if y == 0 {
				dp[x][y] = dp[x-1][y]
			} else {
				dp[x][y] = dp[x-1][y] + dp[x][y-1]
			}
		}
	}
	return dp[w-1][h-1]
}

func main() {
	fmt.Println(searchPathWithMemoization(5, 6))
	fmt.Println(searchPathWithDP(5, 6))
}
