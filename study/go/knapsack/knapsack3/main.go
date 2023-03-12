package main

import "fmt"

func ResolveKnapsack(weight, value []int, maxW int) int {
	// dp[n+1][w]: n番目の要素まで見た時、総量がwになる時の価値の最大値

	n := len(weight)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxW+1)
	}

	for i := 0; i < n; i++ {
		for w := 0; w <= maxW; w++ {
			dp[i+1][w] = dp[i][w]
			if w-weight[i] >= 0 {
				v := dp[i][w-weight[i]] + value[i]
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
