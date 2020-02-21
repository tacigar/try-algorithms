package main

import "fmt"

func calcKnapsackWithMemoization(n int, ws, ps []int, maxw int) int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 100)
	}

	var calcImpl func(n, cw int) int

	calcImpl = func(i, cw int) int {
		if cw > maxw {
			return -100000
		}
		if i == n {
			return 0
		}
		if memo[i][cw] != 0 {
			return memo[i][cw]
		}
		r := calcImpl(i+1, cw)
		l := calcImpl(i+1, cw+ws[i]) + ps[i]
		if r > l {
			memo[i][cw] = r
			return r
		}
		memo[i][cw] = l
		return l
	}

	return calcImpl(0, 0)
}

func calcKnapsackWithDP(n int, ws, ps []int, maxw int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 100)
	}

	for w := 0; w <= maxw; w++ {
		if w >= ws[0] {
			dp[0][w] = ws[0]
		}
	}

	for i := 1; i < n; i++ {
		for w := 0; w <= maxw; w++ {
			if w >= ws[i] {
				r := dp[i-1][w-ws[i]] + ps[i]
				l := dp[i-1][w]
				if r > l {
					dp[i][w] = r
				} else {
					dp[i][w] = l
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n-1][maxw]
}

func main() {
	n := 6
	maxw := 9
	ws := []int{2, 1, 3, 2, 1, 5}
	ps := []int{3, 2, 6, 1, 3, 85}
	fmt.Println(calcKnapsackWithMemoization(n, ws, ps, maxw))
	fmt.Println(calcKnapsackWithDP(n, ws, ps, maxw))
}
