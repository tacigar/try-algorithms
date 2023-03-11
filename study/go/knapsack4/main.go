package main

import "fmt"

func ResolveKnapsack(ws []int, vs []int, maxW int) int {
	n := len(ws)

	// dp[i+1][w]: i番目まで見た時にw以下の重さで最大の価値
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxW+1)
	}

	for i := 0; i < n; i++ {
		for w := 0; w <= maxW; w++ {
			dp[i+1][w] = dp[i][w]
			if w-ws[i] >= 0 {
				v := dp[i][w-ws[i]] + vs[i]
				if dp[i+1][w] < v {
					dp[i+1][w] = v
				}
			}
		}
	}
	return dp[n][maxW]
}

func main() {
	fmt.Println(ResolveKnapsack([]int{2, 1, 3, 2, 1, 5}, []int{3, 2, 6, 1, 3, 85}, 9))
}
